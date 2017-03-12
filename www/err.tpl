{{template "header.tpl"}}
<div class="alert alert-danger" role="alert">
    Error occured.
    Please try again from <a href="/">here</a>.<br/>
    Reason: {{.reason}}
    </div>
{{template "footer.tpl"}}