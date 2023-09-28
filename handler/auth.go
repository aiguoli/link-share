package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"link-share/database"
	"link-share/models"
	"os"
	"time"
)

type signupInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	db := database.DB
	req := new(signupInput)
	if err := c.BodyParser(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't parse JSON",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't hash password",
		})
	}
	user := &models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}
	db.Create(&user)
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "User created successfully",
		"data": user,
	})
}

func Login(c *fiber.Ctx) error {
	db := database.DB
	req := new(loginInput)
	if err := c.BodyParser(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't parse JSON",
		})
	}
	user := new(models.User)
	db.Where("email = ?", req.Email).First(&user)

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code": 401,
			"msg":  "User not found",
		})
	}
	if !CheckPasswordHash(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code": 401,
			"msg":  "Incorrect password",
		})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.ID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	s, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": 500,
			"msg":  "Couldn't generate token",
		})
	}
	return c.JSON(fiber.Map{
		"code":    200,
		"token":   s,
		"expires": time.Now().Add(time.Hour * 72),
	})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
