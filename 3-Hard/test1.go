/*
التحويل من 12 ساعة إلى 24 ساعة والعكس
20 نقطة لكل لغة
صعب
وصف التحدي
قم بكتابة دالة function تستقبل قيمة نصية string تعبر عن الساعة كنظام 12 أو 24 ساعة و تقوم بالتحويل بينها ، و قم بإلحاق نظام 12 ساعة بـ am أو pm ثم قم بارجاع النتيجة من نوع string

المخرجات المتوقعة
الاختبار 1

المدخلات (Inputs)
time = '10:30 am'
المخرجات (Outputs)
'10:30'
الاختبار 2

المدخلات (Inputs)
time = '7:13 pm'
المخرجات (Outputs)
'19:13'
الاختبار 3

المدخلات (Inputs)
time = '19:00'
المخرجات (Outputs)
'7:00 pm'
الاختبار 4

المدخلات (Inputs)
time = '12:00 am'
المخرجات (Outputs)
'0:00'
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 12; i++ {
		fmt.Print(strconv.Itoa(i) + ":45 am-->")
		fmt.Println(T12to24_ViceVersa(strconv.Itoa(i) + ":45 am"))
	}
	for i := 1; i <= 12; i++ {
		fmt.Print(strconv.Itoa(i) + ":26 pm-->")
		fmt.Println(T12to24_ViceVersa(strconv.Itoa(i) + ":26 pm"))
	}
	for i := 0; i <= 23; i++ {
		fmt.Print(strconv.Itoa(i) + ":15-->")
		fmt.Println(T12to24_ViceVersa(strconv.Itoa(i) + ":15"))
	}
	fmt.Println(T12to24_ViceVersa("45:15"))
	fmt.Println(T12to24_ViceVersa("-1:15"))
	fmt.Println(T12to24_ViceVersa("15:15 pm"))

}

func T12to24_ViceVersa(s string) string {
	var answer string
	if len(s) == 5 {
		// in 24 format 2 digits   -->10:00-->23:00
		h := s[:2]
		hour, _ := strconv.Atoi(h)
		if hour > 12 && hour <= 23 {
			hour = hour - 12
			answer = strconv.Itoa(hour) + s[2:] + " pm"
		} else if hour == 12 {
			answer = s + " pm"
		} else if hour == 10 || hour == 11 {
			answer = s + " am"
		}

	} else if len(s) == 4 {
		// in 24 format 1 digits --> 0:00 --> 9:00
		h := s[:1]
		hour, _ := strconv.Atoi(h)
		if hour > 0 && hour < 12 {
			answer = s + " am"
		} else if hour == 0 {
			answer = "12" + s[1:] + " am"
		}

	} else if len(s) == 8 {
		//in 12 format 2 digits
		pam := s[6:]
		h := s[:2]
		hour, _ := strconv.Atoi(h)
		if pam == "am" && hour == 12 {
			hour = 0
			answer = strconv.Itoa(hour) + s[2:5]

		} else if pam == "am" && hour < 12 && hour >= 1 {
			answer = s[:5]
		} else if pam == "pm" && hour < 12 && hour >= 1 {
			hour += 12
			answer = strconv.Itoa(hour) + s[2:5]
		} else if pam == "pm" && hour == 12 {
			answer = s[:5]
		}

	} else if len(s) == 7 {
		//in 12 format 1 digits
		pam := s[5:]
		if pam == "am" {
			answer = s[:4]
		} else if pam == "pm" {
			h := s[:1]
			hour, _ := strconv.Atoi(h)
			hour += 12
			answer = strconv.Itoa(hour) + s[1:4]
		}

	}
	if answer == "" {
		answer = "an error ocuured try again!"
	}
	return answer
}
