package spliteven

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 20, 1, 17, 2, 4}, 2},
		{[]int{3, 20, 0, 0, 0, 1, 17, 2, 4}, 5},
		{[]int{3, 20, 0, 0, 0, 0, 17, 2, 4}, 5},
		{[]int{3, 20, 1, 17, 2, 4, 0, 0, 0, 0, 0}, 2},
		{[]int{4, 5, 9}, -1},
		{[]int{}, -1},
	}
	for i, c := range cases {
		got := Find(c.in)
		fmt.Printf("%03d: ", i)
		fmt.Print(c.in)
		fmt.Printf("Find() == %d, want %d", got, c.want)
		if got != c.want {
			t.Errorf(" >>Failed!\n")
		} else {
			fmt.Printf(" >>Success!\n")
		}
	}
}
