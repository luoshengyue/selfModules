package recipes_api

// Recipes API
//
// This is a sample recipes API.
// You can find out more about the API at https://github.com/luoshengyue/gin_project.
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: luoyuehe <bancangbaize@gmail.com> https://sodasweet.cn
//
// Consumes:
// - application/json
// Produces:
// - application/json
// swagger:meta

//import (
//	"context"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"github.com/go-redis/redis"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"log"
//	"net/http"
//	"time"
//)
//
//type RecipesHandler struct {
//	collection  *mongo.Collection
//	ctx         context.Context
//	redisClient *redis.Client
//}
//
//type Recipe struct {
//	//swagger:ignore
//	//ID           primitive.ObjectID `json:"id" bson:"_id"`
//	ID           string    `json:"id"`
//	Name         string    `json:"name" bson:"name"`
//	Tags         []string  `json:"tags" bson:"tags"`
//	Ingredients  []string  `json:"ingredients" bson:"ingredients"`
//	Instructions []string  `json:"instructions" bson:"instructions"`
//	PublishedAt  time.Time `json:"publishedAt" bson:"publishedAt"`
//}
//
//func NewRecipesHandler(ctx context.Context, collection *mongo.Collection, redisClient *redis.Client) *RecipesHandler {
//	return &RecipesHandler{
//		collection:  collection,
//		ctx:         ctx,
//		redisClient: redisClient,
//	}
//}
//
//// swagger:operation GET /recipes recipes listRecipes
//// Returns list of recipes
//// ---
//// produces:
//// - application/json
//// responses:
////     '200':
////         description: Successful operation
//func (handler *RecipesHandler) ListRecipesHandler(c *gin.Context) {
//	val, err := handler.redisClient.Get("recipes").Result()
//	if err == redis.Nil {
//		log.Printf("Request to MongoDB")
//		cur, err := handler.collection.Find(handler.ctx, bson.M{})
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			return
//		}
//		defer cur.Close(handler.ctx)
//
//		recipes := make([]Recipe, 0)
//		for cur.Next(handler.ctx) {
//			var recipe Recipe
//			cur.Decode(&recipe)
//			recipes = append(recipes, recipe)
//		}
//
//		data, _ := json.Marshal(recipes)
//		handler.redisClient.Set("recipes", string(data), 0)
//		c.JSON(http.StatusOK, recipes)
//	} else if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	} else {
//		log.Printf("Request to Redis")
//		recipes := make([]Recipe, 0)
//		json.Unmarshal([]byte(val), &recipes)
//		c.JSON(http.StatusOK, recipes)
//	}
//}
//
//// swagger:operation POST /recipes recipes newRecipe
//// Create a new recipe
//// ---
//// produces:
//// - application/json
//// responses:
////     '200':
////         description: Successful operation
////     '400':
////         description: Invalid input
//func (handler *RecipesHandler) NewRecipeHandler(c *gin.Context) {
//	var recipe Recipe
//	if err := c.ShouldBindJSON(&recipe); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	recipe.ID = primitive.NewObjectID()
//	recipe.PublishedAt = time.Now()
//	_, err := handler.collection.InsertOne(handler.ctx, recipe)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new recipe"})
//		return
//	}
//
//	log.Println("Remove data from Redis")
//	handler.redisClient.Del("recipes")
//
//	c.JSON(http.StatusOK, recipe)
//}
//
//// swagger:operation PUT /recipes/{id} recipes updateRecipe
//// Update an existing recipe
//// ---
//// parameters:
//// - name: id
////   in: path
////   description: ID of the recipe
////   required: true
////   type: string
//// produces:
//// - application/json
//// responses:
////     '200':
////         description: Successful operation
////     '400':
////         description: Invalid input
////     '404':
////         description: Invalid recipe ID
//func (handler *RecipesHandler) UpdateRecipeHandler(c *gin.Context) {
//	id := c.Param("id")
//	var recipe Recipe
//	if err := c.ShouldBindJSON(&recipe); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	objectId, _ := primitive.ObjectIDFromHex(id)
//	_, err := handler.collection.UpdateOne(handler.ctx, bson.M{
//		"_id": objectId,
//	}, bson.D{{"$set", bson.D{
//		{"name", recipe.Name},
//		{"instructions", recipe.Instructions},
//		{"ingredients", recipe.Ingredients},
//		{"tags", recipe.Tags},
//	}}})
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been updated"})
//}
//
//// swagger:operation DELETE /recipes/{id} recipes deleteRecipe
//// Delete an existing recipe
//// ---
//// produces:
//// - application/json
//// parameters:
////   - name: id
////     in: path
////     description: ID of the recipe
////     required: true
////     type: string
//// responses:
////     '200':
////         description: Successful operation
////     '404':
////         description: Invalid recipe ID
//func (handler *RecipesHandler) DeleteRecipeHandler(c *gin.Context) {
//	id := c.Param("id")
//	objectId, _ := primitive.ObjectIDFromHex(id)
//	_, err := handler.collection.DeleteOne(handler.ctx, bson.M{
//		"_id": objectId,
//	})
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been deleted"})
//}
//
//// swagger:operation GET /recipes/{id} recipes
//// Get one recipe
//// ---
//// produces:
//// - application/json
//// parameters:
////   - name: id
////     in: path
////     description: recipe ID
////     required: true
////     type: string
//// responses:
////     '200':
////         description: Successful operation
//func (handler *RecipesHandler) GetOneRecipeHandler(c *gin.Context) {
//	id := c.Param("id")
//	objectId, _ := primitive.ObjectIDFromHex(id)
//	cur := handler.collection.FindOne(handler.ctx, bson.M{
//		"_id": objectId,
//	})
//	var recipe Recipe
//	err := cur.Decode(&recipe)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, recipe)
//}

// swagger:operation GET /recipes/search recipes findRecipe
// Search recipes based on tags
// ---
// produces:
// - application/json
// parameters:
//   - name: tag
//     in: query
//     description: recipe tag
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
/*func SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)
	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
		}
	}
	c.JSON(http.StatusOK, listOfRecipes)
}*/

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes []Recipe

// swagger:parameters recipes newRecipe
type Recipe struct {
	//swagger:ignore
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// swagger:operation POST /recipes recipes newRecipe
// Create a new recipe
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()

	recipes = append(recipes, recipe)

	c.JSON(http.StatusOK, recipe)
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// Update an existing recipe
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the recipe
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
//     '404':
//         description: Invalid recipe ID
func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	recipes[index] = recipe

	c.JSON(http.StatusOK, recipe)
}

// swagger:operation DELETE /recipes/{id} recipes deleteRecipe
// Delete an existing recipe
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Invalid recipe ID
func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	recipes = append(recipes[:index], recipes[index+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been deleted"})
}

// swagger:operation GET /recipes/search recipes findRecipe
// Search recipes based on tags
// ---
// produces:
// - application/json
// parameters:
//   - name: tag
//     in: query
//     description: recipe tag
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
func SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)

	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}

// swagger:operation GET /recipes/{id} recipes oneRecipe
// Get one recipe
// ---
// produces:
// - application/json
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
// responses:
//     '200':
//         description: Successful operation
//     '404':
//         description: Invalid recipe ID
func GetRecipeHandler(c *gin.Context) {
	id := c.Query("id")
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			c.JSON(http.StatusOK, recipes[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
}
