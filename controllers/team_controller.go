package controllers

import (
	"context"
	"fmt"
	"log"

	helper "gin-mongo-api/helpers"
	logger "gin-mongo-api/log"

	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var teamCollection *mongo.Collection = configs.GetCollection(configs.DB, "teams")


func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func CreateTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var team models.Team

		clientToken := c.Request.Header.Get("token")
        if clientToken == "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
            c.Abort()
            return
        }
        claims, err := helper.ValidateToken(clientToken)
        if err != "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            c.Abort()
            return
        }
		//get admin user by token
		adminId := claims.Uid

		if err := c.BindJSON(&team); err != nil {
			logger.ErrorLogger.Println("error in creating team")

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(team)
		if validationErr != nil {
			logger.ErrorLogger.Println("error in creating team")
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		
		team.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		team.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		team.Admins=[]string{adminId}
		team.Users=append(team.Users,adminId) 
		team.Admins=removeDuplicateStr(team.Admins)
		team.Users=removeDuplicateStr(team.Users)

		team.Id = primitive.NewObjectID()
		team.Team_id = team.Id.Hex()


		// msg := fmt.Sprintf("Team  was not created")  //for testing purposes
        // logger.InfoLogger.Println(msg)  //for testing purposes
		// c.JSON(http.StatusOK, team)    //for testing purposes
	



		resultInsertionNumber, insertErr := teamCollection.InsertOne(ctx, team)
		if insertErr != nil {
			msg := fmt.Sprintf("Team  was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}




func GetTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("teamId")
		var team models.Team
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := teamCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&team)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: team})
	}
}

func EditTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		teamId := c.Param("teamId")
		var team models.Team
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&team); err != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data:err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&team); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data:validationErr.Error()})
			return
		}

		clientToken := c.Request.Header.Get("token")
        if clientToken == "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
            c.Abort()
            return
        }
        claims, err1 := helper.ValidateToken(clientToken)
        if err1 != "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err1})
            c.Abort()
            return
        }
		// get admin user by token
		adminId := claims.Uid

		team.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		log.Println("id"+adminId)
		// log.Println(team.Admins)
		team.Admins=append(team.Admins,adminId) 
		team.Users=append(team.Users,adminId) 
		// log.Println(team.Admins)
		team.Admins=removeDuplicateStr(team.Admins)
		team.Users=removeDuplicateStr(team.Users)

		objId, _ := primitive.ObjectIDFromHex(teamId)
		

		update := bson.M{"name": team.NAME,"description":team.Description, "admins": team.Admins, "users": team.Users}
		result, err := teamCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//get updated user details
		var updatedTeam models.Team 
		defer cancel()
		if result.MatchedCount == 1 {
			err := teamCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedTeam)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			} else {
				c.JSON(http.StatusOK,
					responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data:updatedTeam},
				)
			}
		} else {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: "Invalid ID"})
		}
	}
}

func DeleteTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		teamId := c.Param("teamId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(teamId)

		result, err := teamCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.DefaultResponse{Status: http.StatusNotFound, Message: "error", Data: "User with specified ID not found!"},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: "User successfully deleted!"},
		)
	}
}

func GetAllTeams() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var teams []models.Team
		defer cancel()

		results, err := teamCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleTeam models.Team
			if err = results.Decode(&singleTeam); err != nil {
				c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			}

			teams = append(teams, singleTeam)
		}
		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: teams})
        
	}
}
