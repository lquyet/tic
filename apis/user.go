package apis

import (
	"demo/database"
	"demo/entity"
	"demo/utils"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page != 0 && pageSize != 0 {
		queriedUsers := database.Users[(page-1)*pageSize : page*pageSize]
		res := entity.GetUsersResponse{
			Data: queriedUsers,
			Pagination: entity.Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    len(database.Users),
			},
		}

		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, entity.GetUsersResponse{
		Data: database.Users,
		Pagination: entity.Pagination{
			Page:     1,
			PageSize: len(database.Users),
			Total:    len(database.Users),
		},
	})
}

func UpsertUser(c echo.Context) error {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.ID != "" {
		for i, u := range database.Users {
			if u.ID == user.ID {
				database.Users[i] = user
				return c.JSON(http.StatusOK, user)
			}
		}
	} else {
		user.ID = faker.UUIDDigit()
	}

	database.Users = append(database.Users, user)
	return c.JSON(http.StatusCreated, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	for i, u := range database.Users {
		if u.ID == id {
			database.Users = utils.Remove[entity.User](database.Users, i)
			return c.JSON(http.StatusOK, "User deleted")
		}
	}

	return c.JSON(http.StatusNotFound, "User not found")
}

func ResetUserData(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	if num == 0 {
		num = 10
	}
	database.Users = utils.GenerateMockData[entity.User](num)
	utils.SetUsersAvatarURL()
	return c.JSON(http.StatusOK, "Data reset")
}
