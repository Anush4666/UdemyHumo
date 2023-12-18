package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCities() {
	pkg.City = append(pkg.City, models.City{
		Name:   "Душанбе",
		Region: pkg.Region[0],
	})

	pkg.City = append(pkg.City, models.City{
		Name:   "Худжанд",
		Region: pkg.Region[1],
	})

	pkg.City = append(pkg.City, models.City{
		Name:   "Истаравшан",
		Region: pkg.Region[2],
	})

	pkg.City = append(pkg.City, models.City{
		Name:   "Куляб",
		Region: pkg.Region[3],
	})
}
