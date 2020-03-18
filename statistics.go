package main

import (
	"fmt"
	"github.com/data_science_with_golang/linearAlgebra"
	"math"
	"sort"
)

type data float64

// Mean calculate the mean value of the given data
func Mean(xs []data) float64 {
	var sum float64
	for _, i := range xs {
		sum += float64(i)
	}
	return sum / float64(len(xs))
}

// Median calculate the median of the given data
func medianOdd(xs []data) float64 {
	// If len(xs) is odd, the median is the middle element
	var fs []float64
	for i := 0; i < len(xs); i++ {
		fs = append(fs, float64(xs[i]))
	}
	sort.Float64s(fs)
	return fs[len(fs)/2]
}

func medianEven(xs []data) float64 {
	// If len(xs) is even it's the average of the middle two elements
	var fs []float64
	for i := 0; i < len(xs); i++ {
		fs = append(fs, float64(xs[i]))
	}
	sort.Float64s(fs)
	high_midpoint := len(fs) / 2.0
	return (fs[high_midpoint-1] + fs[high_midpoint]) / 2.0
}

// Median find's the middle-most value of v
func Median(v []data) float64 {
	if len(v)%2 == 0 {
		return medianEven(v)
	} else {
		return medianOdd(v)
	}
}

// Quantile returns the p-th percentile value in xs
func Quantile(xs []data, p float64) float64 {
	var fs []float64
	for i := 0; i < len(xs); i++ {
		fs = append(fs, float64(xs[i]))
	}
	p_index := int(p * float64(len(fs)))
	sort.Float64s(fs)
	return fs[p_index]
}

// Mode returns the list since there is more than one mode
func Mode(xs []data) []data {
	var m []data
	var maxFreq int
	counts := make(map[data]int)
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
func DataRange(xs []data) float64 {
	var fs []float64
	for i := 0; i < len(xs); i++ {
		fs = append(fs, float64(xs[i]))
	}
	sort.Float64s(fs)
	return fs[len(fs)-1] - fs[0]
}

// DeMean will translate xs by subtracting its mean (so the result has mean 0)
func DeMean(xs []data) []data {
	var dm []data
	x_bar := Mean(xs)
	for _, v := range xs {
		dm = append(dm, v-data(x_bar))
	}
	return dm
}

// Variance gives the average squared deviation from the mean
func Variance(xs []data) (float64, error) {
	if len(xs) < 2 {
		return 0, fmt.Errorf("Variance requires at least two elements")
	}

	n := len(xs)
	deviations := DeMean(xs)
	var fs []linearAlgebra.Vector
	for i := 0; i < len(deviations); i++ {
		fs = append(fs, linearAlgebra.Vector(deviations[i]))
	}
	return linearAlgebra.SumOfSquares(fs) / float64(n-1), nil
}

// StdDev gives the squared root of variance
func StdDev(xs []data) (float64, error) {
	s, err := Variance(xs)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(s), nil
}

// InterQuartileRange returns the difference between the 75%-ile and the 25%-ile
func InterQuartileRange(xs []data) float64 {
	return Quantile(xs, 0.75) - Quantile(xs, 0.25)
}

func main() {
	numFriends := []data{100.0, 49, 41, 40, 25, 21, 21, 19, 19, 18, 18, 16, 15, 15, 15, 15, 14, 14, 13, 13, 13, 13, 12, 12, 11, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	// me := Mean(numFriends)
	// fmt.Println(me)

	// var sum float64
	// dme, err := Variance(numFriends)
	// if err != nil {
	// 	fmt.Println(err)
	// }else {
	// 	fmt.Println(dme)
	// }

	s, _ := StdDev(numFriends)
	fmt.Println(s)

	iqr := InterQuartileRange(numFriends)
	fmt.Println(iqr)
}
