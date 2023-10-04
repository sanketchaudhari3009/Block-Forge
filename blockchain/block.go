package blockchain

import (
	"time"
)

type Block struct {
	Pos       int
	Data      BookCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

func (b *Block) GenerateHash() {
	GenerateComplexHash(b)
}

func CreateBlock(prevBlock *Block, checkoutItem BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.GenerateHash()

	return block
}

func (b *Block) ValidateHash(hash string) bool {
	b.GenerateHash()
	if b.Hash != hash {
		return false
	}
	return true
}
