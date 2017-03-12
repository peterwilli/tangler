/*
MIT License

Copyright (c) 2017 Shinya Yagyu

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"strings"

	"github.com/utamaro/giota"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/analyze_tx/", txHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("starting the server at port http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var (
	tmpl    *template.Template
	funcMap = template.FuncMap{
		"localtime": func(stamp int64) string {
			sec := stamp
			var nano int64
			if stamp > 10000000000 {
				nano = stamp % 1000
				sec = stamp / 1000
			}
			return time.Unix(sec, nano).UTC().Format("2006-01-02 15:04:05 MST")
		},
	}
)

func init() {
	var err error
	tmpl = template.New("")
	tmpl.Funcs(funcMap)
	if _, err = tmpl.ParseGlob("www/*.tpl"); err != nil {
		log.Fatal(err)
	}
}

func renderIfError(w http.ResponseWriter, errs ...error) bool {
	for i, err := range errs {
		if err == nil {
			continue
		}
		log.Print(err, i)
		w.WriteHeader(http.StatusNotFound)
		errr := tmpl.ExecuteTemplate(w, "err.tpl", map[string]string{
			"reason": err.Error(),
		})
		if errr != nil {
			log.Fatal(errr)
		}
		return true
	}
	return false
}

func renderTxError(w http.ResponseWriter, trytes giota.Trytes, err error) bool {
	if err == nil {
		return false
	}
	log.Print(err)
	errr := tmpl.ExecuteTemplate(w, "tx.tpl", struct {
		Analyze bool
		Error   string
		Trytes  giota.Trytes
	}{
		true,
		err.Error(),
		trytes,
	})
	if errr != nil {
		log.Print(err)
	}
	return true
}

func txHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if renderIfError(w, err) {
		return
	}
	params := r.PostFormValue("trytes")
	trytes := giota.Trytes(strings.TrimSpace(params))
	if trytes == "" {
		renderTxError(w, "", errors.New("Input transaction trytes above"))
		return
	}
	if err := trytes.IsValid(); renderTxError(w, trytes, err) {
		return
	}
	tx, err := giota.NewTransaction(trytes.Trits())
	if renderTxError(w, trytes, err) {
		return
	}
	err = tmpl.ExecuteTemplate(w, "tx.tpl", struct {
		Analyze bool
		Error   string
		Hash    giota.Trytes
		Trytes  giota.Trytes
		Tx      giota.Transaction
	}{
		true,
		"",
		trytes.Trits().Hash().Trytes(),
		trytes,
		*tx,
	})
	if err != nil {
		log.Print(err)
	}
}

//indexHandler render the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	server := giota.RandomNode()
	//	server := giota.PublicNode[0]
	api := giota.NewAPI(server, nil)

	var ni *giota.GetNodeInfoResponse
	var txs *giota.GetTransactionsToApproveResponse
	var err1 error
	var err2 error
	wd := sync.WaitGroup{}
	wd.Add(2)
	go func() {
		ni, err1 = api.GetNodeInfo()
		wd.Done()
	}()
	go func() {
		anr := &giota.GetTransactionsToApproveRequest{Depth: 0}
		txs, err2 = api.GetTransactionsToApprove(anr)
		wd.Done()
	}()
	wd.Wait()
	if renderIfError(w, err1, err2) {
		return
	}

	err := tmpl.ExecuteTemplate(w, "index.tpl", struct {
		Server   string
		NodeInfo *giota.GetNodeInfoResponse
		Tx       *giota.GetTransactionsToApproveResponse
	}{
		server,
		ni,
		txs,
	})
	if err != nil {
		log.Print(err)
	}
}

//searchHandler render the search result of hash.
func searchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	kind := q.Get("kind")
	hash := q.Get("hash")
	switch kind {
	case "transaction":
		searchTX(w, giota.Trytes(hash))
	case "address":
		searchAddress(w, giota.Address(hash))
	case "bundle":
		searchBundle(w, giota.Trytes(hash))
	default:
		renderIfError(w, errors.New("illegal request"))
	}
}

func searchTX(w http.ResponseWriter, hash giota.Trytes) {
	server := giota.RandomNode()
	//server := giota.PublicNode[0]
	api := giota.NewAPI(server, nil)

	var gt *giota.GetTrytesResponse
	var ni *giota.GetNodeInfoResponse
	var err1 error
	var err2 error
	wd := sync.WaitGroup{}
	wd.Add(2)
	go func() {
		anr := &giota.GetTrytesRequest{Hashes: []giota.Trytes{hash}}
		gt, err1 = api.GetTrytes(anr)
		wd.Done()
	}()
	go func() {
		ni, err2 = api.GetNodeInfo()
		wd.Done()
	}()
	wd.Wait()
	if renderIfError(w, err1, err2) {
		return
	}
	if len(gt.Trytes) == 0 {
		renderIfError(w, errors.New("transaction is not found while GetTrytes"))
		return
	}

	anr := &giota.GetInclusionStatesRequest{Transactions: []giota.Trytes{hash}, Tips: []string{ni.LatestMilestone}}
	resp, err := api.GetInclusionStates(anr)
	if renderIfError(w, err) {
		return
	}
	if len(resp.States) == 0 {
		renderIfError(w, errors.New("transaction is not found while GetInclusionStates"))
		return
	}
	err = tmpl.ExecuteTemplate(w, "tx.tpl", struct {
		Analyze   bool
		Error     string
		Server    string
		Hash      giota.Trytes
		Trytes    giota.Trytes
		Tx        giota.Transaction
		Confirmed bool
	}{
		false,
		"",
		server,
		hash,
		gt.Trytes[0].Trits().Trytes(),
		gt.Trytes[0],
		resp.States[0],
	})
	if err != nil {
		log.Print(err)
	}
}
func searchAddress(w http.ResponseWriter, hash giota.Address) {
	server := giota.RandomNode()
	//server := giota.PublicNode[0]
	api := giota.NewAPI(server, nil)

	var ft *giota.FindTransactionsResponse
	var gb *giota.GetBalancesResponse
	var err1, err2 error
	wd := sync.WaitGroup{}
	wd.Add(2)
	go func() {
		ftr := &giota.FindTransactionsRequest{Addresses: []giota.Address{hash}}
		ft, err1 = api.FindTransactions(ftr)
		wd.Done()
	}()
	go func() {
		ftr := &giota.GetBalancesRequest{Addresses: []giota.Address{hash}}
		gb, err2 = api.GetBalances(ftr)
		wd.Done()
	}()
	wd.Wait()
	if renderIfError(w, err1, err2) {
		return
	}

	if len(ft.Hashes) == 0 {
		renderIfError(w, errors.New("transaction is not found while FindTransactions"))
		return
	}
	if len(gb.Balances) == 0 {
		renderIfError(w, errors.New("transaction is not found while GetBalances"))
		return
	}

	err := tmpl.ExecuteTemplate(w, "address.tpl", struct {
		Server  string
		Hash    giota.Address
		Txs     []giota.Trytes
		Balance *giota.GetBalancesResponse
	}{
		server,
		hash,
		ft.Hashes,
		gb,
	})
	if err != nil {
		log.Print(err)
	}
}
func searchBundle(w http.ResponseWriter, hash giota.Trytes) {
	server := giota.RandomNode()
	//server := giota.PublicNode[0]
	api := giota.NewAPI(server, nil)

	var ft *giota.FindTransactionsResponse
	var err1 error
	wd := sync.WaitGroup{}
	wd.Add(1)
	go func() {
		ftr := &giota.FindTransactionsRequest{Bundles: []giota.Trytes{hash}}
		ft, err1 = api.FindTransactions(ftr)
		wd.Done()
	}()
	wd.Wait()
	if renderIfError(w, err1) {
		return
	}
	if len(ft.Hashes) == 0 {
		renderIfError(w, errors.New("transaction is not found while FindTransactions"))
		return
	}

	err := tmpl.ExecuteTemplate(w, "bundle.tpl", struct {
		Server string
		Hash   giota.Trytes
		Txs    []giota.Trytes
	}{
		server,
		hash,
		ft.Hashes,
	})
	if err != nil {
		log.Print(err)
	}
}
