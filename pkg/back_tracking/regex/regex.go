package regex


// 	text := "baac"
//	p := "ba?c"
var match=false
func rmatch(ti, pj int, p, text string) {
	if match{
		return
	}

	if pj == len(p) {
		if ti == len(text) {
			match=true
		}
		return 
	}

	if p[pj] == '*' {
		for i := 0; i <= len(text)-ti; i++ {
			rmatch(ti+i, pj+1, p, text)
		}
	} else if p[pj] == '?' {
		rmatch(ti, pj+1, p, text)
		rmatch(ti+1, pj+1, p, text) 
	} else if text[ti] == p[pj] { // regular string match
		rmatch(ti+1, pj+1, p, text)
	}
}
