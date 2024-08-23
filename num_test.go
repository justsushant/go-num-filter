package num

import (
	"testing"
	"slices"
)

var (
	isGreaterThan15 = isGreaterThanX[int](15)
	isLessThan6 = isLesserThanX[int](6)
	isLessThan15 = isLesserThanX[int](15)
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
			[]Condition{isOdd, Condition(isGreaterThan5), Condition(isMultipleOf3)},
			[]int{9, 15},
		},
		{
			"check for even, less than 15, and multiple of 3", 
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			[]Condition{isEven, Condition(isLessThan15), Condition(isMultipleOf3)},
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
			[]Condition{isPrime, Condition(isGreaterThan15), Condition(isMultipleOf5)},
			[]int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20},
		},
		{
			"check for less than 6, and multiple of 3",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			[]Condition{Condition(isLessThan6), Condition(isMultipleOf3)},
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

func TestAllConditionsGeneric(t *testing.T) {
	t.Run("testing on integer numbers", func(t *testing.T) {
		testCases := []struct {
			name       string
			nums       []int
			conditions []ConditionGeneric[int]
			expOutput  []int
		}{
			{
				"check for odd, greater than 5, and multiple of 3",
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				[]ConditionGeneric[int]{isOdd, isGreaterThanX(5), isMultipleOfX(3)},
				[]int{9, 15},
			},
			{
				"check for even, less than 15, and multiple of 3",
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				[]ConditionGeneric[int]{isEven, isLesserThanX(15), isMultipleOfX(3)},
				[]int{6, 12},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := CheckAllConditionsGeneric(tc.nums, tc.conditions...)
				want := tc.expOutput

				assertSliceEquality(t, got, want)
			})
		}
	})

	t.Run("testing on Float numbers", func(t *testing.T) {
		testCases := []struct {
			name       string
			nums       []float64
			conditions []ConditionGeneric[float64]
			expOutput  []float64
		}{
			{
				"check for greater than 5, less than 10, and divisible by 2",
				[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0},
				[]ConditionGeneric[float64]{isGreaterThanX(5.9), isLesserThanX(10.1)},
				[]float64{6.6, 7.7, 8.8, 9.9, 10.0},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := CheckAllConditionsGeneric(tc.nums, tc.conditions...)
				want := tc.expOutput

				assertSliceEquality(t, got, want)
			})
		}
	})
}

func TestAnyConditionsGeneric(t *testing.T) {
	t.Run("testing on integer numbers", func(t *testing.T) {
		testCases := []struct {
			name       string
			nums       []int
			conditions []ConditionGeneric[int]
			expOutput  []int
		}{
			{
				"check for prime, greater than 15, and multiple of 5",
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				[]ConditionGeneric[int]{isPrime, isGreaterThanX(15), isMultipleOfX(5)},
				[]int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20},
			},
			{
				"check for less than 6, and multiple of 3",
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				[]ConditionGeneric[int]{isLesserThanX(6), isMultipleOfX(3)},
				[]int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := CheckAnyConditionsGeneric(tc.nums, tc.conditions...)
				want := tc.expOutput

				assertSliceEquality(t, got, want)
			})
		}
	})

	t.Run("testing on float numbers", func(t *testing.T) {
		testCases := []struct {
			name       string
			nums       []float64
			conditions []ConditionGeneric[float64]
			expOutput  []float64
		}{
			{
				"check for greater than 6, or less than 3",
				[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0},
				[]ConditionGeneric[float64]{isGreaterThanX(6.3), isLesserThanX(3.3)},
				[]float64{1.1, 2.2, 6.6, 7.7, 8.8, 9.9, 10.0},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got := CheckAnyConditionsGeneric(tc.nums, tc.conditions...)
				want := tc.expOutput

				assertSliceEquality(t, got, want)
			})
		}
	})
}

func assertSliceEquality[T comparable](t *testing.T, got, want []T) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("Expected %v but got %v", want, got)
	}
}