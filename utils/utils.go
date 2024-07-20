package utils

import (
	"demo/database"
	"demo/entity"
	"github.com/go-faker/faker/v4"
)

type DataFaker interface {
	entity.User | entity.Product
}

func GenerateMockData[T DataFaker](n int) []T {
	var data []T
	for i := 0; i < n; i++ {
		var sample T
		err := faker.FakeData(&sample)
		if err != nil {
			return nil
		}
		data = append(data, sample)
	}
	return data
}

func SetUsersAvatarURL() {
	for i := range database.Users {
		database.Users[i].GenImageURL()
	}
}

func SetProductsImageURL() {
	for i := range database.Products {
		database.Products[i].GenImageURL()
	}
}

func Remove[T DataFaker](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
