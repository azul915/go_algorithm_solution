package chap8

import "fmt"

func Code1() {
	a := []int{4, 3, 12, 7, 11, 1, 9, 8, 14, 6}

	fmt.Println(a[0])
	fmt.Println(a[2])
	a[2] = 5
	fmt.Println(a[2])
}

// Code8.2
type Node struct {
	Next *Node
	Name string
}

// p,q
// p,v,q
func Insert(v *Node, p *Node) {
	v.Next = p.Next
	p.Next = v
}

func Code3() {
	sato := Node{nil, "佐藤"}
	suzuki := Node{nil, "鈴木"}
	sato.Next = &suzuki
	takahashi := Node{nil, "高橋"}
	suzuki.Next = &takahashi
	ito := Node{nil, "伊藤"}
	takahashi.Next = &ito
	watanabe := Node{nil, "渡辺"}
	ito.Next = &watanabe
	yamamoto := Node{nil, "山本"}
	watanabe.Next = &yamamoto
	printLinkedList(&sato)
	fmt.Println("")
	Insert(&Node{nil, "田中"}, &takahashi)
	printLinkedList(&sato)
}

func printLinkedList(head *Node) {
	for head != nil {
		if head.Next == nil {
			fmt.Print(head.Name)
		} else {
			fmt.Printf("%v -> ", head.Name)
		}
		head = head.Next
	}

}

func Code4() {
	sentinel := &Node{}

	init := func() {
		sentinel.Next = sentinel
	}
	init()

	printList := func() {
		cur := sentinel.Next
		for ; cur != sentinel; cur = cur.Next {
			fmt.Printf("%v -> ", cur.Name)
		}
		fmt.Println("")
	}

	insert := func(v *Node) {
		v.Next = sentinel.Next
		sentinel.Next = v
	}

	names := []string{"yamamoto", "watanabe", "ito", "takahashi", "suzuki", "sato"}
	for i, name := range names {
		node := &Node{nil, name}
		insert(node)
		fmt.Printf("step %v:", i)
		printList()
	}
}
