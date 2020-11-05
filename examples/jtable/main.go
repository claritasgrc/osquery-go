package main

import (
	"context"
	"log"
	"os"
	"flag"

	"github.com/kolide/osquery-go"
	"github.com/kolide/osquery-go/plugin/table"
)

func main() {
	socket := flag.String("socket", "", "Path to osquery socket file")
	flag.Parse()
	if *socket == "" {
		log.Fatalf(`Usage: %s --socket SOCKET_PATH`, os.Args[0])
	}

	server, err := osquery.NewExtensionManagerServer("jtable", *socket)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	// Create and register a new table plugin with the server.
	// table.NewPlugin requires the table plugin name,
	// a slice of Columns and a Generate function.
	server.RegisterPlugin(table.NewPlugin("jtable", JtableColumns(), JtableGenerate))
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}

// JtableColumns returns the columns that our table will return.
func JtableColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("path"),
		table.TextColumn("txt"),
		table.TextColumn("hash"),
	}
}

// JtableGenerate will be called whenever the table is queried. It should return
// a full table scan.
func JtableGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	return []map[string]string{
		{
			"path": "/",
			"txt": "helloworld",
			"hash": "asdf",
		},
	}, nil
}
