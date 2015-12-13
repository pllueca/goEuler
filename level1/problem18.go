package main

import "fmt"

type Tree struct {
	v      int
	fl, fr *Tree
}

func (t *Tree) addChild(v int, d bool) {
	child := Tree{v, nil, nil}
	// left child
	if d {
		t.fl = &child
	} else {
		t.fr = &child
	}
}

func (t *Tree) innerPrint() {
	if t.fl != nil {
		t.fl.innerPrint()
	}
	fmt.Print(t.v, " ")
	if t.fr != nil {
		t.fr.innerPrint()
	}
}

func (t *Tree) Print() {
	t.innerPrint()
	fmt.Print("\n")
}

func (t *Tree) printLeaves() {
	b := 0
	if t.fl != nil {
		t.fl.printLeaves()
		b++
	}
	if t.fr != nil {
		t.fr.printLeaves()
		b++
	}
	if b == 0 {
		fmt.Println(t.v)
	}
}

func binTree(depth int) *Tree {
	i := 0
	t := Tree{v: i, nil, nil}

}

func main() {
	t := Tree{1, nil, nil}
	t.addChild(2, true)
	t.addChild(3, false)
	t.Print()
	t.printLeaves()
}
