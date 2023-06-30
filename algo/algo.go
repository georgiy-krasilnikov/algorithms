package algo

func KMP(s string, sub string) bool {
	p := make([]int, len(sub))
	j, i := 0, 1

	for i < len(sub) {
		if sub[j] == sub[i] {
			p[i] = j + 1
			i++
			j++
		} else {
			if j == 0 {
				p[i] = 0
				i++
			} else {
				j = p[j-1]
			}
		}
	}

	m, n := len(sub), len(s)
	i, j = 0, 0

	for i < n {
		if s[i] == sub[j] {
			i++
			j++
			if j == m {
				return true
			}
		} else {
			if j > 0 {
				j = p[j-1]
			} else {
				i++
			}
		}
	}

	return false
}
