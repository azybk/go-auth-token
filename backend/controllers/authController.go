package controllers

import (
	"context"
	"go-auth-token/database"
	"go-auth-token/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SECRET_KEY = "secret"

func Register(ctx *fiber.Ctx) error {
	var data map[string]string

	// err := ctx.BodyParser(&data)
	// if err != nil {
	// 	return err
	// }
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	db := database.OpenConnection()
	defer db.Close()

	konteks := context.Background()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	script := "INSERT INTO user(name, email, password) VALUES(?,?,?)"
	_, err = tx.ExecContext(konteks, script, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	tx.Commit()
	
	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}
	
	var user models.User

	db := database.OpenConnection()
	defer db.Close()

	konteks := context.Background()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	script := "SELECT name, email, password FROM user WHERE email = ?"
	rows, err := tx.QueryContext(konteks, script, data["email"])
	defer rows.Close()

	if err != nil {
		return err
	}

	if !rows.Next() {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = rows.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	tx.Commit()

	
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: user.Email,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SECRET_KEY))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().UTC().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})

}

func User(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	db := database.OpenConnection()
	defer db.Close()

	konteks := context.Background()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	script := "SELECT name, email, password FROM user WHERE email = ?"
	rows, err := tx.QueryContext(konteks, script, claims.Issuer)
	defer rows.Close()

	if !rows.Next() {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	rows.Scan(&user.Name, &user.Email, &user.Password)

	return ctx.JSON(user)
}

func Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}