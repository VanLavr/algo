package main

import (
	"fmt"
	"log"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuxyzABCDEFGHIJKLMNOPQRSTUXYZ"

type Node struct {
	Name     string
	Value    int
	Children []*Node
}

func main() {
	var Root Node
	Root.Name = GenerateRandomName(3)
	RandomTreeGenerator(&Root)
	PrintTree(&Root, false, true)
	fmt.Println("Weight of this tree is:", DepthSearch(&Root, true))
}

func DepthSearch(node *Node, isRoot bool) (sum int) {
	if isRoot {
		sum += node.Value
		fmt.Printf("[%d] ", node.Value)
	}

	for i := 0; i < len(node.Children); i++ {
		fmt.Printf("[%d] ", node.Children[i].Value)
		sum += node.Children[i].Value
		sum += DepthSearch(node.Children[i], false)
	}
	return
}

func PrintTree(node *Node, isRoot bool, metaRoot bool) {
	printer := func(Children []*Node) string {
		var result string
		for _, child := range Children {
			result += fmt.Sprintf("[%s: %d] ", (*child).Name, (*child).Value)
		}

		return result
	}

	if metaRoot {
		fmt.Printf("Root: [%s] its children: %s", node.Name, printer(node.Children))
		fmt.Println()
	} else if isRoot {
		fmt.Printf("Node: [%s] its children: %s", node.Name, printer(node.Children))
		fmt.Println()
	}

	for i := 0; i < len(node.Children); i++ {
		PrintTree(node.Children[i], true, false)
	}
}

// random tree generator
func RandomTreeGenerator(root *Node) {
	log.Println("generating...")
	numberOfChildren := GenerateChildrenNumber()
	if numberOfChildren == 0 {
		return
	}

	for i := 0; i < numberOfChildren; i++ {
		node := new(Node)
		node.Value = rand.Intn(100)
		node.Name = GenerateRandomName(numberOfChildren + 3)
		RandomTreeGenerator(node)
		root.Children = append(root.Children, node)
	}
}

func GenerateRandomName(length int) string {
	word := make([]byte, length)
	for i := range word {
		word[i] = letters[rand.Intn(len(letters))]
	}

	return string(word)
}

func GenerateChildrenNumber() int {
	numberOfChildren := rand.Intn(5)
	numberOfChildren++
	if rand.Intn(2) == 0 {
		numberOfChildren = 0
	}

	return numberOfChildren
}
