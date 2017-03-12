{{template "header.tpl"}}
{{if .Analyze}}
        <form action="/analyze_tx/" method="post">
            <div class="form-group">
                <label for="exampleInputEmail1">Trytes</label>
                <textarea class="form-control" rows="10" name="trytes">{{.Trytes}}</textarea>
            </div>
            <button type="submit" class="btn btn-default">Analyze</button>
        </form>
{{end}}
{{if eq .Error ""}}
        <h1>Transaction Details</h1>
        <div>
            <table class="wrap-table table table-striped">
                <tbody>
                    <tr>
                        <td>Transaction Hash</td>
                        <td><a href="/search/?kind=transaction&hash={{.Hash}}">{{.Hash}}</a></td>
                    </tr>
                    {{if not .Analyze}}
                    <tr>
                        <td>Confirmed</td>
                        <td>{{.Confirmed}}</td>
                    </tr>
                    {{end}}
                    <tr>
                    </tr>
                    <td>Address</td>
                        <td><a href="/search/?kind=address&hash={{.Tx.Address}}">{{.Tx.Address}}</a></td>
                    </tr>
                    <tr>
                        <td>Value</td>
                        <td>{{.Tx.Value}}</td>
                    </tr>
                    <tr>
                        <td>Tag</td>
                        <td>{{.Tx.Tag}}</td>
                    </tr>
                    <td>Timestamp</td>
                    <td>{{localtime .Tx.Timestamp.Unix}}</td>
                    </tr>
                    <tr>
                        <td>Current Index in Bundle</td>
                        <td>{{.Tx.CurrentIndex}}</td>
                    </tr>
                    <tr>
                        <td>Last Index of Bundle</td>
                        <td>{{.Tx.LastIndex}}</td>
                    </tr>
                    <tr>
                        <td>Trunk Transaction Hash</td>
                        <td><a href="/search/?kind=transaction&hash={{.Tx.TrunkTransaction}}">{{.Tx.TrunkTransaction}}</a></td>
                   </tr>
                    <tr>
                        <td>Branch Transaction Hash</td>
                        <td><a href="/search/?kind=transaction&hash={{.Tx.BranchTransaction}}">{{.Tx.BranchTransaction}}</a></td>
                   </tr>
                    <tr>
                        <td>Bundle Hash</td>
                        <td><a href="/search/?kind=bundle&hash={{.Tx.Bundle}}">{{.Tx.Bundle}}</a></td>
                    </tr>
                    <tr>
                        <td>Nonce</td>
                        <td>{{.Tx.Nonce}}</td>
                    </tr>
                    <tr data-toggle="collapse" data-target="#sig" class="clickable">
                        <td>Message or Signature (click to show)</td>
                        <td>
                            <div class="accordion-body collapse" id="sig">
                                {{.Tx.SignatureMessageFragment}}
                        </td>
                    </tr>
                    <tr data-toggle="collapse" data-target="#trytes" class="clickable">
                        <td>Raw trytes (click to show)</td>
                        <td>
                            <div class="accordion-body collapse" id="trytes">
                                {{.Trytes}}
                        </td>
                    </tr>
                </tbody>
            </table>
     {{else}}
     <div class="alert alert-danger top-margin" role="alert">
            {{.Error}}
            </div>
    {{end}}
 {{template "footer.tpl"}}