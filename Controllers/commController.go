package Controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"main/Database"
	"main/Models"
)

func Apply(c *fiber.Ctx) error {

	var data map[string]string

	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	event := Models.Events{
		Title:            data["Title"],
		Organizer:        data["Organizer"],
		Student_Name:     data["Name"],
		Roll:             data["Roll"],
		Branch:           data["Branch"],
		School_concerned: data["School"],
		Description:      data["Desc"],
		Venue:            data["Venue"],
		Start_date:       data["StartDate"],
		End_date:         data["EndDate"],
		Start_time:       data["StartTime"],
		End_time:         data["EndTime"],
		Status:           "Pending"}

	Database.DB.Create(&event)
	return c.JSON(event)
}
