// package zip has a function which gives a slice of tuples
package zip

import "fmt"

type intTuple struct {
	a, b int
}

func zip(a, b []int) ([]intTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: argumens must be of same length")
	}

	r := make([]intTuple, len(a), len(a))

	for i, e := range a {
		r[i] = intTuple{e, b[i]}
	}
	return r, nil
}

// func main() {
// 	d := []int{1, 2, 3}
// 	e := []int{4, 5, 6}
// 	c, _ := zip(d, e)
// 	for _, j := range c {
// 		it := j
// 		fmt.Println(it.a + it.b)
// 		}
// 	}
