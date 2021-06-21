package users

import (
	"fmt"
	"net/http"

	"github.com/commy08/golang_sandbox.git/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

var DB *gorm.DB

func (UsersRouter) register(c echo.Context) error {
	type FromUser struct {
		Username    string `json:"username"`
		Firstname   string `json:"firstname"`
		Lastname    string `json:"lastname"`
		Age         int    `json:"age"`
		DateOfBirth string `json:"date_of_birth"`
	}

	var body FromUser

	if err := c.Bind(&body); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "something went wrong!")
	}

	// DateOfBirth, err := time.Parse("2006-01-02", body.DateOfBirth)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.String(http.StatusInternalServerError, "something went wrong!")
	// }

	result := DB.Find(&models.User{})

	fmt.Printf("%v \n", result)

	return c.JSON(http.StatusCreated, echo.Map{
		"massage": "register complete",
	})
}
