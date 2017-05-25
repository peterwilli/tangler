{{template "header.tpl"}}

        <h1>Node Information</h1>
        <div>
            <table class="wrap-table table table-striped">
                <tbody>
                    <tr>
                        <td>Application Name</td>
                        <td>{{.NodeInfo.AppName}}</td>
                    </tr>
                    <tr>
                        <td>Application Version</td>
                        <td>{{.NodeInfo.AppVersion}}</td>
                    </tr>
                    <tr>
                        <td> Latest Milestone</td>
                        <td><a href="/search/?kind=transaction&hash={{.NodeInfo.LatestMilestone}}">{{.NodeInfo.LatestMilestone}}</a></td>
                    </tr>
                    <tr>
                        <td>Lastest Milestone Index</td>
                        <td>{{.NodeInfo.LatestMilestoneIndex}}</td>
                    </tr>
                    <tr>
                        <td> Lastest Solid Subtangle Milestone</td>
                        <td><a href="/search/?kind=transaction&hash={{.NodeInfo.LatestSolidSubtangleMilestone}}">{{.NodeInfo.LatestSolidSubtangleMilestone}}</a></td>
                    </tr>
                    <tr>
                        <td>Lastest Solid Subtangle Milestone Index</td>
                        <td>{{.NodeInfo.LatestSolidSubtangleMilestoneIndex}}</td>
                    </tr>
                    <tr>
                        <td>Number of Neighbors</td>
                        <td>{{.NodeInfo.Neighbors}}</td>
                    </tr>
                    <tr>
                        <td>Time</td>
                        <td>{{localtime .NodeInfo.Time}}</td>
                    </tr>
                    <tr>
                        <td>Number of tips</td>
                        <td>{{.NodeInfo.Tips}}</td>
                    </tr>
                </tbody>
            </table>

            <h1>Transactions To Be Approved</h1>
            <table class="wrap-table table">
                <thead>
                </thead>
                <tbody>
                    <tr>
                        <td>Trunk</td>
                        <td><a href="/search/?kind=transaction&hash={{.Tx.TrunkTransaction}}">{{.Tx.TrunkTransaction}}</a></td>
                    </tr>
                    <tr>
                        <td>Branch</td>
                        <td><a href="/search/?kind=transaction&hash={{.Tx.BranchTransaction}}">{{.Tx.BranchTransaction}}</a></td>
                    </tr>
                </tbody>
            </table>
        </div>

{{template "footer.tpl"}}
