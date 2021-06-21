gocover-cobertura (fork)
======================
Forked because it was not possible to import and use the Cobertura Coverage definitions as a library.

- This fork moves the core functionalities of this tool to two packages `cobertura` and `gocover`.
- Added a `go.mod` file to work with Go modules.
- I also changed the `class.NumLines()` function to count the list of lines in the class, instead of the length of all of its methods. This allows it to also work for parsing Cobertura coverage files generated with `pytest`, while still working for Golang coverages, since the resulting Cobertura compatible XML already contains every source line as both lines and methods.

> _Fork author: [Bart van Oort (bvobart)](https://github.com/bvobart)_

Original ReadMe 
================

[![Build Status](https://travis-ci.org/t-yuki/gocover-cobertura.svg?branch=master)](https://travis-ci.org/t-yuki/gocover-cobertura)
[![Coverage Status](https://coveralls.io/repos/github/t-yuki/gocover-cobertura/badge.svg?branch=master)](https://coveralls.io/github/t-yuki/gocover-cobertura?branch=master)

go tool cover XML (Cobertura) export
====================================

This is a simple helper tool for generating XML output in [Cobertura](http://cobertura.sourceforge.net/) format
for CIs like [Jenkins](https://wiki.jenkins-ci.org/display/JENKINS/Cobertura+Plugin) and others
from [go tool cover](https://code.google.com/p/go.tools/) output.

Installation
------------

Just type the following to install the program and its dependencies:

    $ go get code.google.com/p/go.tools/cmd/cover
    $ go get github.com/t-yuki/gocover-cobertura

Usage
-----

`gocover-cobertura` reads from the standard input:

    $ go test -coverprofile=coverage.txt -covermode count github.com/gorilla/mux
    $ gocover-cobertura < coverage.txt > coverage.xml

Authors
-------

* [Yukinari Toyota (t-yuki)](https://github.com/t-yuki)

Thanks
------

This tool is originated from [gocov-xml](https://github.com/AlekSi/gocov-xml) by [Alexey Palazhchenko (AlekSi)](https://github.com/AlekSi)
