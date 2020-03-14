package main

import (
	"fmt"
	//"github.com/data_science_with_golang/linearAlgebra"
	"sort"
)

type data []float64

// Mean calculate the mean value of the given data
func Mean(xs data) float64 {
	var sum float64
	for _, i := range xs {
		sum += i
	}
	return sum / float64(len(xs))
}

// Median calculate the median of the given data
func medianOdd(xs data) float64 {
	// If len(xs) is odd, the median is the middle element
	sort.Float64s(xs)
	return xs[len(xs)/2.0]
}

func medianEven(xs data) float64 {
	// If len(xs) is even it's the average of the middle two elements
	sort.Float64s(xs)
	high_midpoint := len(xs) / 2.0
	return (xs[high_midpoint-1] + xs[high_midpoint]) / 2.0
}

// Median find's the middle-most value of v
func Median(v data) float64 {
	if len(v)%2 == 0 {
		return medianEven(v)
	} else {
		return medianOdd(v)
	}
}

// Quantile returns the p-th percentile value in xs
func Quantile(xs data, p float64) float64 {
	p_index := int(p * float64(len(xs)))
	sort.Float64s(xs)
	return xs[p_index]
}

// Mode returns the list since there is more than one mode
func Mode(xs data) data {
	var m data
	var maxFreq int
	counts := make(map[float64]int)
	// Below loop is to add data to map
	for _, i := range xs {
		counts[i]++
		if counts[i] > maxFreq {
			maxFreq = counts[i]
		}
	}

	// this loop gives the mode of the data
	for i, _ := range counts {
		if maxFreq == counts[i] {
			m = append(m, i)
		}
	}
	return m
}

// Dispersion / range / spreadness of data
func DataRange(xs data) float64 {
	sort.Float64s(xs)
	return xs[len(xs)-1] - xs[0]
}

func DeMean(xs data) data {
	var dm data
	x_bar := Mean(xs)
	for _, v := range xs {
		dm = append(dm, v-x_bar)
	}
	return dm
}

func Variance(xs data) (float64, error) {
	if len(xs) < 2 {
		return nil, fmt.Errorf("Variance requires at least two elements")
	}
	n := len(xs)
	deviations := DeMean(xs)
		s := linearAlgebra.SumOfSquares(deviations)
	return s / (n - 1), nil
}

func main() {
	numFriends := data{100.0, 49, 41, 40, 25, 21, 21, 19, 19, 18, 18, 16, 15, 15, 15, 15, 14, 14, 13, 13, 13, 13, 12, 12, 11, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	me := Mode(numFriends)
	fmt.Println(me)

}
