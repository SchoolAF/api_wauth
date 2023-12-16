package handler
import (
	"api/database"
	"api/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//Create a user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
   // Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
	 return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Something's wrong with your input", "data": err})
	}
   err = db.Create(&user).Error
	if err != nil {
	 return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Could not create user", "data": err})
	}
   // Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message":  "User has created", "data": user})
}

// Get All Users from db
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User
   // find all users in the database
	db.Find(&users)
   // If no user found, return an error
	if len(users) == 0 {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
   // return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

// GetSingleUser from db
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db
   // get id params
	id := c.Params("id")
   var user model.User
   // find single user in the database by id
	db.Find(&user, "id = ?", id)
   if user.ID == uuid.Nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
   return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
	 Username string `json:"username"`
	}
   db := database.DB.Db
   var user model.User
   // get id params
	id := c.Params("id")
   // find single user in the database by id
	db.Find(&user, "id = ?", id)
   if user.ID == uuid.Nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
   var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
	 return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
   user.Username = updateUserData.Username
   // Save the Changes
	db.Save(&user)
   // Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
   // get id params
	id := c.Params("id")
   // find single user in the database by id
	db.Find(&user, "id = ?", id)
   if user.ID == uuid.Nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
   }
   err := db.Delete(&user, "id = ?", id).Error
   if err != nil {
	 return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
   return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

func CheckPhoneNumber(c *fiber.Ctx) error {
    db := database.DB.Db

    phoneNumber := c.Params("phoneNumber")

    if phoneNumber == "" {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Phone number is required", "data": nil})
    }

    var user model.User

    // Check if the phone number exists in the database
    if err := db.Where("phone = ?", phoneNumber).First(&user).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error checking phone number", "data": err})
    }

    // If the user is found, respond with true; otherwise, respond with false
    isRegistered := user.ID != uuid.Nil
    return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Phone number registration status", "data": isRegistered})
}