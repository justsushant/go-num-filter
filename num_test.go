package num

import (
	"testing"
	"slices"
)

func TestEvenNumbers(t *testing.T) {
	got := CheckEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{2, 4, 6, 8, 10}

	assertSliceEquality(t, got, want)
}

func TestOddNumbers(t *testing.T) {
	got := CheckOddNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{1, 3, 5, 7, 9}

	assertSliceEquality(t, got, want)
}

func TestPrimeNumbers(t *testing.T) {
	got := CheckPrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{2, 3, 5, 7}

	assertSliceEquality(t, got, want)
}

func TestOddPrimeNumbers(t *testing.T) {
	got := CheckOddPrimeNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	want := []int{3, 5, 7}

	assertSliceEquality(t, got, want)
}

func TestEvenMultiplesOf5(t *testing.T) {
	got := CheckEvenMultiplesOf5([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	want := []int{10, 20}

	assertSliceEquality(t, got, want)
}

func TestOddMultipleOf3GreaterThan10(t *testing.T) {
	got := CheckOddMultipleOf3GreaterThan10([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	want := []int{15}

	assertSliceEquality(t, got, want)
}

func TestAllConditions(t *testing.T) {
	testCases := []struct{
		name string
		nums []int
		conditions []Condition
		expOutput []int
	}{
		{
			"check for odd, greater than 5, and multiple of 3",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			[]Condition{isOdd, isGreaterThan5, isMultipleOf3},
			[]int{9, 15},
		},
		{
			"check for even, less than 15, and multiple of 3", 
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			[]Condition{isEven, isLessThan15, isMultipleOf3},
			[]int{6, 12},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckAllConditions(tc.nums, tc.conditions...)
			want := tc.expOutput

			assertSliceEquality(t, got, want)
		})
	}
}

func TestAnyConditions(t *testing.T) {
	testCases := []struct{
		name string
		nums []int
		conditions []Condition
		expOutput []int
	}{
		{
			"check for prime, greater than 15, and multiple of 5",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			[]Condition{isPrime, isGreaterThan15, isMultipleOf5},
			[]int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20},
		},
		{
			"check for less than 6, and multiple of 3",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			[]Condition{isLessThan6, isMultipleOf3},
			[]int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckAnyConditions(tc.nums, tc.conditions...)
			want := tc.expOutput

			assertSliceEquality(t, got, want)
		})
	}
}

func assertSliceEquality(t *testing.T, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}