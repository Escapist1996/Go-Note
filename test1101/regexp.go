package main

import (
	"regexp"
	"fmt"
)

func Regexp() {
	var str = "Msg: hp12345mqwrw"
	var valid = regexp.MustCompile("[0-9]")
	fmt.Println(valid.FindAllStringSubmatch(str, -1))

}