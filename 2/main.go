package main

import (
	"fmt"
	"net/http"

	"github.com/btoll/PeteTheDog/1/blockchain"
	"github.com/btoll/PeteTheDog/1/petethedog"
)

var PORT = "12345"

// Create the blockchain.
var pete = petethedog.New()

func entryHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Oh, Pete, the Great and Powerful!")
}

func listHandler(w http.ResponseWriter, req *http.Request) {
	list := pete.List
	fmt.Fprintf(w, "Size %d\n\n", list.Len())
	i := 0
	for e := list.Front(); e != nil; e = e.Next() {
		b := e.Value.(*blockchain.Block)
		fmt.Fprintf(w, "Block\t\t%d\n", i)
		//		fmt.Fprintf(w, "Hash\t\t%s\n", getHash(b.LastHash, b.Msg))
		fmt.Fprintf(w, "Last Hash\t%s\n", b.LastHash)
		fmt.Fprintf(w, "Proof\t\t%d\n\n", b.Proof)
		i++
	}
}

// This expects a POST:
//
//      curl -X POST localhost:12345/newblock --data "msg=helloworld"
//
func newBlockHandler(w http.ResponseWriter, req *http.Request) {
	msg := req.FormValue("msg")
	block := pete.NewBlock(msg)
	fmt.Fprintf(w, "%d\n", block.Proof)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", entryHandler)
	mux.HandleFunc("/list", listHandler)
	mux.HandleFunc("/newblock", newBlockHandler)

	fmt.Printf("Server listening on port %s\n", PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		panic(err)
	}
}
