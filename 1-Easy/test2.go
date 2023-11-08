/*
النصوص المتشابهة
5 نقاط لكل لغة
سهل
وصف التحدي
قم بكتابة دالة
function
تستقبل قيمتين نصية من نوع
string
وتقوم بالتحقق فيما إذا كانتا متشابهتين أم غير متشابهتين.

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
name1 = 'Anas'
name2 = 'anas'
المخرجات (Outputs)
'غير متشابهتين'
الاختبار 2

المدخلات (Inputs)
name1 = 'Omer'
name2 = 'Omar'
المخرجات (Outputs)
'غير متشابهتين'
الاختبار 3

المدخلات (Inputs)
name1 = 'Abdullah'
name2 = 'Abdullah'
المخرجات (Outputs)
'متشابهتين'
الاختبار 4

المدخلات (Inputs)
name1 = 'Ayman'
name2 = 'Mishal'
المخرجات (Outputs)

*/

package main

import (
	"fmt"
)

func main() {
	var name1, name2 string
	//	name1 = "Anas"
	//	name2 = "anas"
	//
	//	name1 = "Omer"
	//	name2 = "Omar"
	//
	name1 = "Abdullah"
	name2 = "Abdullah"

	//	name1 = "Ayman"
	//	name2 = "Mishal"

	if areTheSame(name1, name2) {
		fmt.Println("strings are the same")
	} else {
		fmt.Println("strings are NOT the same")

	}
}
func areTheSame(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
