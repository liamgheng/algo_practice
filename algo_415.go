package algo_practice

import "strconv"

// MARK:  这里涉及一些字符和字符串操作，需要熟悉下
func addStrings(num1 string, num2 string) string {
	num1Index := len(num1) - 1
	num2Index := len(num2) - 1
	prev := 0  // 进位
	result := ""

	for num1Index >= 0 || num2Index >= 0 || prev > 0 {
		num1Value := 0
		if num1Index >= 0 {
			// MARK 注意这里的处理
			num1Value = int(num1[num1Index] - '0')
			num1Index -= 1
		}

		num2Value := 0
		if num2Index >= 0 {
			num2Value = int(num2[num2Index] - '0')
			num2Index -= 1
		}

		summation := prev + num1Value + num2Value
		result = strconv.Itoa(summation % 10) + result
		prev = summation / 10
	}

	return result
}
