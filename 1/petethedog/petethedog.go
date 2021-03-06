package petethedog

import (
	"container/list"

	"github.com/btoll/PeteTheDog/1/blockchain"
)

type PeteTheDog struct {
	List *list.List
}

func New() *PeteTheDog {
	return &PeteTheDog{
		List: blockchain.InitLinkedList(),
	}
}

func (p *PeteTheDog) NewBlock(data string) *blockchain.Block {
	block := blockchain.NewBlock(p.List, data)
	return block
}
