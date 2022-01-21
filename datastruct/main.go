package main

import "github.com/leetCodeTasks/datastruct/list"

func main() {
	exampleList := list.New([]int{1, 2, 3, 4})
	exampleList.Print()
	exampleList.Add(5)
	exampleList.Print()
	exampleList.Remove(5)
	exampleList.Print()
	exampleList.Remove(1)
	exampleList.Print()
}
