package main

import "fmt"

type Human struct {
	Name        string
	LastName    string
	Age         int
	PhoneNumber string
}

func NewHuman(name string, lastName string, age int, phone string) *Human {
	return &Human{
		Name:        name,
		LastName:    lastName,
		Age:         age,
		PhoneNumber: phone,
	}
}

func (h *Human) Rename(name string, lastName string) {
	h.Name = name
	h.LastName = lastName
}

func (h *Human) ChangePhoneNumber(phone string) {
	h.PhoneNumber = phone
}

func (h *Human) BirthdayToday() {
	h.Age++
}

type Action struct {
	Human
	Role string
}

func main() {
	person := Action{
		Human: Human{Name: "Ilyas", LastName: "Yunusov", Age: 19, PhoneNumber: "89178439502"},
		Role:  "водитель",
	}

	person.Rename("Ilya", "Ivanov")
	person.ChangePhoneNumber("8921039419")
	person.BirthdayToday()

	fmt.Println("Данные о человеке:")
	fmt.Println("Имя:", person.Name, ", Фамилия:", person.LastName, ", Возраст:", person.Age, ", Номер телефона:", person.PhoneNumber, ", Деятельность:", person.Role)
}