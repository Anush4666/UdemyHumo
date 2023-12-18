package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCourse() {
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "Golang",
		Price:    500,
		Category: pkg.Category[0],
		Duration: 3,
	})

	pkg.Course = append(pkg.Course, models.Course{
		Name:     "UX/UI",
		Price:    350,
		Category: pkg.Category[1],
		Duration: 5,
	})

	pkg.Course = append(pkg.Course, models.Course{
		Name:     "Python",
		Price:    450,
		Category: pkg.Category[2],
		Duration: 4,
	})

	pkg.Course = append(pkg.Course, models.Course{
		Name:     "C++",
		Price:    600,
		Category: pkg.Category[3],
		Duration: 10,
	})
}
