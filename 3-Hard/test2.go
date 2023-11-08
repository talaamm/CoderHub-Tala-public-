/*

فصل الكلمات
20 نقطة لكل لغة
صعب
وصف التحدي
قم بكتابة دالة function تستقبل قيمة نصية من نوع string ، تقوم الدالة بفصل الكلمات عند ملاحظة حرف كبير بوضع مسافه وتحويله لحرف صغير، ثم قم بارجاع النتيجة من نوع string

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
txt = 'wePlayTennis'
المخرجات (Outputs)
'we play tennis'
الاختبار 2

المدخلات (Inputs)
txt = 'iLikeCats'
المخرجات (Outputs)
'i like cats'
الاختبار 3

المدخلات (Inputs)
txt = 'computerScience'
المخرجات (Outputs)
'computer science'
الاختبار 4

المدخلات (Inputs)
txt = 'thankYou'
المخرجات (Outputs)
'thank you'
*/

package main

import "fmt"

func main() {
	fmt.Println(splitByCap("wePlayCards"))
	fmt.Println(splitByCap("iLikeCats"))
	fmt.Println(splitByCap("computerScience"))
	fmt.Println(splitByCap("thankYou"))

}

func splitByCap(s string) string {
	//32 diff
	arr := []rune{}
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			arr = append(arr, ' ')
			arr = append(arr, r+32)
		} else {
			arr = append(arr, r)
		}
	}
	return string(arr)
}
