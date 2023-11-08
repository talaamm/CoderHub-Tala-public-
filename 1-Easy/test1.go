/*
مساحة المستطيل
5 نقاط لكل لغة
سهل
وصف التحدي
قم بكتابة دالة function تستقبل عددين صحيحين موجبين من نوع integer يمثلان قيم الطول والعرض للمستطيل. تقوم الدالة بحساب مساحة المستطيل وإرجاع نتيجته.

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
width = 9
height = 8
المخرجات (Outputs)
72
الاختبار 2

المدخلات (Inputs)
width = 5
height = 5
المخرجات (Outputs)
25
الاختبار 3

المدخلات (Inputs)
width = 4
height = 3
المخرجات (Outputs)
12
الاختبار 4

المدخلات (Inputs)
width = 7
height = 4
المخرجات (Outputs)
28
*/

package main

import "fmt"

func mostateel_Area(width, height int) int {
	if width > 0 && height > 0 {
		return width * height
	}
	return -1
}

func main() {
	fmt.Println(mostateel_Area(9, 8))
	fmt.Println(mostateel_Area(5, 5))
	fmt.Println(mostateel_Area(4, 3))
	fmt.Println(mostateel_Area(7, 4))

}
