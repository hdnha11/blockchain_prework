package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	genesisBlock := Block{"Genesis Block", time.Now().Unix(), []byte{}, []byte{}}
	genesisBlock.Hash = genesisBlock.calculateHash()

	Blockchain = []Block{genesisBlock}
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	newBlock := Block{data, time.Now().Unix(), []byte{}, []byte{}}
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = newBlock.calculateHash()

	return newBlock
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	lastestBlock := Blockchain[len(Blockchain)-1]
	err := b.isValidBlock(lastestBlock)

	if err != nil {
		return err
	}

	Blockchain = append(Blockchain, b)

	return nil
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}

func (b *Block) isValidBlock(previousBlock Block) error {
	blockHash := b.calculateHash()

	switch {
	case !bytes.Equal(b.Hash, blockHash):
		return fmt.Errorf("Invalid hash: %x %x", b.Hash, blockHash)
	case !bytes.Equal(b.PrevHash, previousBlock.Hash):
		return fmt.Errorf("Invalid previous hash")
	}

	return nil
}
