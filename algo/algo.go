package algo

func KMP(s string, sub string) bool { //алгоритм Кнута — Морриса — Пратта
	p := make([]int, len(sub))
	j, i := 0, 1

	for i < len(sub) { //цикл формирования массива "р"
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

	for i < n { //поиск подстроки (образа) в тексте
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

func BHR(s string, sub string) bool { //алгоритм Бойера — Мура — Хорспула
	m := len(sub)
	set := NewSet[string](m)
	d := make(map[string]int, m)
	//формирование таблицы смещения
	for i := m - 2; i > 0; i-- { //для букв с первой до предпоследней
		if !set.Map[string(sub[i])] {
			d[string(sub[i])] = m - i - 1
			set.Add(string(sub[i]))
		}
	}

	if !set.Map[string(sub[m-1])] { //последняя буква
		d[string(sub[m-1])] = m
	}

	d["*"] = m //остальные

	n := len(s)

	if n >= m {
		i := m - 1

		for i < n {
			k := 0
			for j := m - 1; j > 0; j-- {
				if s[i-k] != sub[j] {
					var off int
					if j == m-1 {
						off = d["*"]
						if d[string(s[i])] == 0 {
							off = d[string(s[i])]
						}
					} else {
						off = d[string(sub[i])]
					}

					i += off
					break
				}
				k++

				if j == 0 {
					return true
				}
			}
		}
	}

	return false
}
