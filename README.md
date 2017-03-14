[![Build Status](https://travis-ci.org/utamaro/tangler.svg?branch=master)](https://travis-ci.org/utamaro/tangler)
[![GoDoc](https://godoc.org/github.com/utamaro/tangler?status.svg)](https://godoc.org/github.com/utamaro/tangler)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/utamaro/tangler/master/LICENSE)


# tangler

## Overview

Yet another iota tangle explorer.

Visit [here](https://tangler.herokuapp.com/) for demo.

## Requirements

This requires

* git
* go 1.7+


## Installation

    $ mkdir tmp
    $ cd tmp
    $ mkdir src
    $ mkdir bin
    $ mkdir pkg
    $ exoprt GOPATH=`pwd`
    $ go get github.com/utamaro/tangler
    $ go build
    $ ./tangler

# Using Uour IRI Server

By default tangler uses public nodes listed at [here](http://iotasupport.com/lightwallet.shtml).
If you want to use your own IRI server,  you would need to change every lines in main.go

```
server := giota.RandomNode()
```

to

```
server:="http://localhost:14265"
```

or something.

# Contribution
Improvements to the codebase and pull requests are encouraged.


