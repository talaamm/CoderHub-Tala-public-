/*
صيغة البريد الإلكتروني
10 نقاط لكل لغة
متوسط
وصف التحدي
قم بكتابة دالة function تستقبل قيمة نصية من نوع string. تقوم الدالة بالتحقق ما إذا كانت القيمة النصية تحتوي على بريد إلكتروني صحيح أم لا، عن طريق إرجاع قيمة من نوع boolean.

شروط البريد الإلكتروني الصحيح

أن لا يبدأ رمز

أن يحتوي الجزء الأول على خانة على الأقل (ما قبل الرمز @)

أن يحتوي على الرمز @

أن يتبع رمز @ خانة على الأقل

أن يتبعه بعد ذلك بالرمز .

أن يتبع الرمز . بخانتين على الأقل

لا مانع من أن يحتوي الجزء الذي يسبق @ على الرموز التالية _.-

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
email = 'example@example.com'
المخرجات (Outputs)
true
الاختبار 2

المدخلات (Inputs)
email = 'example@example.c'
المخرجات (Outputs)
false
الاختبار 3

المدخلات (Inputs)
email = 'example@com'
المخرجات (Outputs)
false
الاختبار 4

المدخلات (Inputs)
email = '@example.com'
المخرجات (Outputs)
false
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isEmail("example@example.com"))
	fmt.Println(isEmail("example@example.c"))
	fmt.Println(isEmail("examp.fkr-le@example.com"))
	fmt.Println(isEmail("exa_mple@example.com"))
	fmt.Println(isEmail("@example.com"))
	fmt.Println(isEmail("example@com"))
	fmt.Println(isEmail("example.com"))

}

func isEmail(s string) bool {
	//s:= tala@gmail.com
	pt := strings.Split(s, "@") //["tala" "gmail.com"]
	if len(pt) != 2 {
		return false
	}
	pt1 := pt[0]                             //tala
	ptt := strings.Split(pt[len(pt)-1], ".") //["gmail" "com"]
	if len(ptt) != 2 {
		return false
	}
	pt2 := ptt[0] //gmail
	pt3 := ptt[1] //com

	if len(pt1) == 0 {
		return false
	}
	if len(pt2) == 0 {
		return false
	}
	if len(pt3) < 2 {
		return false
	}
	return true
}
