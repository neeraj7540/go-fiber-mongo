package controllers
import (
	"github.com/kamva/mgm/v3"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go-fiber-mongo/models"
)

//GetAllUser 

func GetAllUser(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.User{})
	users := []models.User{}
	err :=collection.SimpleFind(&users, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map {
			"ok": false,
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok":true,
		"users": users,
	})

}



//Create User



func CreateUser(ctx *fiber.Ctx) {
	params :=new(struct {
	       Name string
		   Address string
		   Description string
	})

	ctx.BodyParser(&params)

	if len(params.Name) == 0 || len(params.Address) == 0 || len(params.Description) == 0 {
      ctx.Status(400).JSON(fiber.Map {
		  "ok" : false,
		  "error": "Name and Address or Description not specified.",
	  })
	  return
	}
	user := models.CreateUser(params.Name,params.Address,params.Description)
	err := mgm.Coll(user).Create(user)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map {
			"ok" : false,
			"error" : err.Error(),
		})
		return
	}
	ctx.JSON(fiber.Map {
		"ok": true,
		"user":user,
	})
}

//GetUserID

func GetUserByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	user := &models.User{}
	collection := mgm.Coll(user)
	err := collection.FindByID(id,user)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "user not found.",
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": user,
	})
}

//Upate User

func UpdateUser(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	user := &models.User{}
	collection := mgm.Coll(user)

	err := collection.FindByID(id, user)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "User not found.",
		})
		return
	}

	user.Status  =! user.Status 

	err = collection.Update(user)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"user": user,
	})
}

