package ranges

import (
	"fmt"
	"testing"
)

func TestToRange(t *testing.T) {
	tt := []struct {
		in  []int
		out string
	}{
		{[]int{}, ""},
		{[]int{1}, "1"},
		{[]int{999}, "999"},
		{[]int{11, 12, 13, 15, 16, 17, 18, 19, 20, 50, 51, 52, 998, 999}, "11-13,15-20,50-52,998-999"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			if out := ToRange(tc.in...); out != tc.out {
				t.Errorf("unexpected results (want: %q, got %q)", tc.out, out)
			}
		})
	}
}
