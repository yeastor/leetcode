package main

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, cur := range nums {
		sum += cur
		if sum < cur {
			sum = cur
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {

}
