package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anujdatar/goth-stack/database"
	"github.com/anujdatar/goth-stack/models"
	"github.com/anujdatar/goth-stack/utils"
	"github.com/gin-gonic/gin"
)

func RegisterNewUser(c *gin.Context) {
	var input models.UserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error. Unable to bind JSON: " + err.Error(),
		})
		return
	}

	if input.Email == "" ||
		input.Password == "" ||
		input.FirstName == "" ||
		input.LastName == "" ||
		input.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error. Missing required fields",
		})
		return
	}
	if input.Username == "" {
		input.Username = input.Email
	}

	if input.AdminPassword == os.Getenv("ADMIN_PASSWORD") {
		input.Role = "admin"
	} else {
		input.Role = "user"
	}

	input.PhoneVerification = utils.GenerateRandomNumber(6)
	input.EmailVerification = utils.GenerateRandomString(6)
	input.AccountState = "active"
	input.IncorrectPwdCount = 0
	input.PwdResetFlag = 0
	input.PwdResetCode = ""
	input.InviteCode = ""

	fmt.Println(&input)

	NewUser(&input)

	c.JSON(http.StatusOK, gin.H{
		"message": "creating new user",
	})
}

func NewUser(u *models.UserInput) {
	fmt.Println("new user")

	currentTime := time.Unix(0, time.Now().UnixNano())
	u.CreatedAt = currentTime
	u.UpdatedAt = currentTime

	_, err := database.Database.Exec(
		database.NewUserQueryStr(),
		u.FirstName,
		u.LastName,
		u.Username,
		u.Email,
		u.Password,
		u.Phone,
		u.Role,
		u.AccountState,
		u.IncorrectPwdCount,
		u.PwdResetFlag,
		u.PwdResetCode,
		u.InviteCode,
		u.PhoneVerification,
		u.EmailVerification,
		u.CreatedAt,
		u.UpdatedAt,
	)
	if err != nil {
		log.Fatal("Error. Unable to create new user: ", err)
	}
}

func UpdateUser() {
	fmt.Println("update user")
}

func DeleteUser() {
	fmt.Println("delete user")
}
