package algorithm

import (
	"mr.jackpot-backend/utility/util"
)





type MinEditDist struct {
	rarr, carr []int
	r, c int
	dp [][]int
}



func (m *MinEditDist) TransferUnicodeToIntArray(str string) []int {

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


func (m *MinEditDist) InitArray() {

	m.r, m.c = len(m.rarr)-1, len(m.carr)-1
}


func (m *MinEditDist) InitDPTable() {
	r, c := m.r, m.c

	m.dp = make([]([]int), r+1)
	for i := 0; i <= r; i++ {
		m.dp[i] = make([]int, c+1)
	}
}


func (m *MinEditDist) FillDPTable() {
	r, c := m.r, m.c
	dp := m.dp
	rarr := m.rarr
	carr := m.carr

	for i := 1; i <= r; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= c; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {

			if rarr[i] == carr[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = util.Min(dp[i-1][j-1], util.Min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
}


func (m *MinEditDist) GetMaxValue() int {
	return m.dp[m.r][m.c]
}


type MinEditDistSolution struct {
	MinEditDist
}

func (m *MinEditDistSolution) SolveMinEditDist(rarr, carr []int) int {
	m.rarr = rarr
	m.carr = carr
	m.InitArray()
	m.InitDPTable()
	m.FillDPTable()
	return m.GetMaxValue()
}

func SolveMinEditDist(tar, cmp []int) int {
	m := MinEditDistSolution{}
	return m.SolveMinEditDist(tar, cmp)
}
