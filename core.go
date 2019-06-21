package mnv

import (
	"fmt"
	"regexp"
)

var emum = map[string]string{
	"cn":  "^(\\+?0?86\\-?)?1[345789]\\d{9}$",                                            // 中国大陆
	"tw":  "^(\\+?886\\-?|0)?9\\d{8}$",                                                   // 中国香港
	"hk":  "^(\\+?852\\-?)?[569]\\d{3}\\-?\\d{4}$",                                       // 中国台湾
	"mys": "^(\\+?6?01){1}(([145]{1}(\\-|\\s)?\\d{7,8})|([236789]{1}(\\s|\\-)?\\d{7}))$", // 马来西亚
	// "us":  "^(\\+?1)?[2-9]\\d{2}[2-9](?!11)\\d{6}$",                                      // 美国
	// "vn": "^(\\+?84|0)?((1(2([0-9])|6([2-9])|88|99))|(9((?!5)[0-9])))([0-9]{7})$", // 越南
}

func Veri(country string, tel string) bool {

	if _, exists := emum[country]; !exists {
		fmt.Println("Sorry,we didn't have this country mobile number")
		return false
	}

	return veri(tel, emum[country])
}

func Country(tel string) []string {

	r := []string{}

	for k, v := range emum {
		fmt.Println(k)
		if veri(tel, v) {
			r = append(r, k)
		}
	}

	return r
}

func veri(tel string, p string) bool {
	reg := regexp.MustCompile(p)
	return reg.MatchString(tel)
}
