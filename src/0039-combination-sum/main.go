package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs func(sum, idx int, list []int)
	dfs = func(sum, idx int, list []int) {
		if sum == target {
			res = append(res, append([]int(nil), list...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum + candidates[i] > target {
				continue
			}
			dfs(sum + candidates[i], i, append(list, candidates[i]))
		}
	}
	dfs(0, 0,  []int{})
	return res
}