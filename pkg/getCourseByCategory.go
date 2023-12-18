package pkg

import (
	"fmt"
)

func GetCourseByCategory() {

	var course string
	var has bool
	fmt.Println("Введите категорию:")
	fmt.Scan(&course)
	for _, v := range Course {
		if v.Category.Name == course {
			has = true
		}
	}
	if has {
		fmt.Println("Курсы в данной категории:")
		for _, v := range Course {
			if v.Category.Name == course {
				fmt.Println(v.Name)
			}
		}
	}
	if !has {
		fmt.Println("Категория не найдена в нашем БД!")
	}
}
