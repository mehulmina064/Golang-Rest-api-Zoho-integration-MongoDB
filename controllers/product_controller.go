package controllers

import (
	"context"
	"encoding/json"
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

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")


func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var Product models.Product

	

		if err := c.BindJSON(&Product); err != nil {
			logger.ErrorLogger.Println("error in creating Product")

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(Product)
		if validationErr != nil {
			logger.ErrorLogger.Println("error in creating Product")
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		
		Product.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		Product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		

		Product.Id = primitive.NewObjectID()
		Product.Product_id = Product.Id.Hex()


		// msg := fmt.Sprintf("Product  was not created")  //for testing purposes
        // logger.InfoLogger.Println(msg)  //for testing purposes
		// c.JSON(http.StatusOK, Product)    //for testing purposes
	



		resultInsertionNumber, insertErr := productCollection.InsertOne(ctx, Product)
		if insertErr != nil {
			msg := fmt.Sprintf("Product  was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)

	}
}




func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("productId")
		var Product models.Product
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := productCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&Product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: Product})
	}
}

func EditProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ProductId := c.Param("productId")
		var Product models.Product
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&Product); err != nil {
			c.JSON(http.StatusBadRequest, responses.DefaultResponse{Status: http.StatusBadRequest, Message: "error", Data:err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&Product); validationErr != nil {
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

		Product.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		log.Println("id"+adminId)
		// log.Println(Product.Admins)
		

		objId, _ := primitive.ObjectIDFromHex(ProductId)
		

		update := bson.M{
			"name": Product.NAME,
			"pimcoreId": Product.PimcoreId,
			"Brand":Product.Brand,
			"manufacturer": Product.Manufacturer,
			"categoryId": Product.CategoryId,
			"categoryName": Product.CategoryName,
			"hsnOrSac": Product.HsnOrSac,
			"imageName": Product.ImageName,
			"imageType": Product.ImageType,
			"status": Product.Status,
			"isLinkedWithZohoCrm": Product.IsLinkedWithZohoCrm,
			"zohoCrmProductId": Product.ZohoCrmProductId,
			"crmOwnerId": Product.CrmOwnerId,
			"unit": Product.Unit,
			"description": Product.Description,
			"itemTaxPreferences": Product.ItemTaxPreferences,
			"rate": Product.Rate,
			"accountId": Product.AccountId,
			"accountName":Product.AccountName,
			"taxId": Product.TaxId,
			"taxName": Product.TaxName,
            "taxPercentage": Product.TaxPercentage,
			"taxType": Product.TaxType,
			"isDefaultTaxApplied": Product.IsDefaultTaxApplied,
			"isTaxable": Product.IsTaxable,
			"taxExemptionId": Product.TaxExemptionId,
			"taxExemptionCode": Product.TaxExemptionCode,
			"taxabilityType": Product.TaxabilityType,
			"documents": Product.Documents,
			"purchaseDescription": Product.PurchaseDescription,
			"priceBookRate": Product.PriceBookRate,
			"salesRate": Product.SalesRate,
			"purchaseRate": Product.PurchaseRate,
			"purchaseAccountId": Product.PurchaseAccountId,
			"purchaseAccountName": Product.PurchaseAccountName,
			"inventoryAccountId": Product.InventoryAccountId,
			"inventoryAccountName": Product.InventoryAccountName,
			"tags": Product.Tags,
			"itemType": Product.ItemType,
			"productType": Product.ProductType,
			"isReturnable": Product.IsReturnable,
			"reorderLevel": Product.ReorderLevel,
			"minimumOrderQuantity": Product.MinimumOrderQuantity,
			"maximumOrderQuantity": Product.MaximumOrderQuantity,
			"initialStock": Product.InitialStock,
			"initialStockRate": Product.InitialStockRate,
			"totalInitialStock": Product.TotalInitialStock,
			"stockOnHand": Product.StockOnHand,
			"vendorId": Product.VendorId,
			"vendorName": Product.VendorName,
            "assetValue": Product.AssetValue,
			"availableStock": Product.AvailableStock,
			"actualAvailableStock": Product.ActualAvailableStock,
			"committedStock": Product.CommittedStock,
            "actualCommittedStock":Product.ActualCommittedStock,
			"availableForSaleStock": Product.AvailableForSaleStock,
			"trackBatchNumber": Product.TrackBatchNumber,
            "isComboProduct":Product.IsComboProduct,
			"isAdvancedTrackingMissing": Product.IsAdvancedTrackingMissing,
			"customFields": Product.CustomFields,
			"salesChannels": Product.SalesChannels,
			"warehouses": Product.Warehouses,
			"branches": Product.Branches,
			"preferredVendors": Product.PreferredVendors,
			"packageDetails": Product.PackageDetails,
			"product_id": Product.Product_id,
			"zohoCreatedTime": Product.ZohoCreatedTime,
			"zohoLastModifiedTime": Product.ZohoLastModifiedTime,
			"updated_at": Product.Updated_at,
		}
		result, err := productCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		//get updated user details
		var updatedProduct models.Product 
		defer cancel()
		if result.MatchedCount == 1 {
			err := productCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedProduct)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
				return
			} else {
				c.JSON(http.StatusOK,
					responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data:updatedProduct},
				)
			}
		} else {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data: "Invalid ID"})
		}
		// emp, _ := json.Marshal(updatedProduct)
        bytes:=[]byte("")
		log.Println(json.Unmarshal(bytes,&updatedProduct))
	}

}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ProductId := c.Param("productId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ProductId)

		result, err := productCollection.DeleteOne(ctx, bson.M{"_id": objId})

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

func GetAllProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("in Products")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var Products []models.Product
		defer cancel()

		results, err := productCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			return
		}
		// c.JSON(http.StatusOK,results)
		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleProduct models.Product
			if err = results.Decode(&singleProduct); err != nil {
				c.JSON(http.StatusInternalServerError, responses.DefaultResponse{Status: http.StatusInternalServerError, Message: "error", Data:err.Error()})
			}

			Products = append(Products, singleProduct)
		}
		c.JSON(http.StatusOK, responses.DefaultResponse{Status: http.StatusOK, Message: "success", Data: Products})
        
	}
}
