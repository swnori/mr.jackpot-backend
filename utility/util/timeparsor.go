package util

import "time"



func TimeParsor(str string) (time.Time, error) {
	return time.Parse(time.RFC1123, str)
}