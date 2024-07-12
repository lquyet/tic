package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	ID               string `json:"id" faker:"uuid_digit"`
	Name             string `json:"name" faker:"name"`
	Email            string `json:"email" faker:"email"`
	Phone            string `json:"phone" faker:"phone_number"`
	CreditCardNumber string `json:"creditCardNumber" faker:"cc_number"`
	Avatar           string `json:"avatar"`
	JoinedDate       int64  `json:"joinedDate" faker:"unix_time"`
	Age              int    `json:"age" faker:"oneof:18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

// text, number, date-time (unix time in second), link url

type GetUsersResponse struct {
	Data       []User     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func generateMockUsers(n int) []User {
	var data []User
	for i := 0; i < n; i++ {
		var sample User
		err := faker.FakeData(&sample)
		if err != nil {
			return nil
		}

		sample.Avatar = fmt.Sprintf("https://picsum.photos/id/%d/300/200", rand.Intn(99)+1)
		data = append(data, sample)
	}
	return data
}

var users []User

func getUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page != 0 && pageSize != 0 {
		queriedUsers := users[(page-1)*pageSize : page*pageSize]
		res := GetUsersResponse{
			Data: queriedUsers,
			Pagination: Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    len(users),
			},
		}

		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, GetUsersResponse{
		Data: users,
		Pagination: Pagination{
			Page:     1,
			PageSize: len(users),
			Total:    len(users),
		},
	})
}

func upsertUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.ID != "" {
		for i, u := range users {
			if u.ID == user.ID {
				users[i] = user
				return c.JSON(http.StatusOK, user)
			}
		}
	} else {
		user.ID = faker.UUIDDigit()
	}

	users = append(users, user)
	return c.JSON(http.StatusCreated, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	for i, u := range users {
		if u.ID == id {
			users = remove(users, i)
			return c.JSON(http.StatusOK, "User deleted")
		}
	}

	return c.JSON(http.StatusNotFound, "User not found")
}

func remove(s []User, i int) []User {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func resetUserData(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	if num == 0 {
		num = 10
	}
	users = generateMockUsers(num)
	return c.JSON(http.StatusOK, "Data reset")
}

func main() {
	e := echo.New()

	// First init users data
	users = generateMockUsers(10)

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Routes
	e.GET("/users", getUsers)
	e.POST("/users", upsertUser)
	e.DELETE("/users/:id", deleteUser)

	e.POST("/users/reset", resetUserData)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
