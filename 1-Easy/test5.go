/*
الجمع التسلسلي
5 نقاط لكل لغة
سهل
وصف التحدي
قم بكتابة دالة function تستقبل عدد صحيح integer. تقوم الـدالة بجمع جميع الأعداد من 1 إلى الرقم الذي تم تمريره وإرجاع النتيجة.

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
num = 3
المخرجات (Outputs)
6
الاختبار 2

المدخلات (Inputs)
num = 13
المخرجات (Outputs)
91
الاختبار 3

المدخلات (Inputs)
num = 15
المخرجات (Outputs)
120
الاختبار 4

المدخلات (Inputs)
num = 9
المخرجات (Outputs)
45
*/

package main

import "fmt"

func main() {
	fmt.Println(easyLoop(3))
	fmt.Println(easyLoop(13))
	fmt.Println(easyLoop(15))
	fmt.Println(easyLoop(9))

}

func easyLoop(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
