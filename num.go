package num

import (
	"golang.org/x/exp/constraints"
)


var (
	isMultipleOf3 = isMultipleOfX[int](3)
	isMultipleOf5 = isMultipleOfX[int](5)
	isGreaterThan5 = isGreaterThanX[int](5)
	isGreaterThan10 = isGreaterThanX[int](10)
)

type Condition func(int) bool
type ConditionGeneric[T any] func(T) bool

func CheckEvenNumbers(nums []int) []int {
	var result []int
	for _, num := range nums {
		if isEven(num) {
			result = append(result, num)
		}
	}

	return result
}

func CheckOddNumbers(nums []int) []int {
	var result []int
	for _, num := range nums {
		if isOdd(num) {
			result = append(result, num)
		}
	}

	return result
}

func CheckPrimeNumbers(nums []int) []int {
	var result []int

	for _, num := range nums {
		if isPrime(num) {
			result = append(result, num)
		}
	}

	return result
}

func CheckOddPrimeNumbers(nums []int) []int {
	var result []int

	for _, num := range nums {
		if isOdd(num) && isPrime(num) {
			result = append(result, num)
		}
	}

	return result
}

func CheckEvenMultiplesOf5(nums []int) []int {
	var result []int

	for _, num := range nums {
		if isEven(num) && isMultipleOf5(num) {
			result = append(result, num)
		}
	}

	return result
}

func CheckOddMultipleOf3GreaterThan10(nums []int) []int {
	var result []int

	for _, num := range nums {
		if isGreaterThan10(num) && isOdd(num) && isMultipleOf3(num) {
			result = append(result, num)
		}
	}

	return result
}

// checks if all conditions are true for each element
func CheckAllConditions(nums []int, conditions ...Condition) []int {
	var result []int
	var flag bool

	for _, num := range nums {
		flag = true
		for _, condition := range conditions {
			if !condition(num) {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, num)
		}
	}

	return result
}

// checks if any conditions are true for each element
func CheckAnyConditions(nums []int, conditions ...Condition) []int {
	var result []int

	for _, num := range nums {
		for _, condition := range conditions {
			if condition(num) {
				result = append(result, num)
				break
			}
		}
	}

	return result
}


// checks if all conditions are true for each element, generic version
func CheckAllConditionsGeneric[T constraints.Ordered](num []T, conditions ...ConditionGeneric[T]) []T {
    var result []T
	var flag bool

    for _, num := range num {		
		flag = true
		for _, condition := range conditions {
			if !condition(num) {
				flag = false
				break
			}
		}

		if flag {
			result = append(result, num)
		}
    }

    return result
}

// checks if any condition is true for each element, generic version
func CheckAnyConditionsGeneric[T constraints.Ordered](nums []T, conditions ...ConditionGeneric[T]) []T {
    var result []T

    for _, num := range nums {
		for _, condition := range conditions {
			if condition(num) {
				result = append(result, num)
				break
			}
		}
    }

    return result
}


func isOdd(num int) bool {
	return num % 2 != 0
}

func isEven(num int) bool {
	return num % 2 == 0
}

func isPrime(num int) bool {
	if num == 0 || num == 1 {
		return false
	}

	for i := 2; i < num-1; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}

func isMultipleOfX[T constraints.Integer](x T) ConditionGeneric[T] {
    return func(num T) bool {
        return num%x == 0
    }
}


func isGreaterThanX[T constraints.Ordered](x T) ConditionGeneric[T] {
	return func(num T) bool {
		return num > x
	}
}

func isLesserThanX[T constraints.Ordered](x T) ConditionGeneric[T] {
	return func(num T) bool {
		return num < x
	}
}
