package main

import (
	"fmt"
)

type user struct {
	name	string
	email	string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}
func main() {
	var bill user
	bill = user{"Bill", "bill@email.com"}
	bill.notify()
	
	sansa := &user{"Sansa", "sansa@email.com"}
	sansa.notify()
	
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	sansa.changeEmail("sansa@newdomain.com")
	sansa.notify()
}

