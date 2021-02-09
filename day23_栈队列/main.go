package main

import "fmt"

type MyQueue struct {
}

/** Initialize your data structure here. */
func Constructor() MyQueue {

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

}

/** Get the front element. */
func (this *MyQueue) Peek() int {

}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

}

/**
 * Your MyQueue object will be instantiated and called as such:

 */

func main() {
	obj := Constructor()
	obj.Push(2)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_)
}
