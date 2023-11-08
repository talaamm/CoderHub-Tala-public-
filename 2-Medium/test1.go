/*
أطول سلسلة أرقام متناوبة
10 نقاط لكل لغة
متوسط
وصف التحدي
قم بكتابة دالة function تستقبل قيمة نصية من نوع string. تقوم هذه الدالة بإيجاد أطول سلسلة فرعية بالتناوب بين أعداد فردية وأعداد زوجية. والمعنى من ذلك أن يجب ان تكون السلسلة مكونة من عدد زوجي ويتليه عدد فردي.. ولا يشترط في ذلك ان تبدأ السلسلة بعدد زوجي. فمن الممكن أن تبدأ بعدد زوجي أو فردي.

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
digits = '2105787220351146'
المخرجات (Outputs)
'2105'
الاختبار 2

المدخلات (Inputs)
digits = '1263654081858902'
المخرجات (Outputs)
'8185890'
الاختبار 3

المدخلات (Inputs)
digits = '334090830025543'
المخرجات (Outputs)
'090'
الاختبار 4

المدخلات (Inputs)
digits = '6769423178839463'
المخرجات (Outputs)
'67694'
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(longestNumOddEven("2105787220351146"))     //2105
	fmt.Println(longestNumOddEven("1263654081858902"))     //8185890
	fmt.Println(longestNumOddEven("3340901830025543"))     //0901830
	fmt.Println(longestNumOddEven("6769423178839463"))     //67694
	fmt.Println(longestNumOddEven("12345678912345688595")) //12345689

}

// fmt.Println(longestNumOddEven("334090830025543"))   the same but with len instead of indx
func longestNumOddEven(s string) string {
	numbers := []string{}

	theLong := 0

	var even1num bool

	Fnum, _ := strconv.Atoi(string(s[0]))
	if Fnum%2 == 0 {
		even1num = true
	} else {
		even1num = false
	}

	for indx, harf := range s {
		CurrNum, _ := strconv.Atoi(string(harf))
		even := false
		if CurrNum%2 == 0 {
			even = true
		} else {
			even = false
		}
		if even1num {
			if (even && indx%2 == 0) || !even && indx%2 == 1 {
				theLong = theLong*10 + CurrNum
			} else {
				str := strconv.Itoa(theLong)
				numbers = append(numbers, str)
				even1num, even, theLong = false, false, 0

				Fnum, _ = strconv.Atoi(string(harf))
				if Fnum%2 == 0 {
					even1num = true
				} else {
					even1num = false
				}
				theLong = theLong*10 + CurrNum
			}
		} else {
			if (!even && indx%2 == 0) || (even && indx%2 != 0) {
				theLong = theLong*10 + CurrNum
			} else {
				str := strconv.Itoa(theLong)
				numbers = append(numbers, str)
				even1num, even, theLong = false, false, 0

				Fnum, _ = strconv.Atoi(string(harf))
				if Fnum%2 == 0 {
					even1num = true
				} else {
					even1num = false
				}
				theLong = theLong*10 + CurrNum
			}
		}

	}
	return maxLenStrings(numbers)
}

func maxLenStrings(arr []string) string {
	lenSs := make([]int, len(arr))
	for i, str := range arr {
		lenSs[i] = len(str)
	}

	max := float64(lenSs[0])
	for _, num := range lenSs {
		max = math.Max(max, float64(num))
	}
	for i := range lenSs {
		if max == float64(lenSs[i]) {
			return arr[i]
		}
	}
	return "error"
}
