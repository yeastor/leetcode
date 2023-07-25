package main

func trap(height []int) int {
	lenght := len(height)
	if lenght == 2 {
		return 0
	}

	var min, move, cur int
	if height[0] >= height[lenght-1] {
		min = height[lenght-1]
		move = 2
		cur = lenght - 2
	} else {
		min = height[0]
		move = 1
		cur = 1
	}

	curl := 0

	currR := lenght - 1
	res := 0

	for i := 1; i <= lenght-2; i++ {
		if height[cur] < min {
			res += min - height[cur]
			if move == 1 {
				cur++
			} else {
				cur--
			}
		} else if move == 1 {
			if height[cur] >= height[currR] {
				move = 2
				min = height[currR]
				curl = cur
				cur = currR - 1
			} else if height[cur] >= height[curl] {
				min = height[cur]
				curl = cur
				cur++
			} else {
				cur++
			}
		} else if move == 2 {
			if height[cur] >= height[curl] {
				move = 1
				min = height[curl]
				currR = cur
				cur = curl + 1
			} else if height[cur] >= height[currR] {
				min = height[cur]
				currR = cur
				cur--
			} else {
				cur--
			}
		}

	}

	return res
}

func main() {

}
