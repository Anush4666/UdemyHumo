package pkg

import (
	"fmt"
)

func GetUsersCourse() {
	var user string
	var has bool
	fmt.Println("Введите имя пользователя:")
	fmt.Scan(&user)
	for _, v := range Users {
		if user == v.Name {
			has = true
		}
	}
	if has {
		fmt.Println("Курсы на которые подписаны пользователи:")
		for _, v := range Subscription {
			if v.User.Name == user {
				fmt.Println(v.Course.Name)
			}
		}
	}
	if !has {
		fmt.Println("Пользователь не найден на нашем БД!")
	}
}
