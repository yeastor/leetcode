package main

import "strconv"

func countAndSay(n int) string {
	cur := "1"
	for i := 2; i <= n; i++ {
		cur = say(cur)
	}
	return cur
}

func say(s string) string {
	var countN int
	curS := s[0]
	var res []byte
	for _, sym := range []byte(s) {
		if sym != curS {
			res = append(append(res, strconv.Itoa(countN)[0]), curS)
			countN = 0
			curS = sym
		}
		countN++
	}
	res = append(append(res, strconv.Itoa(countN)[0]), curS)
	return string(res)
}

func main() {

}
