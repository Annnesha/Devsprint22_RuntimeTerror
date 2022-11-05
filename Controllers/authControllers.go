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

const SecretKey