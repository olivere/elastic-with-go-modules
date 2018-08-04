# Using Elastic with Go modules

Go 1.11 comes with modules support. This repository is a testbed
for modules support.

To use this package with Go 1.11 (beta3 as of now), do:

```
$ go get golang.org/dl/go1.11beta3
$ go1.11beta3 download
$ go1.11beta3 version
go version go1.11beta3 darwin/amd64
$ cat go.mod
module github/elastic-with-go-modules

require github.com/olivere/elastic/v6 v6.1.27
$ go1.11beta3 build
go: finding github.com/olivere/elastic/v6 v6.1.27
go: downloading github.com/olivere/elastic/v6 v6.1.27
$ ./elastic-with-go-modules
Go version is go1.11beta3
Connecting to Elasticsearch succeeded and we're talking to version 6.3.2
```

To do the same with Go 1.10.3, do:

```
$ go version
go version go1.10.3 darwin/amd64
$ go get -v ./...
$ go build
$ ./elastic-with-go-modules
Go version is go1.10.3
Connecting to Elasticsearch succeeded and we're talking to version 6.3.2
```

Notice that you *must* use Go 1.10.3 (or Go 1.9.7) because limited module support has only been [introduced in those versions](https://go.googlesource.com/go/+/d4e21288e444d3ffd30d1a0737f15ea3fc3b8ad9). Earlier versions of Go wouldn't find the import path `"github.com/olivere/elastic/v6"`.

Unfortunately, tools like [`dep`](https://github.com/golang/dep) don't support this either as of writing this. Follow [#1962](https://github.com/golang/dep/issues/1962) and [this PR](https://github.com/golang/dep/pull/1963) to watch progress on "minimal module awareness".

## License

MIT-LICENSE. See [LICENSE](http://olivere.mit-license.org/)
or the LICENSE file provided in the repository for details.
