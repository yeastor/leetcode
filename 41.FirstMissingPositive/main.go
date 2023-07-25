package main

func firstMissingPositive(n []int) int {

	arrLen := len(n) + 1
	newArr := make([]bool, arrLen)
	for _, i := range n {
		if i > 0 && i <= arrLen {
			newArr[i-1] = true
		}
	}

	for i := 0; i <= arrLen; i++ {
		if newArr[i] == false {
			return i + 1
		}
	}

	return arrLen
}

func main() {

}
