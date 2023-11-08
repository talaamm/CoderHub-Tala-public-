/*
أحرف العلة
20 نقطة لكل لغة
صعب
وصف التحدي
قم بكتابة دالة function تستقبل متغيرين، المتغير الأول عبارة عن نص من نوع string و المتغير الثاني عبارة عن عدد من نوع integer ، تقوم الدالة باستخراج أحرف العلة من النص المدخل إعتمادًا على العدد المدخل. في حال أن العدد المدخل أكبر من حروف العلة : تقوم الدالة بإظهار رسالة invalid ، و في حال أن العدد المدخل أصغر أو مساوي لحروف العلة : تقوم بإرجاع الحروف من نوع string بناء على العدد المدخل.

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
phrase = 'Queen'
n = 3
المخرجات (Outputs)
'uee'
الاختبار 2

المدخلات (Inputs)
phrase = 'First Time'
n = 3
المخرجات (Outputs)
'iie'
الاختبار 3

المدخلات (Inputs)
phrase = 'Apple'
n = 2
المخرجات (Outputs)
'Ae'
الاختبار 4

المدخلات (Inputs)
phrase = 'Riyadh'
n = 2
المخرجات (Outputs)
'ia'
*/

package main

import "fmt"

func main() {
	fmt.Println(weweLetters("Queen", 3))
	fmt.Println(weweLetters("First time", 3))
	fmt.Println(weweLetters("Apple", 2))
	fmt.Println(weweLetters("Riyadh", 2))
	fmt.Println(weweLetters("noor and tala", 4))
	fmt.Println(weweLetters(" wewe monami", 15))

}

func weweLetters(s string, n int) string {
	answer := []rune{}

	for _, i := range s {
		if i == 'A' || i == 'a' || i == 'o' || i == 'O' || i == 'i' || i == 'I' || i == 'u' || i == 'U' || i == 'e' || i == 'E' {
			answer = append(answer, i)
		}
	}
	if len(answer) == n {
		return string(answer)
	} else if n < len(answer) {
		str := ""
		for b := 0; b < n; b++ {
			str += string(answer[b])
		}
		return str
	}
	return "invalid"
}
