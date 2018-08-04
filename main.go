// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

// Connect simply connects to Elasticsearch, but uses Go modules
// as introduced with Go 1.11.
//
// Example
//
//
//     go run main.go -url=http://127.0.0.1:9200 -sniff=false
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	elastic "github.com/olivere/elastic/v6"
)

func main() {
	var (
		url   = flag.String("url", "http://localhost:9200", "Elasticsearch URL")
		sniff = flag.Bool("sniff", true, "Enable or disable sniffing")
	)
	flag.Parse()
	log.SetFlags(0)

	if *url == "" {
		*url = "http://127.0.0.1:9200"
	}

	fmt.Printf("Go version is %s\n", runtime.Version())

	// Create an Elasticsearch client
	fmt.Print("Connecting to Elasticsearch ")
	client, err := elastic.NewClient(elastic.SetURL(*url), elastic.SetSniff(*sniff))
	if err != nil {
		fmt.Printf("failed: %v", err)
		os.Exit(1)
	}
	fmt.Print("succeeded ")

	version, err := client.ElasticsearchVersion(*url)
	if err != nil {
		fmt.Printf("but couldn't retrieve version: %v", err)
		os.Exit(1)
	}
	fmt.Printf("and we're talking to version %s\n", version)
}
