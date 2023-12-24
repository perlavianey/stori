package database

type Transaction struct {
	Id        string
	Date      string
	Amount    float64
	IdAccount string
	Filename  string
	Timestamp string
}

type Account struct {
	Id    string
	Name  string
	Email string
}
