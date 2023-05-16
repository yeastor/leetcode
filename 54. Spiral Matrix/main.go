package main

func spiralOrder(matrix [][]int) []int {

	m := len(matrix[0])
	n := len(matrix)
	incX := 1
	incY := 0
	if m == 1 {
		incX = 0
		incY = 1
	}
	doneX := 0
	doneY := 0
	totalEl := n * m
	res := make([]int, totalEl)
	cur := 0
	i := 0
	j := 0
	res[cur] = matrix[i][j]
	cur++
	m--
	n--
	for k := totalEl - 1; k > 0; k-- {
		i += incY
		j += incX
		res[cur] = matrix[i][j]
		cur++
		if incX > 0 && j == m {
			m--
			incX = 0
			incY = 1
		} else if incY > 0 && i == n {
			incX = -1
			incY = 0
			doneY++
		} else if incX < 0 && j == doneX {
			incX = 0
			incY = -1
			doneX++
		} else if incY < 0 && i == doneY {
			n--
			incX = 1
			incY = 0
		}
	}
	return res
}

func main() {

}
