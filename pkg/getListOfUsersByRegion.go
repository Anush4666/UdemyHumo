package pkg

import (
	"fmt"
)

func GetListOfUsersByRegion() {

	var region string
	var has bool
	fmt.Println("Введите регион:")
	fmt.Scan(&region)
	for _, v := range Region {
		if v.Name == region {
			has = true
		}
	}
	if has {
		fmt.Println("Пользователи в данном регионе:")
		for _, v := range Users {
			if v.City.Region.Name == region {
				fmt.Println(v.Name)
			}
		}
	}
	if !has {
		fmt.Println("Данного региона нет в нашей БД!")

	}

}
