package main

 

import "fmt"

 

type myError struct {

  err string

}

 

func (e *myError) Error() string {

  return e.err

}

 

 // func do() error {

func do() interface{} {

  var me *myError

  fmt.Printf("me == nil ? %v\n", me == nil) //返回true

  return me

}

 

func main() {

  err := do()

  fmt.Println(err == nil) //返回false

}


