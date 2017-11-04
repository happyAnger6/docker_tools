package main

import (
	"github.com/opencontainers/go-digest"
	"os"
	"fmt"
	"strings"
	_ "crypto/sha256"
)

type DiffID digest.Digest
type ChainID digest.Digest

func CreateChainID(dgsts []DiffID) ChainID {
	return createChainIDFromParent("", dgsts...)
}

func createChainIDFromParent(parent ChainID, dgsts ...DiffID) ChainID {
	if len(dgsts) == 0 {
		return parent
	}
	if parent == "" {
		return createChainIDFromParent(ChainID(dgsts[0]), dgsts[1:]...)
	}
	// H = "H(n-1) SHA256(n)"
	dgst := digest.FromBytes([]byte(string(parent) + " " + string(dgsts[0])))
	return createChainIDFromParent(ChainID(dgst), dgsts[1:]...)
}

func main(){
	if len(os.Args) == 0 {
		fmt.Printf("usage:%s <... ID>\r\n", os.Args[0])
		os.Exit(-1)
	}
	diffIDs := make([]DiffID, 1)
	for _, id := range os.Args[1:]{
		hsh := strings.TrimSpace(string(digest.Canonical) + ":" + id)
		dig, err := digest.Parse(hsh)
		if err != nil {
			fmt.Printf("invalid digest:%v :%v\r\n", hsh, err)
			os.Exit(-1)
		}
		fmt.Printf("append digest:%v\r\n", id)
		diffIDs = append(diffIDs, (DiffID(dig)))
	}
	chainID := CreateChainID(diffIDs)
	fmt.Printf("chainID: %v\n", chainID)
}
