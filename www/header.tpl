<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>Tangler - Tangle Explore -</title>

    <!-- Bootstrap -->
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <link href="/css/custom.css" rel="stylesheet">

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>

<body>
    <nav class="navbar navbar-inverse navbar-fixed-top">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false"
                    aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
                <a class="navbar-brand" href="/">
                    <span><img src="/img/logo.png" class="logo">Unofficial  Tangle Explorer</span></a>
            </div>
            <div id="navbar" class="navbar-collapse collapse">
              <ul class="nav navbar-nav">
            <li class="active"><a href="/analyze_tx/">Anylze Transaction</a></li>
            </ul>
                <form class="navbar-form navbar-right" action="/search/">
                    <select class="form-control" name="kind">
                        <option>address</option>
                        <option>transaction</option>
                        <option>bundle</option>
                    </select>
                    <div class="form-group">
                        <input type="text" name="hash" placeholder="Enter Hash to Search" class="form-control">
                    </div>
                    <button type="submit" class="btn btn-success">Search</button>
                </form>
            </div>
            <!--/.navbar-collapse -->
        </div>
    </nav>

    <div class="container">