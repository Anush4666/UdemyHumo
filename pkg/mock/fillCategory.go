package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCategory() {
	pkg.Category = append(pkg.Category, models.Category{
		Name: "Back-End",
	})

	pkg.Category = append(pkg.Category, models.Category{
		Name: "Design",
	})

	pkg.Category = append(pkg.Category, models.Category{
		Name: "Программирование",
	})

	pkg.Category = append(pkg.Category, models.Category{
		Name: "Программирование",
	})

}
