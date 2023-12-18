package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillUsers() {
	pkg.Users = append(pkg.Users, models.Users{
		Name:      "Anushervon",
		BirthYear: 2005,
		Balance:   20000,
		City:      pkg.City[0],
		Address:   "someAddress",
		Mail:      "someMail",
		Phone:     "+992988888888",
	})

	pkg.Users = append(pkg.Users, models.Users{
		Name:      "Umed",
		BirthYear: 2004,
		Balance:   25000,
		City:      pkg.City[1],
		Address:   "someAddress",
		Mail:      "someMail",
		Phone:     "+992977777777",
	})

	pkg.Users = append(pkg.Users, models.Users{
		Name:      "Ehson",
		BirthYear: 2005,
		Balance:   13000,
		City:      pkg.City[2],
		Address:   "someAddress",
		Mail:      "someMail",
		Phone:     "+992966666666",
	})

	pkg.Users = append(pkg.Users, models.Users{
		Name:      "Iskandar",
		BirthYear: 2002,
		Balance:   30000,
		City:      pkg.City[3],
		Address:   "someAddress",
		Mail:      "someMail",
		Phone:     "+992955555555",
	})
}
