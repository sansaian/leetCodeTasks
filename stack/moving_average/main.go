package main

import "fmt"

func main() {
	ma := Constructor(3)
	fmt.Println(ma.Next(1))
	fmt.Println(ma.Next(10))
	fmt.Println(ma.Next(3))
	fmt.Println(ma.Next(5))
}

type MovingAverage struct {
	sumPrfx    []int
	windowSize int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		sumPrfx:    []int{0},
		windowSize: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.sumPrfx = append(this.sumPrfx, this.sumPrfx[len(this.sumPrfx)-1]+val)
	var start int
	if len(this.sumPrfx) > this.windowSize {
		start = len(this.sumPrfx) - 1 - this.windowSize
	}
	sum := this.sumPrfx[len(this.sumPrfx)-1] - this.sumPrfx[start]
	return float64(sum) / float64(len(this.sumPrfx)-start-1)
}
