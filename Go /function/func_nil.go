package main

type Boy struct {
	Name string
	Age  int
}

func (this *Boy) SetName(name string) {
	this.Name = name
}
func (this *Boy) GetName() string {
	return this.Name
}
func main() {
	// boy := &Boy{Name: "H", Age: 21}
	// boy.SetName("P")
	// // boy = nil // error
	// fmt.Println(boy.GetName())
	// fmt.Println(boy)

	// boy := &Boy{Name: "H", Age: 21}
	// boy = nil
	// fmt.Println(boy)

	// fmt.Println(Boy{"H", 21}.GetName())

	b := Boy{}
	b.SetName("H")
}
