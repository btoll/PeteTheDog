package petethedog

import (
	"container/list"

	"github.com/btoll/PeteTheDog/1/blockchain"
)

type PeteTheDog struct {
	CurrentTransactions []blockchain.Transaction `json:"current_transactions"`
	List                *list.List
}

func New() *PeteTheDog {
	return &PeteTheDog{
		CurrentTransactions: []blockchain.Transaction{},
		List:                blockchain.InitLinkedList(),
	}
}

func (p *PeteTheDog) AddTransaction() {
}

func (p *PeteTheDog) NewBlock(data string) *blockchain.Block {
	block := blockchain.NewBlock(p.List, data)
	return block
}
