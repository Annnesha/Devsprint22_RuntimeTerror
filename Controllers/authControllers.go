package Controllers

import(
  "os"
  "time"
  "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

  "main/Database"
  "main/Models"
  "strings"
)

var SecretKey string = os.Getenv("SecretKey")


func Register(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  Password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	if (data["Email"][0] > 0 && data["Email"][len(data["Email"])-11:len(data["Email"])]=="@kiit.ac.in"){
		user := Models.Student{
			First_Name: data["FirstName"],
			Last_Name : data["LastName"],
			Email : data["Email"],
      Roll : data["ID"],
			Password: Password,
      Branch: data["School"],
		}

		Database.DB.Create(&user)
		return c.JSON(user)

	} else if (data["Email"][len(data["Email"])-11:len(data["Email"])]=="@kiit.ac.in") {
		user := Models.Teacher{
			First_Name: data["FirstName"],
			Last_Name:  data["LastName"],
			Email:      data["Email"],
      Faculty_ID: data["ID"],
      Position: data["Designation"],
      School: data["School"],
			Password:   Password,
		}

		Database.DB.Create(&user)
		return c.JSON(user)
	} else{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message" : "Invalid credentials",
		})
	}
}


func Login(c *fiber.Ctx) error {
	var data map[string]string
  
	if err := c.BodyParser(&data); err != nil {
		return err
	}
  
  if (data["Email"][0] > 0 && strings.Contains((data["Email"]), "@kiit.ac.in")){

    var user Models.Student
    
	rows:= Database.DB.Select("Password").Where("Email like ?", data["Email"]).Table("Student").First(&user)
    if rows==nil{
      c.Status(fiber.StatusNotFound)
      return c.JSON(fiber.Map{
        "message" : "Invalid email",
      })
    }
   if err:=bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["Password"])); err!=nil{
      c.Status(fiber.StatusBadRequest)
      return c.JSON(fiber.Map{
        "message" : "Invalid password",
      })
    }

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
      Issuer: user.Email,
    })

  token, err := claims.SignedString([]byte(SecretKey))

  if err!=nil{
    c.Status(fiber.StatusInternalServerError)
    return c.JSON(fiber.Map{
      "message" : "Could Not Login",
    })
  }
  
    cookie := fiber.Cookie{
    Name: "jwt",
    Value: token,
    Expires: time.Now().Add(time.Hour * 24),
    HTTPOnly: true,
    }

  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "message": "Successfully logged in",
  })
  }
  

  if (strings.Contains((data["Email"]),"@kiit.ac.in")){

    var user Models.Teacher

  rows := Database.DB.Select("Password").Where("Email like ?", data["Email"]).Table("Teacher").First(&user)
    if rows==nil{
      c.Status(fiber.StatusNotFound)
      return c.JSON(fiber.Map{
        "message" : "Invalid email",
      })
    }

  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["Password"])); err!=nil{
      c.Status(fiber.StatusBadRequest)
      return c.JSON(fiber.Map{
        "message" : "Invalid password",
      })
    }

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
      Issuer: user.Email,
    })

  token, err := claims.SignedString([]byte(SecretKey))

  if err!=nil{
    c.Status(fiber.StatusInternalServerError)
    return c.JSON(fiber.Map{
      "message" : "Could Not Login",
    })
  }
  
    cookie := fiber.Cookie{
    Name: "jwt",
    Value: token,
    Expires: time.Now().Add(time.Hour * 24),
    HTTPOnly: true,
    }

  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "message": "Successfully logged in",
  
  })
  }

  return c.JSON(fiber.Map{
    "message" : "Unknown error occurred",
  })

}

func Student(c *fiber.Ctx) error {
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

	var user Models.Student

	Database.DB.Where("Email = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Teacher(c *fiber.Ctx) error {
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

	var user Models.Teacher

	Database.DB.Where("Email = ?", claims.Issuer).First(&user)

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
