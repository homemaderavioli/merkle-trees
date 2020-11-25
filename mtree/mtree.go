package mtree

import (
	"fmt"
)

type InclusionProof struct {
	hashData string
	rootHash string
	hashes   []string
	bitfield []int
}

func ValidateInclusionProof(proof *InclusionProof) bool {
	if proof == nil {
		return false
	}
	rootHash := proof.hashData
	for i := 0; i < len(proof.hashes); i++ {
		bit := proof.bitfield[i]
		if bit == 1 {
			rootHash = ConcatenatedHash(rootHash, proof.hashes[i])
		} else if bit == 0 {
			rootHash = ConcatenatedHash(proof.hashes[i], rootHash)
		}
	}
	if rootHash == proof.rootHash {
		return true
	}
	return false
}

func BuildInclusionProof(rootNode *Node, data string) *InclusionProof {
	if rootNode == nil {
		return nil
	}

	targetHash := HashHex(data)

	stackPathToTarget, targetNode := findPathToTarget(rootNode, targetHash)

	if stackPathToTarget == nil || targetNode == nil {
		return nil
	}

	hashes, bitfield := buildProof(stackPathToTarget, targetNode)

	return &InclusionProof{
		targetHash,
		rootNode.NodeHash,
		hashes,
		bitfield,
	}
}

func findPathToTarget(rootNode *Node, target string) (*nodeStack, *Node) {
	var targetNode *Node = nil

	var nodeStack nodeStack
	node := rootNode
	for {
		if node != nil {
			if node.NodeHash == target {
				targetNode = node
				break
			}
			nodeStack.Put(node)
			node = node.LeftNode
		} else {
			if nodeStack.Empty() {
				break
			}
			temp := nodeStack.Peek()
			nodeStack.Pop()
			if temp.RightNode != nil {
				node = temp.RightNode
			}
		}
	}

	if targetNode == nil {
		return nil, nil
	}

	return &nodeStack, targetNode
}

func buildProof(nodeStack *nodeStack, targetNode *Node) ([]string, []int) {
	hashes := []string{}
	bitfield := []int{}

	tempHash := targetNode.NodeHash

	for !nodeStack.Empty() {
		n := nodeStack.Peek()
		if n.LeftNode.NodeHash == tempHash {
			bitfield = append(bitfield, 1)
			hashes = append(hashes, n.RightNode.NodeHash)
		}
		if n.RightNode.NodeHash == tempHash {
			bitfield = append(bitfield, 0)
			hashes = append(hashes, n.LeftNode.NodeHash)
		}
		tempHash = n.NodeHash
		nodeStack.Pop()
	}
	return hashes, bitfield
}

func NewTree(data []string) *Node {
	leafNodes := make([]*Node, 0)
	for _, s := range data {
		hash := HashHex(s)
		node := &Node{
			hash,
			nil,
			nil,
		}
		leafNodes = append(leafNodes, node)
	}

	return buildTree(leafNodes, len(data))
}

func buildTree(nodes []*Node, length int) *Node {
	if length == 1 {
		return nodes[0]
	}
	if length == 2 {
		return newNode(nodes[0], nodes[1])
	}

	half := 0
	if length%2 == 0 {
		half = length / 2
	} else {
		half = length/2 + 1
	}
	return newNode(buildTree(nodes[:half], half), buildTree(nodes[half:], length-half))
}

func traverseTree(rootNode *Node) {
	if rootNode == nil {
		return
	}
	traverseTree(rootNode.LeftNode)
	traverseTree(rootNode.RightNode)
	fmt.Printf("%s\n", rootNode.NodeHash)
}

type Node struct {
	NodeHash  string
	LeftNode  *Node
	RightNode *Node
}

func newNode(leftNode *Node, rightNode *Node) *Node {
	nodeHash := ConcatenatedHash(leftNode.NodeHash, rightNode.NodeHash)
	node := &Node{
		nodeHash,
		leftNode,
		rightNode,
	}
	return node
}
