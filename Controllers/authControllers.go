package Controllers

import(
  "strconv"
  "os"
  "time"
  "database/sql"

  "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

  "main/Database"
  "main/Models"
  "strings"
)

const SecretKey = os.Getenv("SecretKey")

func Register(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return 
  }

  password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	if (data["email"][0] > 0 && email[len(email)-11:len(email)]=="@kiit.ac.in"){
		user := models.Student{
			First_name: data["first_name"],
			Last_name : data["last_name"],
			Email : data["email"],
			Password: password,
		}

		database.DB.Create(&user)
		return c.JSON(user)

	} else if (email[len(email)-11:len(email)]=="@kiit.ac.in") {
		user := Models.Teacher{
			First_name: data["first_name"],
			Last_name:  data["last_name"],
			Email:      data["email"],
			Password:   password,
		}

		database.DB.Create(&user)
		return c.JSON(user)
	}

	else{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message" : "Invalid credentials"
		})
	}
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
  
	if err := c.BodyParser(&data); err != nil {
		return err
	}
  
  if (data["email"][0] > 0 && strings.Contains((data["email"], "@kiit.ac.in"))){

    var user Models.Student
    
	rows, err := database.DB.Query("SELECT * FROM Student WHERE email = ?", data["email"]).First(&user)
    if err!=nil{
      return err
    }
    if rows==nil{
      c.Status(fiber.StatusNotFound)
      return c.JSON(fiber.Map{
        "message" : "Invalid email"
      })
    }
    password := database.DB.Query("SELECT password FROM Student WHERE password = ?", data["email"]).First(&user)

    err := bcrypt.CompareHashAndPassword([]byte(password), []byte(data["password"]))

    if err!=nil{
      c.Status(fiber.StatusBadRequest)
      return c.JSON({
        "message" : "Invalid password"
      })
    }

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
      Issuer: user.Email
    })
  }

  token, err := claims.SignString([]byte(SecretKey))

  if err!=nil{
    c.Status(fiber.StatusInternalServerError)
    return c.JSON(fiber.Map{
      "message" : "Could Not Login"
    })
  
    cookie := fiber.Cookie{
    Name: "jwt",
    Value: token,
    Expires: time.Now().Add(time.Hour * 24),
    HTTPOnly: true,
    }

  c.Cookie(&cookie)

  return c.JSON(fiber,Map{
    "message": "Successfully logged in"
  
  })
  }

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user Models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
