package util


func TransferUnicodeToIntArray(str string) []int {

	runeStr := []rune(str)
	intArr := []int{0}

	for _, runestr := range runeStr {
		num := int(runestr) - 44032;

		fst := num / 28 / 21
		sec := num / 28 % 21
		trd := num % 28

		if trd == 0 {
			intArr = append(intArr, fst, sec)
		} else {
			intArr = append(intArr, fst, sec, trd)
		}
	}

	return intArr
}