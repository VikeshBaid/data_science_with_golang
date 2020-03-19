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

// var dailyHr []data
// for

// Covariance measures how to variable vary in tandem from their Mean
func Covariance(xs, ys []data) (float64, error) {
	if len(xs) != len(ys) {
		return 0, fmt.Errorf("Both inputs should have same number of elements")
	}

	var vx, wy []linearAlgebra.Vector
	x := DeMean(xs)
	y := DeMean(ys)
	for i := 0; i < len(x); i++ {
		vx = append(vx, linearAlgebra.Vector(x[i]))
	}
	for i := 0; i < len(y); i++ {
		wy = append(wy, linearAlgebra.Vector(y[i]))
	}
	return linearAlgebra.Dot(vx, wy) / float64(len(xs)-1), nil
}

// Correlation measures thr ewlation b/w two or more variable. Correlation lies b/w -1(negative corr) and +1(positive corr)
func Correlation(xs, ys []data) float64 {
	stdX, _ := StdDev(xs)
	stdY, _ := StdDev(ys)
	cov, _ := Covariance(xs, ys)
	if stdX > 0 && stdY > 0 {
		return cov / stdX / stdY
	} else {
		return 0
	}
}

func main() {
	numFriends := []data{100.0, 49, 41, 40, 25, 21, 21, 19, 19, 18, 18, 16, 15, 15, 15, 15, 14, 14, 13, 13, 13, 13, 12, 12, 11, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	dailyMin := []data{1, 68.77, 51.25, 52.08, 38.36, 44.54, 57.13, 51.4, 41.42, 31.22, 34.76, 54.01, 38.79, 47.59, 49.1, 27.66, 41.03, 36.73, 48.65, 28.12, 46.62, 35.57, 32.98, 35, 26.07, 23.77, 39.73, 40.57, 31.65, 31.21, 36.32, 20.45, 21.93, 26.02, 27.34, 23.49, 46.94, 30.5, 33.8, 24.23, 21.4, 27.94, 32.24, 40.57, 25.07, 19.42, 22.39, 18.42, 46.96, 23.72, 26.41, 26.97, 36.76, 40.32, 35.02, 29.47, 30.2, 31, 38.11, 38.18, 36.31, 21.03, 30.86, 36.07, 28.66, 29.08, 37.28, 15.28, 24.17, 22.31, 30.17, 25.53, 19.85, 35.37, 44.6, 17.23, 13.47, 26.33, 35.02, 32.09, 24.81, 19.33, 28.77, 24.26, 31.98, 25.73, 24.86, 16.28, 34.51, 15.23, 39.72, 40.8, 26.06, 35.76, 34.76, 16.13, 44.04, 18.03, 19.65, 32.62, 35.59, 39.43, 14.18, 35.24, 40.13, 41.82, 35.45, 36.07, 43.67, 24.61, 20.9, 21.9, 18.79, 27.61, 27.21, 26.61, 29.77, 20.59, 27.53, 13.82, 33.2, 25, 33.1, 36.65, 18.63, 14.87, 22.2, 36.81, 25.53, 24.62, 26.25, 18.21, 28.08, 19.42, 29.79, 32.8, 35.99, 28.32, 27.79, 35.88, 29.06, 36.28, 14.1, 36.63, 37.49, 26.9, 18.58, 38.48, 24.48, 18.95, 33.55, 14.24, 29.04, 32.51, 25.63, 22.22, 19, 32.73, 15.16, 13.9, 27.2, 32.01, 29.27, 33, 13.74, 20.42, 27.32, 18.23, 35.35, 28.48, 9.08, 24.62, 20.12, 35.26, 19.92, 31.02, 16.49, 12.16, 30.7, 31.22, 34.65, 13.13, 27.51, 33.2, 31.57, 14.1, 33.42, 17.44, 10.12, 24.42, 9.82, 23.39, 30.93, 15.03, 21.67, 31.09, 33.29, 22.61, 26.89, 23.48, 8.38, 27.81, 32.35, 23.84}

	// me := Mean(numFriends)
	// fmt.Println(me)

	// var sum float64
	// dme, err := Variance(numFriends)
	// if err != nil {
	// 	fmt.Println(err)
	// }else {
	// 	fmt.Println(dme)
	// }

	// s, _ := StdDev(numFriends)
	// fmt.Println(s)
	//
	// iqr := InterQuartileRange(numFriends)
	// fmt.Println(iqr)

	// cov, _ := Covariance(numFriends, dailyMin)
	// fmt.Println(cov)

	corr := Correlation(numFriends, dailyMin)
	fmt.Println(corr)
}
