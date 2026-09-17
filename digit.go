package chinese

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	numbers = map[rune]rune{
		48: '零',
		49: '一',
		50: '二',
		51: '三',
		52: '四',
		53: '五',
		54: '六',
		55: '七',
		56: '八',
		57: '九',
	}
	reDate     = regexp.MustCompile(`((\d{2,4}年)|(\d{1,2}月)|(\d{1,2}日))`)
	reFraction = regexp.MustCompile(`\d+/\d+`)
	reCelsius  = regexp.MustCompile(`-?\d+℃`)
	rePercent  = regexp.MustCompile(`-?(\d+\.)?\d+%`)
	reDecimal  = regexp.MustCompile(`-?(\d+\.)?\d+`)
	reDateYear = regexp.MustCompile(`\d+年`)
	reNumber   = regexp.MustCompile(`\d+`)
)

func Digit2Chinese(s string) string {
	s = reDate.ReplaceAllStringFunc(s, replaceDate)
	s = reFraction.ReplaceAllStringFunc(s, replaceFraction)
	s = reCelsius.ReplaceAllStringFunc(s, replaceCelsius)
	s = rePercent.ReplaceAllStringFunc(s, replacePercent)
	return reDecimal.ReplaceAllStringFunc(s, replaceDecimal)
}

func replaceDate(s string) string {
	s = reDateYear.ReplaceAllStringFunc(s, replaceDirect)
	return reNumber.ReplaceAllStringFunc(s, replaceNumber)
}

func replaceFraction(s string) string {
	a := strings.Split(s, "/")
	return replaceNumber(a[1]) + "分之" + replaceNumber(a[0])
}

func replaceCelsius(s string) string {
	runes := []rune(s)
	var r string
	if runes[0] == '-' {
		r += "零下"
		runes = runes[1 : len(runes)-1]
	} else {
		runes = runes[:len(runes)-1]
	}
	r += replaceNumber(string(runes))
	return r + "摄氏度"
}

func replacePercent(s string) string {
	return "百分之" + replaceDecimal(s[:len(s)-1])
}

func replaceDecimal(s string) string {
	var r string
	if s[0] == '-' {
		r += "负"
		s = s[1:]
	}
	a := strings.Split(s, ".")
	r += replaceNumber(a[0])
	if len(a) == 2 {
		r += "点" + replaceDirect(a[1])
	}
	return r
}

func replaceDirect(s string) string {
	runes := []rune(s)
	results := make([]rune, 0, len(runes))
	for _, v := range runes {
		if n, ok := numbers[v]; ok {
			results = append(results, n)
		} else {
			results = append(results, v)
		}
	}
	return string(results)
}

func replaceNumber(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return ""
	}
	return numberToChinese(n)
}

func numberToChinese(n int) string {
	var r string
	if yi := n / 100000000; yi > 0 {
		r += numberToChinese(yi) + "亿"
		n %= 100000000
	}
	if wan := n / 10000; wan > 0 {
		r += numberToChinese(wan) + "万"
		n %= 10000
	} else if r != "" {
		r += "零"
	}
	if qian := n / 1000; qian > 0 {
		r += getNumberString(qian) + "千"
		n %= 1000
	} else if r != "" {
		r += "零"
	}
	if bai := n / 100; bai > 0 {
		r += getNumberString(bai) + "百"
		n %= 100
	} else if r != "" {
		r += "零"
	}
	if shi := n / 10; shi > 0 {
		if r == "" && shi == 1 {
			r += "十"
		} else {
			r += getNumberString(shi) + "十"
		}
		n %= 10
	} else if r != "" {
		r += "零"
	}
	if n > 0 {
		r += getNumberString(n)
	} else if r == "" {
		r = "零"
	}
	return r
}

func getNumberString(n int) string {
	return string(numbers[rune(n)+48])
}
