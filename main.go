package main

import (
	mtree "main/mtree"
)

func main() {

	data := []string{"hello world", "merkle tree", "blockchain"}
	merkleTree := mtree.NewTree(data)

}
