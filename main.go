package main

import (
	"github.com/wealdtech/go-merkletree"
	"fmt"
	"encoding/hex"
	"sort"
	"github.com/wealdtech/go-merkletree/keccak256"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	whitelist := []string{
		"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		"0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
	}
	sort.Sort(sort.StringSlice(whitelist))
	data := [][]byte{}
	for i := 0; i < len(whitelist); i++ {
		addr := common.HexToAddress(whitelist[i])
		data = append(data, addr.Bytes())
	}

	tree, err := merkletree.NewUsing(data, keccak256.New(), nil)
	if err != nil {
		panic(err)
	}

	root := tree.Root()
	fmt.Println("root: 0x" + hex.EncodeToString(root))
	fmt.Println("tree str: 0x" + tree.String())
}
