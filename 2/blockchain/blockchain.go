package blockchain

import (
	"container/list"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Blockchain interface {
	AddTransaction()
	NewBlock(string) *Block
}

type Hasher interface {
	Hash()
}

type Miner interface {
	Mine() (string, error)
}

type Block struct {
	Index        int            `json:"index"`
	Timestamp    int64          `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
	Proof        int            `json:"proof"`
	LastHash     string         `json:"last_hash"`
}

type Transaction struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int    `json:"amount"`
}

func incr() func() int {
	i := -1
	return func() int {
		i++
		return i
	}
}

func mine(msg string) (int, error) {
	var encoded string
	var hashCollisionPrefix string

	hash := sha256.New()
	k := 5

	for i := 0; i < k; i++ {
		hashCollisionPrefix += "0"
	}

	nonce := 0

	for {
		_, err := hash.Write([]byte(msg + strconv.Itoa(nonce)))
		if err != nil {
			return -1, errors.New(fmt.Sprintf("Error: %s could not be hashed!", msg+strconv.Itoa(nonce)))
		}

		encoded = hex.EncodeToString(hash.Sum(nil))
		if encoded[:k] == hashCollisionPrefix {
			return nonce, nil
		}

		hash.Reset()
		nonce += 1
	}

	return -1, nil
}

var counter = incr()

func InitLinkedList() *list.List {
	head := &Block{
		Index:        counter(),
		Timestamp:    time.Now().Unix(),
		Transactions: []*Transaction{},
		Proof:        0,
	}
	l := list.New()
	l.PushFront(head)
	return l
}

func NewBlock(l *list.List, msg string) *Block {
	hashed, err := mine(msg)
	if err != nil {
		// TODO
		hashed = -1
	}
	newBlock := &Block{
		Index:        counter(),
		Timestamp:    time.Now().Unix(),
		Transactions: []*Transaction{},
		Proof:        hashed,
	}
	l.PushBack(newBlock)
	return newBlock
}
