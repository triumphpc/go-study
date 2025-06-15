package main

func subarraySum(nums []int, k int) int {
	sum := 0
	result := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		count = 1
		for j := i + 1; j < len(nums); j++ {
			nums[i] += nums[j]
			if sum <= k {
				count++

				if count > result {
					result = count
				}
			} else {
				break
			}
		}
	}

	return result
}

func subarraySum2(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			sum := 0
			for i := start; i < end; i++ {
				sum += nums[i]
			}
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySum3(nums []int, k int) int {
	result := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		count = 1
		for j := i + 1; j < len(nums); j++ {
			nums[i] += nums[j]
			if nums[i] <= k {
				count++

				if count > result {
					result = count
				}
			} else {
				break
			}
		}
	}

	return result
}
