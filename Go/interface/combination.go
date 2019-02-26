package file

// 组合

type Cat interface {
	cat()
}

type Dog interface {
	dog()
}

type Animal interface {
	Cat
	Dog
}

func get(i Animal) {
	i.cat()
	i.dog()
}
