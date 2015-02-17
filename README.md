# unlinked.in [![GoDoc](https://godoc.org/github.com/internavenue/unlinked.in?status.svg)](https://godoc.org/github.com/internavenue/unlinked.in) [![Build Status](https://travis-ci.org/internavenue/unlinked.in.svg?branch=master)](https://travis-ci.org/internavenue/unlinked.in)
Unlinked.in help to retrieve and convert LinkedIn profile to:
  * [HR-XML JSON Lightweight Recruiting](http://www.hropenstandards.org/) Candidate
  * [JSON Resume](https://jsonresume.org/)
  * [Europass](http://interop.europass.cedefop.europa.eu/) LearnerInfo

# Build
  * Install Go
  * In your terminal
```
go get -u github.com/internavenue/unlinked.in
cd $GOPATH/src/github.com/internavenue/unlinked.in
go build
```
A "unlinked.in" binary file will be generate.

# Config & Run
  * Creat a "cfg.json" with content can be found at "[cfg.json.example](https://github.com/internavenue/unlinked.in/blob/master/cfg.json.example)"
  * In your terminal
```
cd $GOPATH/src/github.com/internavenue/unlinked.in
./unlinked.in -cfgfile='/path/to/cfg.json'
```
  * Flags and Config options:
    * Flags:
      * cfgfile: full path to the JSON config file.
    * Config options:
      * HTTPAddress: the network address will listening on. (eg: localhost:80)
      * SiteURL:  the real-full url that "unlinked.in" will serve - without trailing slash. (eg: http://localhost)
      * TemplateDir: path to UI template folder (a working example included in the project)
      * SessionKey: A secure and random string for cookie encrytion.
      * APIKey: LinkedIn Application API Key (can be create here https://www.linkedin.com/secure/developer)
      * SecretKey: LinkedIn Application Secret Key

# Licence

Released under GPLv3 