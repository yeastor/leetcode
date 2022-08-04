package main

func maxArea(height []int) int {
	len := len(height)
	//half := int(math.Floor(float64(len / 2)))
	maxArea := 0
	left := 0
	right := len - 1
	min := 0
	moveSide := 1 // 1 - right, 2 - left, 3 - both
	for left != right {
		a := height[left]
		b := height[right]
		min = a
		moveSide = 1
		if b < a {
			min = b
			moveSide = 2
		}
		if a == b {
			moveSide = 3
		}

		area := min * (right - left)
		if area > maxArea {
			maxArea = area
		}

		if moveSide == 1 {
			left = moveRight(height, left, right, a)
		}
		if moveSide == 2 {
			right = moveLeft(height, left, right, b)
		}
		if moveSide == 3 {
			tempLeft := moveRight(height, left, right, a)
			tempRight := moveLeft(height, left, right, b)
			if right-tempLeft > tempRight-left {
				left = tempLeft
			} else {
				right = tempRight
			}
		}
		if left >= right {
			return maxArea
		}

	}
	return maxArea
}

func moveRight(height []int, left int, right int, min int) int {
	for i := left + 1; i <= right; i++ {
		if height[i] > min {
			return i
		}
	}
	return right
}

func moveLeft(height []int, left int, right int, min int) int {
	for i := right - 1; i >= left; i-- {
		if height[i] > min {
			return i
		}
	}
	return left
}

func main() {

}
