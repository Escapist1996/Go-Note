package main

import "fmt"

type name string



func main() {
	var name string = "bl"
	var i,j,k int = 1,2,3
	var(
		name string
		age int
	   )
	str := "string"
	_,ret := 2,3

	const (
		ConfigKey_finance_db ConfigKey = 0
		ConfigKey_afterloan_db ConfigKey = 1
		ConfigKey_dayhole_db ConfigKey = 2
		ConfigKey_loanaudit_db ConfigKey = 3
		  )
	var config_key map[int64] {
		"finance_db",
		"afterloan_db",
		"dayhole_db",
		)

	mympa := make(map[string]int)

	for name, job := range map2 {
		fmt.Println(name, job)
	}
	/*
	var myname name = "zs"
	bt := []byte(myname)
	fmt.Println(len(bt))
	fmt.Println("%d,%d", bt[0], bt[1])
	*/
}
