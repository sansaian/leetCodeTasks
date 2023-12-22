package main

func main() {

}

type StockSpanner struct {
	stack    []int
	ansStack []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	ans := 1
	for len(this.stack) != 0 && this.stack[len(this.stack)-1] <= price {
		this.stack = this.stack[:len(this.stack)-1]
		ans += this.ansStack[len(this.ansStack)-1]
		this.ansStack = this.ansStack[:len(this.ansStack)-1]
	}
	this.stack = append(this.stack, price)
	this.ansStack = append(this.ansStack, ans)
	return ans
}
