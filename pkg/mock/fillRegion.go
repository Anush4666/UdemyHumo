package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillRegion() {
	pkg.Region = append(pkg.Region, models.Region{
		Name: "РРП",
	})

	pkg.Region = append(pkg.Region, models.Region{
		Name: "Согдийская_область",
	})

	pkg.Region = append(pkg.Region, models.Region{
		Name: "Согдийская_область",
	})

	pkg.Region = append(pkg.Region, models.Region{
		Name: "Хатлонская_область",
	})
}
