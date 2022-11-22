package util


func Max[T int | int16 | int32 | int64 | float32 | float64] (a,b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T int | int16 | int32 | int64 | float32 | float64] (a,b T) T {
	if a > b {
		return b
	}
	return a
}

func MaxIdx[T int | int16 | int32 | int64 | float32 | float64] (array []T) int {
	var (
		MaxIdx int = 0;
		MaxValue T = 0
	)

	for i, e := range array {
		if MaxValue < e {
			MaxValue = e
			MaxIdx = i
		}
	}

	return MaxIdx
}

func IntAll(slice []int32) []int {
	ret := make([]int, 0, len(slice))
	for _, v := range slice {
		ret = append(ret, int(v))
	}
	return ret
}
