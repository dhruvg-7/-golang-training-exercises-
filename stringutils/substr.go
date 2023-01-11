package stringutils

func OverLapString(a, b string) string {

	if len(a) == 0 || len(b) == 0 {
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
