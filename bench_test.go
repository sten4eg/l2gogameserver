package main

import (
	"errors"
	"math/rand"
	"testing"
)

var notPositive = errors.New("not Positive!")

const size = 10000

var myDataForTest []int

func init() {
	myDataForTest = make([]int, 0, size)
	for i := 0; i < size; i++ {
		myDataForTest = append(myDataForTest, (rand.Int()%1000)-500)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum_abs_value_exception(myDataForTest)
	}
}
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum_abs_value(myDataForTest)
	}
}

func sum_abs_value_exception(a []int) int {
	var sum int
	for _, x := range a {
		err := get_positive_value(x)
		if err != nil {
			sum += -x
		} else {
			sum += x
		}

	}

	return sum
}

func get_positive_value(a int) error {
	if a < 0 {
		return notPositive
	}
	return nil
}
func sum_abs_value(a []int) int {
	var sum int
	for _, x := range a {
		if x < 0 {
			sum += -x
		} else {
			sum += x
		}
	}

	return sum
}
