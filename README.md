# Kli

kli is Klivo's official CLI and it uses the Go Library [cobra](https://github.com/spf13/cobra) to simplify the CLI workflow.

Usage
-----

``` sh
$ kli new-member --data [data.txt]
#=> :klivo: member created with id: 35d413bf-e107-4a33-9449-372c3ce9735f

$ kli batch-members --data [member-batch.csv]
#=> :klivo: members created. To se full list of members, click here

$ kli block-member --id [id]
#=> :klivo: member blocked

```

See [usage examples](https://kli.github.com/#developer) or the [full reference
documentation](https://kli.github.com/kli.html) to see all available commands
and flags.

Installation
------------

#### Standalone

`kli` can be easily installed as an executable. Download the [latest
binary][latest] for your system and put it anywhere in your executable path.

#### Source

Prerequisites for building from source are:

* `make`
* [Go](https://golang.org/doc/install)


This assumes support for [Go 1.11+
modules](https://github.com/golang/go/wiki/Modules). If you are building on an
older version of Go, you will need to clone the repository into
`$GOPATH/src/github.com/github/kli`.

# Contributing

1. Fork it
2. [Setup SSH key](https://help.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh)
2. Download your fork to your PC (`git clone git@github.com:klivohealth/kli.git && cd kli`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Make changes and add them (`git add .`)
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create new pull request

Meta
----

* Bugs: <https://github.com/klivohealth/kli/issues>
* Authors: <https://github.com/klivohealth/kli/graphs/contributors>
