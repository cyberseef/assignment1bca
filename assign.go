package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	Nonce    int
	Trans    string
	Prevhash string
	Hash     string
}

func CalculateHash1(hashh string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashh)))
}

func Newblock(n int, t string) *Block {
	s := new(Block)
	s.Nonce = n
	s.Trans = t
	s.Hash = CalculateHash1(strconv.Itoa(n) + t)
	return s
}

type Blocklist struct {
	list []*Block
}

func (ls *Blocklist) AddBlock(n int, t string) *Block {
	st := Newblock(n, t)
	var x = VerifyChain(ls)
	if x == true {
		ls.list = append(ls.list, st)
		CalculateHash(ls)
		fmt.Println("Block Added")
		return st
	} else {
		return nil
	}
}
func ListBlocks(obj *Blocklist) {

	var l = len(obj.list)
	for i := l - 1; i >= 0; i-- {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println("Previous Hash:   ", obj.list[i].Prevhash)
		fmt.Println("Current Hash:   ", obj.list[i].Hash)
		fmt.Println("Nonce:    ", obj.list[i].Nonce)
		fmt.Println("Transaction ID:   ", obj.list[i].Trans)

	}

}

func CalculateHash(stud *Blocklist) {
	var l = len(stud.list)
	for i := l - 1; i >= 0; i-- {
		sum := sha256.Sum256([]byte(stud.list[i].GetString()))
		stud.list[i].Hash = fmt.Sprintf("%x", sum)
		if i < len(stud.list)-1 {
			stud.list[i+1].Prevhash = fmt.Sprintf("%x", sum)
		}
	}
}
func (s *Block) GetString() string {

	var r = ""
	r += strconv.Itoa(s.Nonce)
	r += s.Trans + s.Prevhash
	return r
}

func VerifyChain(stud *Blocklist) bool {
	var st = ""
	var l = len(stud.list)
	for i := l - 1; i >= 0; i-- {
		sum := sha256.Sum256([]byte(stud.list[i].GetString()))
		st = fmt.Sprintf("%x", sum)

		if st != stud.list[i].Hash {
			fmt.Printf("Block is Tempered at Block #. %d\n", i)
			return false
		}
	}
	fmt.Println("Blockchain isValid")
	return true
}
func ChangeBlock(stud *Blocklist, n int, t string) {
	var l = len(stud.list)
	for i := l - 1; i >= 0; i-- {
		if n == stud.list[i].Nonce {

			stud.list[i].Trans = t
			fmt.Println("block Changed")
			return
		}
	}
	fmt.Println("Block Not exist")
}

func main() {
	blockchain := new(Blocklist)
	nonn := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	blockchain.AddBlock(nonn[0], "welcome to the blockchain")

	blockchain.AddBlock(nonn[1], "just testing")

	blockchain.AddBlock(nonn[2], "this is the third block")

	ListBlocks(blockchain)
	fmt.Println("Changing Block with Nonce: 3")
	ChangeBlock(blockchain, nonn[2], "Kamran to Meraj")

	fmt.Println("Adding New Block ")
	blockchain.AddBlock(nonn[3], "New Transaction")

	fmt.Println("Verifying Chain ")
	VerifyChain(blockchain)
	fmt.Println("Recalculate Hashes")
	CalculateHash(blockchain)
	fmt.Println("Verifying Blockchain")
	VerifyChain(blockchain)

}
