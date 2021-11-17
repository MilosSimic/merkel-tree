package main

import (
	"fmt"
	"crypto/sha1"
	"encoding/hex"
)

const (
	START = 0
	STEP = 2
)

type MerkelRoot struct {
	root *Node
}

func (mr *MerkelRoot) String() string{
	return mr.root.String()
}

type Node struct {
	data []byte
	left *Node
	right *Node
}

func (n *Node) String() string {
	return hex.EncodeToString(n.data[:])
}

func process(parts []Node) []Node {
	nodes := []Node{}
	if len(parts) % 2 != 0 {
		parts = append(parts, Node{data:[]byte{}}) // IF can't create binary tree, add empty elems to make it work :D
	}

	start := START
	step := STEP
	for true {
		if step <= len(parts){
			elems := parts[start:step]
			l := elems[0]
			r:= elems[1]
			d := sha1.Sum(append(l.data[:], r.data[:]...))
			nodes = append(nodes, Node{left: &l, right: &r, data: d[:]})
			start = step
			step = step + STEP
		}else{
			break
		}
	}

	if len(nodes) == 1 {
		return nodes
	}else{
		return process(nodes)
	}
}

func NewMerkelTree(parts []Node) *MerkelRoot {
	elems := process(parts)
	return &MerkelRoot{root: &elems[0]}
}

func main(){
	nodes := []Node{
		Node{data:[]byte("a")},
		Node{data:[]byte("b")},
		Node{data:[]byte("c")},
		Node{data:[]byte("d")},
		Node{data:[]byte("e")},
		// Node{data:"f"},
	}

	r := NewMerkelTree(nodes)
	fmt.Println(r)
}
