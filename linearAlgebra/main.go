//
package linearAlgebra

import (
	"fmt"
	"math"
)

// Creating a Vector type using slice
type Vector float64

// Add adds two Vector and returns a Vector
// It adds the corresponding element of both Vectors
// to each other
func Add(v, w []Vector) ([]Vector, error) {
	var x []Vector
	if len(v) != len(w) {
		return nil, fmt.Errorf("error: Vectors must be of same length")
	} else {
		for i, _ := range v {
			for j, _ := range w {
				if i == j {
					x = append(x, v[i]+w[j])
				}
			}
		}
	}
	return x, nil
}

// Sub subtracts two Vector and returns a Vector
// It subtract's the corresponding element
func Sub(v, w []Vector) ([]Vector, error) {
	var x []Vector
	if len(v) != len(w) {
		return nil, fmt.Errorf("error: Vectors mus be of same length")
	} else {
		for i, _ := range v {
			for j, _ := range w {
				if i == j {
					x = append(x, v[i]-w[j])
				}
			}
		}
	}
	return x, nil
}

// VectorSum adds all the factors component wise
// i.e all the first element will be add to give the resultant
// vectors first element
// it checks if the given vectors is nil or not
// it also checks if the length of all the given vectors are same or not
func VectorSum(vectors ...[]Vector) ([]Vector, error) {
	var x []Vector
	if vectors == nil {
		return nil, fmt.Errorf("error: no vector is given")
	}

	// for checking the length of elements
	num_elements := len(vectors[0])
	for _, j := range vectors {
		if len(j) != num_elements {
			return nil, fmt.Errorf("all vectors should be of same size")
		}
	}

	// for addtion of vector elements
	for i := 0; i < num_elements; i++ {
		// sum variable should be of vector type
		// go addition of vector componentwise
		var sum Vector
		for _, j := range vectors {
			sum += j[i]
		}
		x = append(x, sum)
	}
	return x, nil
}

// ScalarMultiply multiplies scalar value to a vector
// and return the resultant vector
func ScalarMultiply(f float64, v []Vector) []Vector {
	var x []Vector
	for _, i := range v {
		x = append(x, Vector(f)*i)
	}
	return x
}

// VectorMean gives the Mean Vector
func VectorMean(vectors ...[]Vector) ([]Vector, error) {
	var x []Vector
	n := float64(len(vectors))

	s, err := VectorSum(vectors...)
	if err != nil {
		return nil, err
	}
	x = append(x, ScalarMultiply(1/n, s)...)
	return x, nil
}

// Dot computes v_1 * w_1 + ... + v_n * w_n
func Dot(v, w []Vector) float64 {
	var x float64
	for i := 0; i < len(v); i++ {
		x += float64(v[i] * w[i])
	}
	return x
}

// SumOfSquares compute v_1 * v_1 + ..... + v_n * v_n
func SumOfSquares(v []Vector) float64 {
	return Dot(v, v)
}

// Magnitude compute magnitude (or length) of Vector
func Magnitude(v []Vector) float64 {
	return math.Sqrt(SumOfSquares(v))
}

// SquaredDistance compute square distace between two vectors
func SquaredDistance(v, w []Vector) float64 {
	x, _ := Sub(v, w)
	return SumOfSquares(x)
}

func Distance(v, w []Vector) float64 {
	return math.Sqrt(SquaredDistance(v, w))
}

//// Matrix

type Matrix [][]Vector

// Shape returns number of rows and number of columns of m
func Shape(m Matrix) []int {
	num_rows, num_cols := len(m), len(m[0])
	si := []int{num_rows, num_cols}
	return si
}

// GetRow returns ith row of Matrix
func GetRow(m Matrix, i int) []Vector {
	return m[i]
}

// GetCow returns ith row of Matrix
func GetCol(m Matrix, j int) []Vector {
	var x []Vector
	for _, i := range m {
		x = append(x, i[j])
	}
	return x
}

// MakeMatrix will create a empty Matrix of given shape
func MakeMatrix(nr, nc int) (Matrix, error) {
	if nc != nr {
		return nil, fmt.Errorf("error: number of rows and columns should be equal")
	}
	a := make([][]Vector, nr)
	for i := range a {
		a[i] = make([]Vector, nc)
	}
	return a, nil
}

// IdentityMatrix will create a identity matrix of given shape
func IdentityMatrix(n int) Matrix {
	x, _ := MakeMatrix(n, n)
	for i, v := range x {
		for j, _ := range v {
			if i == j {
				x[i][j] = 1
			}
		}
	}
	return x
}

func main() {
	// var sum Vector
	// fmt.Println(sum+)
	// a := []Vector{1, 2}
	// b := []Vector{3, 4}
	//
	// ma := Matrix{a, b}
	// fmt.Println(ma)
	// c := []Vector{5, 6, 7}
	// d := []Vector{7, 8}

	// Test of Add
	// c, err := Add(a, b)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(c)
	// }

	// Test of Sub
	// d, err := Sub(a, b)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(d)
	// }

	// Test of VectorSum
	// e, err := VectorSum(a, b, c, d)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(e)
	// }

	// Test of ScalarMultiply
	// n := ScalarMultiply(3, a)
	// m := ScalarMultiply(3, b)
	// o := ScalarMultiply(3, c)
	// p := ScalarMultiply(3, d)
	// fmt.Println(n)
	// fmt.Println(m)
	// fmt.Println(o)
	// fmt.Println(p)

	// Test of func VectorMean
	// v, err := VectorMean(a, b, c)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 		fmt.Println(v)
	// }

	// Test of func Dot
	// v := Dot(a, b)
	// fmt.Println(v)

	// Test of func SumOfSquares
	// v := SumOfSquares(a)
	// fmt.Println(v)

	// Test of func Magnitude
	// v := Magnitude(a)
	// fmt.Println(v)

	// Test of func SquaredDistance
	// v := SquaredDistance(a, b)
	// fmt.Println(v)

	// Test of func Distance
	// v := Distance(a, b)
	// fmt.Println(v)

	// Test of Shape
	// m := Shape(ma)
	// fmt.Println(m)

	// Test of func GetRow
	// mr := GetRol(ma, 0)
	// fmt.Println(mr)

	// Test of func GetCol
	// mc := GetCol(ma, 0)
	// fmt.Println(mc)

	// Test of func MakeMatrix
	// m, err := MakeMatrix(3, 3)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(m)
	// }

	// Test of func IdentityMatrix
	im := IdentityMatrix(3)
	fmt.Println(im)
	// for _, i := range im {
	// 	fmt.Printf("%v\n", i)
	// }

}
