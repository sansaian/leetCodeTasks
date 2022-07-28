package main

func main() {
	myList := New([]int{1, 2, 3, 4, 5, 6})
	myList.Add(7)
	myList.Add(8)
	myList.Add(9)
	myList.Print()
	myList.Remove(3)
	myList.Remove(1)
	myList.Remove(9)
	myList.Print()
}
