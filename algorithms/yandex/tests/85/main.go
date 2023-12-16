package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
)

// https://telegra.ph/Golang-zadacha-Kombinaciya-summ-II-11-30
func main() {
	task(os.Stdin, os.Stdout)
}

func task(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(out)
	defer func() {
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
	}()

	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	// todo

}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combos := make([][]int, 0)

	var dfs func([]int, int, []int)
	dfs = func(nums []int, targ int, path []int) {
		if targ < 0 {
			return
		}
		if targ == 0 {
			combos = append(combos, path)
			return
		}

		for i, n := range nums {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			newPath := make([]int, len(path))
			copy(newPath, path)
			dfs(nums[i+1:], targ-n, append(newPath, n))
		}
		return
	}

	dfs(candidates, target, []int{})
	return combos
}
