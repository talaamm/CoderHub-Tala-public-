/*
حساب مابعد الفاصلة
5 نقاط لكل لغة
سهل
وصف التحدي
قم بكتابة دالة function تستقبل قيمة نصية من نوع string عبارة عن على عدد عشري أو صحيح. تقوم الدالة بحساب عدد الأرقام التي بعد الفاصلة وإرجاع النتيجة على شكل قيمة من نوع integer.

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
num = '3.967'
المخرجات (Outputs)
3
الاختبار 2

المدخلات (Inputs)
num = '2.9'
المخرجات (Outputs)
1
الاختبار 3

المدخلات (Inputs)
num = '200'
المخرجات (Outputs)
0
الاختبار 4

المدخلات (Inputs)
num = '5.19054'
المخرجات (Outputs)
5
*/

package main

import "fmt"

func main() {
	fmt.Println(afterComma("3.952"))
	fmt.Println(afterComma("2.9"))
	fmt.Println(afterComma("200"))
	fmt.Println(afterComma("561.8645846"))

}

func afterComma(s string) int {
	str := ""
	for i, r := range s {
		if r == '.' {
			str = s[i+1:]
		}
	}
	return len(str)
}
