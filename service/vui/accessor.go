package vui

import (
	"strings"

	"mr.jackpot-backend/utility/algorithm"
	"mr.jackpot-backend/utility/util"
)




func (v *VUIAccessor) TransferToFloat(value, sz int) float64 {
	return 1 - float64(value) / float64(sz)
}

func (v *VUIAccessor) GetTargetCandidate(message string) []string {
	candidate := make([]string, 0)

	candidate = append(candidate, message)
	candidate = append(candidate, strings.Split(message, " ")...)
	candidate = append(candidate, strings.ReplaceAll(message, " ", ""))

	return candidate
}

func (v *VUIAccessor) GetTargetId(message string, targetList []string) int {

	for i := range targetList {
		targetList[i] = strings.ReplaceAll(targetList[i], " ", "")
	}

	sz := len(targetList)
	point := make([]float64, sz)
	for _, candidate := range v.GetTargetCandidate(message) {
		for idx, target := range targetList {
			candidateArr := util.TransferUnicodeToIntArray(candidate)
			targetArr := util.TransferUnicodeToIntArray(target)
			p := algorithm.SolveMinEditDist(candidateArr, targetArr)

			point[idx] = util.Max(point[idx], v.TransferToFloat(p, len(candidateArr) + len(targetArr)))
		}
	}

	idx := util.MaxIdx(point)

	if point[idx] < v.threshold {
		return -1
	} else {
		return idx
	}
}



