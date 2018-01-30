package ranges

import (
	"bytes"
	"strconv"
)

// Range will convert a list of integers into string grouping together
// consecutive digits together.
func ToRange(a ...int) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(a[0])
	}

	buf := bytes.Buffer{}
	for i := 0; i < len(a); i++ {
		buf.WriteString(strconv.Itoa(a[i]))
		start := a[i]
		for i < len(a)-1 && a[i]+1 == a[i+1] {
			i++
		}
		if start != a[i] {
			buf.WriteString("-" + strconv.Itoa(a[i]))
		}

		if i != len(a)-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
