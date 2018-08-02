# Using Elastic with Go modules

Go 1.11 comes with modules support. This repository is a testbed
for modules support.

In a nutshell:

```
$ go version
go version go1.11beta2 darwin/amd64
$ cat go.mod
module github/elastic-with-go-modules

require github.com/olivere/elastic/v6 v6.1.27
$ go build
go: finding github.com/olivere/elastic/v6 v6.1.27
go: downloading github.com/olivere/elastic/v6 v6.1.27
$ ./elastic-with-go-modules
Connection succeeded
Elasticsearch version 6.3.2
```

## License

MIT-LICENSE. See [LICENSE](http://olivere.mit-license.org/)
or the LICENSE file provided in the repository for details.
