package chap9

import "fmt"

func Code1() {
	const MAX = 10000
	var stack [MAX]int
	var top int

	init := func() {
		top = 0
	}

	isEmpty := func() bool {
		return top == 0
	}

	isFull := func() bool {
		return top == MAX
	}

	push := func(x int) {
		if isFull() {
			fmt.Println("error: stack is already full.")
			return
		}
		stack[top] = x
		top++
	}

	pop := func() int {
		if isEmpty() {
			fmt.Println("error: stack is already empty.")
			return -1
		}
		top--
		return stack[top]
	}

	tostring := func() {
		if isEmpty() {
			fmt.Println("error: stack is already empty.")
			return
		}

		fmt.Print("[")
		for i := 0; i < top; i++ {
			if i < top-1 {
				fmt.Printf("%v ", stack[i])
			} else {
				fmt.Printf("%v", stack[i])
			}
		}
		fmt.Println("]")
	}

	init()

	push(3)
	push(5)
	push(7)
	tostring()
	fmt.Println(pop())
	fmt.Println(pop())
	push(9)
	tostring()
}

func Code2() {
	const MAX = 10000
	var queue [MAX]int
	head, tail := 0, 0

	init := func() {
		head, tail = 0, 0
	}

	isEmpty := func() bool {
		return head == tail
	}
	isFull := func() bool {
		return head == (tail+1)%MAX
	}

	enqueue := func(x int) {
		if isFull() {
			fmt.Println("error: queue is already full.")
			return
		}
		queue[tail] = x
		tail++
		if tail == MAX {
			tail = 0
		}
	}

	dequeue := func() int {
		if isEmpty() {
			fmt.Println("error: queue is already empty.")
			return -1
		}
		res := queue[head]
		head++
		if head == MAX {
			head = 0
		}
		return res
	}

	tostring := func() {
		if isEmpty() {
			fmt.Println("error: queue is already empty.")
			return
		}

		fmt.Print("[")
		for i := head; i < tail; i++ {
			if i < tail-1 {
				fmt.Printf("%v ", queue[i])
			} else {
				fmt.Printf("%v", queue[i])
			}
		}
		fmt.Println("]")
	}

	init()

	enqueue(3)
	enqueue(5)
	enqueue(7)
	tostring()

	fmt.Println(dequeue())
	fmt.Println(dequeue())
	enqueue(9)
	tostring()
}
