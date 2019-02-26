package main

import (
	"fmt"
	"regexp"
)
type Level uint32


const (
	DebugLevel Level = iota + 2
	InfoLevel
	WarnLevel  
	ErrorLevel
)

func Regexp_Func() {
	var str = "financeMsg: hp12345mqwrw"
	// var valid = regexp.MustCompile("[0-9]")
	// fmt.Println(valid.FindAllStringSubmatch(str, -1))
	re := regexp.MustCompile("finance")
	fmt.Println(re.FindString(str))
}

func main() {
	// Regexp_Func()
	Regexp()	

}