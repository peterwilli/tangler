{{template "header.tpl"}}  
        <h1>Bundle Details</h1>
        <div>
            <table class="wrap-table table table-striped">
                <tbody>
                    <tr>
                        <td>Bundle Hash</td>
                        <td><a href="/search/?kind=bundle&hash={{.Hash}}">{{.Hash}}</a></td>
                    </tr>
                    <tr>
                        <td>Transactions in this bundle</td>
                        <td>
                        {{range $index,$element:=.Txs}}
                        <a href="/search/?kind=transaction&hash={{$element}}">{{$element}}</a></br>
                        {{end}}
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
{{template "footer.tpl"}}  
