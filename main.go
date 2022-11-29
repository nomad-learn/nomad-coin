package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/nomad-learn/nomad-coin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/home.html"))
		blockchain.GetBlockchain().AddBlock("Second")
		tmpl.Execute(rw, homeData{"Home", blockchain.GetBlockchain().AllBlocks()})
	})

	fmt.Printf("listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
