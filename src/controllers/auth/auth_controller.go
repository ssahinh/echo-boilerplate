package auth

import (
	"ModaLast/src/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gookit/validate"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Register(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		user := models.User{FullName: c.FormValue("fullname"), Email: c.FormValue("email"),
			Password: c.FormValue("password")}
		v := validate.Struct(user)

		if v.Validate() {
			err = db.Debug().Model(&models.User{}).Create(&user).Error
			if err != nil {
				return c.JSON(http.StatusBadRequest, "Register Error")
			}

			return c.JSON(http.StatusOK, user)
		} else {
			return c.JSON(http.StatusBadRequest, v.Errors)
		}
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		user := models.User{}

		email := c.FormValue("email")
		password := c.FormValue("password")

		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			log.Println(err)
			return err
		}

		err = models.VerifyPassword(user.Password, password)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println(err)
			return err
		}

		// set custom claims
		claims := &jwtCustomClaims{
			"Jon Snow",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
