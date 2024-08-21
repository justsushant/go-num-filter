package num

var (
	isMultipleOf3 = isMultipleOfX(3)
	isMultipleOf5 = isMultipleOfX(5)
	isGreaterThan5 = isGreaterThanX(5)
	isGreaterThan10 = isGreaterThanX(10)
	isGreaterThan15 = isGreaterThanX(15)
	isLessThan6 = isLessThanX(6)
	isLessThan15 = isLessThanX(15)
)

type Condition func(int) bool

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

func isMultipleOfX(x int) Condition {
	return func(num int) bool {
		return num % x == 0
	}
}

func isGreaterThanX(x int) Condition {
	return func(num int) bool {
		return num > x
	}
}

func isLessThanX(x int) Condition {
	return func(num int) bool {
		return x > num
	}
}
