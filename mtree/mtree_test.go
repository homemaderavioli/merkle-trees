package mtree

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestValidateInclusionProof(t *testing.T) {
	cases := []struct {
		inputTree *Node
		inputData string
		want      bool
	}{
		{
			NewTree([]string{"hello world", "merkle tree", "blockchain"}),
			"hello world",
			true,
		},
		{
			NewTree([]string{"hello world", "merkle tree", "blockchain"}),
			"hello mars",
			false,
		},
	}

	for _, c := range cases {
		proof := BuildInclusionProof(c.inputTree, c.inputData)
		got := ValidateInclusionProof(proof)
		assert.Equal(t, c.want, got)
	}
}

func TestBuildInclusionProof(t *testing.T) {
	rootNode := NewTree([]string{"hello world", "merkle tree", "blockchain"})
	data := "hello world"
	got := BuildInclusionProof(rootNode, data)
	want := &InclusionProof{
		"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		"97675815e25e3f37f26f4783ba740ecaa9b4fa2872069c0fdf1f2c0c1e7f590d",
		[]string{"305d7abb3a8e806f3806e85f7702d51e2f58269a4e4759fa8f7be7facd8e0fd9", "ef7797e13d3a75526946a3bcf00daec9fc9c9c4d51ddc7cc5df888f74dd434d1"},
		[]int{1, 1},
	}
	assert.Equal(t, want, got)
}

func TestNewTree(t *testing.T) {
	cases := []struct {
		input []string
		want  string
	}{
		{[]string{"hello world"}, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{[]string{"hello world", "merkle tree"}, "77780c7c12ffae0ba949febb05181b542b48db7f3603be474afd0de19ff0af40"},
		{[]string{"hello world", "merkle tree", "blockchain"}, "97675815e25e3f37f26f4783ba740ecaa9b4fa2872069c0fdf1f2c0c1e7f590d"},
		//{[]string{"hello world", "merkle tree", "blockchain", "ethereum"}, "cj"},
		//{[]string{"hello world", "merkle tree", "blockchain", "ethereum", "bitcoin"}, "d"},
	}

	for _, c := range cases {
		got := NewTree(c.input)
		//traverseTree(got)
		//fmt.Printf("====\n")
		assert.Equal(t, c.want, got.NodeHash)
	}
}

func TestNewNode(t *testing.T) {
	node1 := &Node{
		"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		nil,
		nil,
	}
	node2 := &Node{
		"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		nil,
		nil,
	}
	want := "bc62d4b80d9e36da29c16c5d4d9f11731f36052c72401a76c23c0fb5a9b74423"
	got := newNode(node1, node2)
	assert.Equal(t, want, got.NodeHash)
}
