package pkg

import (
	"fmt"
)

func GetUsersExpense() {
	var (
		name    string
		has     bool
		balance int
	)
	fmt.Println("Введите имя пользователя:")
	fmt.Scan(&name)

	for _, course := range Subscription {
		if course.User.Name == name {
			balance += course.Price
			has = true
		}
	}

	if has {
		fmt.Println("Сумма которую потратил студент", balance)
	} else {
		fmt.Println("Данного пользователя нет в нашей бд!")
	}
}
