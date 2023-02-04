package controllers

import (
	"context"
	"fmt"
	"log"
	logger "gin-mongo-api/log"


	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	helper "gin-mongo-api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf(" passowrd is incorrect 🧐")
		check = false
	}

	return check, msg
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please Provide correct Format data 🤨", Data:err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please Provide correct Format data 🤨", Data:validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error 😬", Data:"error occurred while checking for the email"})
			return
		}
		if count > 0 {
			logger.WarningLogger.Println("This email is already exists")
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please login not register 😟", Data:"User with this email is already exists"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		count, err = userCollection.CountDocuments(ctx, bson.M{"contactNumber": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Error 😬", Data:"error occurred while checking for the phone number"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please login not register 😟", Data:"User with this phone is already exists"})
			return
		}

		
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Id = primitive.NewObjectID()
		user.User_id = user.Id.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken	

		//default values
		user.Roles=[]string{"USER","ADMIN","CLIENT"}
		user.Teams=[]string{}

		// c.JSON(http.StatusOK, user)

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "Error 😥", Data:msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK,responses.DefaultResponse{Status: http.StatusOK, Message: "Successful Registered here it your Details 🥳",Data:map[string]interface{}{"DatabaseResponse":resultInsertionNumber}})

	}
}

// Login is the api used to to get a single user by email
func LoginWithEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.LoginWithEmail

		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please Provide correct Format data 🤨", Data:err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusNotFound,responses.ErrorResponse{Status: http.StatusNotFound, Message: "User not found on this email ☹️",Data:err.Error()})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusBadRequest,responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please provide correct password 🧐",Data:msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK,responses.DefaultResponse{Status: http.StatusOK, Message: "Successful login here it your Details 🥳",Data:map[string]interface{}{"token":foundUser.Token,"refresh_token":foundUser.Refresh_token}})

	}
}

// Login is the api used to to get a single user by phone number
func LoginWithMobile() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.LoginWithMobile
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please Provide correct Format data 🤨", Data:err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"contactNumber": user.Phone}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusNotFound,responses.ErrorResponse{Status: http.StatusNotFound, Message: "User not found on this email ☹️",Data:err.Error()})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if passwordIsValid != true {
			c.JSON(http.StatusBadRequest,responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please provide correct password 🧐",Data:msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		c.JSON(http.StatusOK,responses.DefaultResponse{Status: http.StatusOK, Message: "Successful login here it your Details 🥳",Data:map[string]interface{}{"token":foundUser.Token,"refresh_token":foundUser.Refresh_token}})


	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "User not found ☹️", Data:err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: user})
	}
}

func EditAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "Please Provide correct Format data 🤨", Data:err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "error", Data:validationErr.Error()})
			return
		}

		update := bson.M{"firstName": user.First_name,"lastName":user.Last_name, "location": user.Location, "title": user.Title,"address":user.Address,"age":user.Age}
		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//get updated user details
		var updatedUser models.User
		defer cancel()
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			} else {
				c.JSON(http.StatusOK,
					responses.DefaultResponse{Status: http.StatusOK, Message: "success 😃", Data:updatedUser},
				)
			}
		} else {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error ☹️", Data: "Invalid ID"})
		}

	}
}

func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.ErrorResponse{Status: http.StatusNotFound, Message: "error", Data: "User with specified ID not found!"},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: "User successfully deleted!"},
		)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			}

			users = append(users, singleUser)
		}
		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: users})
        
	}
}
