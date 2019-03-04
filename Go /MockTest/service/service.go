package service

type db struct{}

type DB interface {
	FetchMessage(lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type greeter struct {
	datebase DB
	lang     string
}

type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}

func (d *db) FetchMessage(lang string) (string, error) {
	if lang == "en" {
		return "hello", nil
	}
	if lang == "es" {
		return "holla", nil
	}
	return "bzzzzz", nil
}

func (d *db) FetchDefaultMessage() (string, error) {
	return "default message", nil
}

func (g greeter) Greet() string {
	msg, _ := g.datebase.FetchMessage(g.lang)
	return "Message is: " + msg
}

func (g greeter) GreetInDefaultMsg() string {
	msg, _ := g.datebase.FetchDefaultMessage()
	return "Message is: " + msg
}

func NewDB() DB {
	return new(db)
}

func NewGreeter(db DB, lang string) GreeterService {
	return greeter{db, lang}
}
