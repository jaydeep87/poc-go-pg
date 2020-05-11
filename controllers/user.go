package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
)

type User struct {
	ID        string    `form:"id" json:"id"`
	Name      string    `form:"name" json:"name"`
	Email     string    `form:"email" json:"email"`
	Desc      string    `form:"desc" json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create User Table
func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&User{}, opts)
	if createError != nil {
		log.Printf("Error while creating  user table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("User table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllUsers(c *gin.Context) {
	var users []User
	err := dbConnect.Model(&users).Select()

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    users,
	})
	return
}

func CreateUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	name := user.Name
	email := user.Email
	desc := user.Desc
	id := guuid.New().String()

	insertError := dbConnect.Insert(&User{
		ID:        id,
		Name:      name,
		Email:     email,
		Desc:      desc,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new User into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
	return
}

func GetSingleUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &User{ID: userId}
	err := dbConnect.Select(user)

	if err != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single User",
		"data":    user,
	})
	return
}

func EditUser(c *gin.Context) {
	userId := c.Param("userId")
	var user User
	c.Bind(&user)
	desc := user.Desc

	_, err := dbConnect.Model(&User{}).Set("desc = ?", desc).Where("id = ?", userId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User Edited Successfully",
	})
	return
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &User{ID: userId}

	err := dbConnect.Delete(user)
	if err != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
	return
}
