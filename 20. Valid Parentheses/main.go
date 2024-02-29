package main

func isValid(s string) bool {
	pairs := map[byte]byte{
		byte('}'): byte('{'),
		byte(']'): byte('['),
		byte(')'): byte('('),
	}

	var openPars []byte
	openLen := 0
	for _, parenthes := range []byte(s) {
		if openPar, ok := pairs[(parenthes)]; ok {
			if openLen == 0 {
				return false
			}
			if openPars[openLen-1] == openPar {
				openPars = openPars[:openLen-1]
				openLen--
			} else {
				return false
			}
		} else {
			openPars = append(openPars, parenthes)
			openLen++
		}
	}
	return openLen == 0
}

func main() {

}
