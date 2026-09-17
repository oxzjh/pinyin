package chinese

import (
	"regexp"
	"strings"
)

var (
	gInitials    = []string{"b", "p", "m", "f", "d", "t", "n", "l", "g", "k", "h", "j", "q", "x", "zh", "ch", "sh", "r", "z", "c", "s", "y", "w"}
	gPhoneticReg = regexp.MustCompile("[āáǎàēéěèōóǒòīíǐìūúǔùüǘǚǜńňǹḿ]")
	gToneReg     = regexp.MustCompile("^([a-z]+)([1-4])([a-z]*)$")
)

func Pinyin(r rune) []string {
	s, ok := pinyinDict[r]
	if ok {
		return strings.Split(s, ",")
	}
	return nil
}

func FirstPinyin(r rune) string {
	s, ok := pinyinDict[r]
	if ok {
		return strings.SplitN(s, ",", 2)[0]
	}
	return ""
}

// 声母
func ToInitial(py string) string {
	for _, v := range gInitials {
		if strings.HasPrefix(py, v) {
			return v
		}
	}
	return ""
}

// 声调(1,2,3,4)在韵母后
func ToToneRhyme(py string) string {
	return gPhoneticReg.ReplaceAllStringFunc(py, func(s string) string {
		return phoneticDict[s]
	})
}

// 声调(1,2,3,4)在最后
func ToToneTail(py string) string {
	return gToneReg.ReplaceAllString(ToToneRhyme(py), "$1$3$2")
}

// 没有声调
func ToNormal(py string) string {
	return gToneReg.ReplaceAllString(ToToneRhyme(py), "$1$3")
}
