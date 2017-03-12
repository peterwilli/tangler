
{{template "header.tpl"}}
        <h1>Address Details</h1>
        <div>
            <table class="wrap-table table table-striped">
                <tbody>
                    <tr>
                        <td>Address</td>
                        <td><a href="/search/?kind=address&hash={{.Hash}}">{{.Hash}}</a></td>
                    </tr>
                    <tr>
                        <td>Balance</td>
                        <td>{{index .Balance.Balances 0}}</td>
                    </tr>
                    <tr>
                        <td>Latest confirmed milestone</td>
                        <td><a href="/search/?kind=transaction&hash={{.Balance.Milestone}}">{{.Balance.Milestone}}</a></td>
                    </tr>
                    <tr>
                        <td>Latest confirmed milestone index</td>
                        <td>{{.Balance.MilestoneIndex}}</td>
                    </tr>
                    <tr>
                        <td>Related Transactions</td>
                        <td>
                        {{range $index,$element:=.Txs}}
                        <a href="/search/?kind=transaction&hash={{$element}}">{{$element}}</a></br>
                        {{end}}
                        </td>
                    </tr>
                </tbody>
            </table>
{{template "footer.tpl"}}
