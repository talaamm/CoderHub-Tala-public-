/*
تحويل رقم ثنائي إلى رقم ثماني
20 نقطة لكل لغة
صعب
وصف التحدي
قم بكتابة دالة function تستقبل متغير من نوع string يعبر عن قيمة ثمانية binary number، ثم قم بإرجاع النتيجة بعد التحويل الى قيمة ست عشرية octal number بنوع int

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
b = '101010101010'
المخرجات (Outputs)
5252
الاختبار 2

المدخلات (Inputs)
b = '1000000000'
المخرجات (Outputs)
1000
الاختبار 3

المدخلات (Inputs)
b = '10101111000'
المخرجات (Outputs)
2570
الاختبار 4

المدخلات (Inputs)
b = '1'
المخرجات (Outputs)
1
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(binToOct("101010101010"))
	fmt.Println(binToOct("1000000000"))
	fmt.Println(binToOct("10101111000"))
	fmt.Println(binToOct("1"))

}

func binToOct(s string) int {
	num, _ := strconv.ParseInt(s, 2, 32)
	oct, _ := strconv.Atoi(fmt.Sprintf("%o", num))
	return (oct)
}
