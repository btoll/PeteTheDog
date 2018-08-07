package main

import (
	"container/list"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

var l = list.New()
var hash = sha256.New()

type Block struct {
	LastHash string
	Msg      string
}

func add(block *Block) (*Block, error) {
	if l.Len() > 0 {
		if err := isBlockValid(block); err != nil {
			return nil, err
		}
	}
	l.PushBack(block)
	return block, nil
}

func getHash(lastHash, msg string) string {
	hash.Write([]byte(lastHash + msg))
	defer hash.Reset()
	return hex.EncodeToString(hash.Sum(nil))
}

func isBlockValid(block *Block) error {
	if e := l.Back(); e != nil {
		prev := e.Value.(*Block)
		if block.LastHash != getHash(prev.LastHash, prev.Msg) {
			return errors.New("Hashes don't match!")
		}
	}
	return nil
}

func newBlock(msg string) (*Block, error) {
	hash := "0"
	if e := l.Back(); e != nil {
		prev := e.Value.(*Block)
		hash = getHash(prev.LastHash, prev.Msg)
	}
	return add(&Block{
		Msg:      msg,
		LastHash: hash,
	})
}

func printBlocks() {
	fmt.Fprintf(os.Stdout, "Size %d\n\n", l.Len())
	for e, i := l.Front(), 0; e != nil; e = e.Next() {
		b := e.Value.(*Block)
		fmt.Fprintf(os.Stdout, "Block\t\t%d\n", i)
		fmt.Fprintf(os.Stdout, "Hash\t\t%s\n", getHash(b.LastHash, b.Msg))
		fmt.Fprintf(os.Stdout, "Last Hash\t%s\n", b.LastHash)
		fmt.Fprintf(os.Stdout, "Msg\t\t%s\n\n", b.Msg)
		i++
	}
}

func main() {
	newBlock("genesis block")
	newBlock("huck")
	newBlock("utley")
	newBlock("molly")
	newBlock("pete")
	newBlock("lily")
	newBlock("rupert")
	printBlocks()
}
