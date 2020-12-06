package negativesequencedegree

// in := []int{6, 1, 5, 4, 3, 2}
// nsd is 5+3+2+1=11,
// detail is (6,1),(6,5),(6,4),(6,3),(6,2),
// (5,4),(5,3),(5,2),
// (4,3),(4,2),
// (3,2)
func nsd(in []int, p, r int) int {
	if p >= r {
		return 0
	}

	q := (p + r) / 2
	// caculate the nsd of the first half part
	k1 := nsd(in, p, q)
	// caculate the nsd of the second half part
	k2 := nsd(in, q+1, r)
	// caculate the nsd between the first half part and the second half part
	k3 := merge(in, p, q, r)
	return k1 + k2 + k3
}

// merge return the nsd between in[p:q+1] an in[q+1:r]
func merge(in []int, p, q, r int) (result int) {
	tmp := make([]int, r-p+1)
	i, j, k := p, q+1, 0
	for i <= q && j <= r {
		if in[i] <= in[j] {
			tmp[k] = in[i]
			i++
		} else {
			result += q - i + 1
			tmp[k] = in[j]
			j++
		}
		k++
	}

	for i <= q {
		tmp[k] = in[i]
		k++
		i++
	}

	for j <= r {
		tmp[k] = in[j]
		k++
		j++
	}

	copy(in[p:r+1], tmp)

	return result
}
