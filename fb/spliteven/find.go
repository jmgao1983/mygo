package spliteven

func Find(s []int) int {
	if len(s) < 3 {
		return -1
	}
	head := 0
	tail := len(s) - 1
	sum1 := s[head]
	sum2 := s[tail]
	for head < tail {
		if sum1 > sum2 {
			tail -= 1
			sum2 += s[tail]
		} else {
			head += 1
			sum1 += s[head]
		}
	}
	if sum1 == sum2 {
		return head
	}
	return -1
}
