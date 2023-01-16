package stringutils

import (
	"fmt"
	"time"
)

func timeCommenter(t time.Duration) {
	if t < 22 {
		fmt.Println("too fast ", t)
	} else {
		fmt.Println("too slow ", t)
	}
}
func OverLapString(a, b string) string {

	t := time.Now()
	defer timeCommenter(time.Since(t))

	if a == "" || b == "" {
		return ""
	}

	ans := ""
	count := 0

	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {

			for j := 0; j < len(b); j++ {
				if a[i+j] != b[j] {
					count = 0
					break
				} else {

					count++
					ans += string(a[i+j])
				}
			}
			if count == len(b) {

				break
			} else {
				ans = ""
			}
		}
	}

	return ans
}
