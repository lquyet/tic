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

func GetProducts(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page != 0 && pageSize != 0 {
		queriedProducts := database.Products[(page-1)*pageSize : page*pageSize]
		res := entity.GetProductsResponse{
			Data: queriedProducts,
			Pagination: entity.Pagination{
				Page:     page,
				PageSize: pageSize,
				Total:    len(database.Products),
			},
		}

		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, entity.GetProductsResponse{
		Data: database.Products,
		Pagination: entity.Pagination{
			Page:     1,
			PageSize: len(database.Products),
			Total:    len(database.Products),
		},
	})
}

func UpsertProduct(c echo.Context) error {
	product := entity.Product{}
	if err := c.Bind(&product); err != nil {
		return err
	}

	if product.ID != "" {
		for i, u := range database.Products {
			if u.ID == product.ID {
				database.Products[i] = product
				return c.JSON(http.StatusOK, product)
			}
		}
	} else {
		product.ID = faker.UUIDDigit()
	}

	database.Products = append(database.Products, product)
	return c.JSON(http.StatusCreated, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	for i, u := range database.Products {
		if u.ID == id {
			database.Products = utils.Remove[entity.Product](database.Products, i)
			return c.JSON(http.StatusOK, "Product deleted")
		}
	}

	return c.JSON(http.StatusNotFound, "Product not found")
}

func ResetProductData(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	if num == 0 {
		num = 10
	}
	database.Products = utils.GenerateMockData[entity.Product](num)
	utils.SetProductsImageURL()
	return c.JSON(http.StatusOK, "Data reset")
}
