package auth

import (
	"ModaLast/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gookit/validate"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
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
				return c.JSON(http.StatusBadRequest, echo.Map{
					"success": false,
					"errors":  "Register error",
				})
			}

			// Create token with claims
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = user.FullName
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, echo.Map{
				"success": true,
				"data":    user,
				"token":   t,
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"success": false,
				"errors":  v.Errors.String(),
			})
		}
	}
}

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		user := models.User{}

		email := c.FormValue("email")
		password := c.FormValue("password")

		userErr := db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if userErr != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"errors":  "Register error",
				"success": false,
			})
		}

		err = models.VerifyPassword(user.Password, password)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Create token with claims
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.FullName
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"success": true,
			"token":   t,
		})
	}
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
