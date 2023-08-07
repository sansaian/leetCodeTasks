package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
}

type stats struct {
	win  int
	lose int
}

func findWinners(matches [][]int) [][]int {

	stat := make(map[int]stats)
	for _, game := range matches {
		stat[game[0]] = stats{
			win:  stat[game[0]].win + 1,
			lose: stat[game[0]].lose,
		}

		stat[game[1]] = stats{
			win:  stat[game[1]].win,
			lose: stat[game[1]].lose + 1,
		}
	}

	result := [][]int{{}, {}}
	for k, v := range stat {
		if v.lose == 0 {
			result[0] = append(result[0], k)
			continue
		}
		if v.lose == 1 {
			result[1] = append(result[1], k)
		}
	}
	sort.Ints(result[0])
	sort.Ints(result[1])
	return result
}
