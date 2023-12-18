package pkg

import "fmt"

func GetMiddleAgeOfUsers() {

	var course string
	var average int
	var has bool
	fmt.Println("Введите название курса:")
	fmt.Scan(&course)
	for _, v := range Course {
		if v.Name == course {
			has = true
		}
	}
	if has {
		fmt.Println("Средний возраст пользователей в данном курсе:")
		for _, v := range Users {
			average += 2023 - v.BirthYear
		}
		fmt.Println(average/len(Users) - 1)
	}
	if !has {
		fmt.Println("Курс не найден в нашем БД!")
	}
}
