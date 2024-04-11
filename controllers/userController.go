package controllers

import (
	"example/testapp/initializers"
	"example/testapp/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UserCreate(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	initializers.DB.Db.Create(&user)
	return c.Status(202).JSON(user)
}

func GetUserById(c *fiber.Ctx) error {
	user := new(models.User)
	fmt.Println(user, "user--------")
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Db.Where("id = ?", user.ID).Find(&user)

	return c.Status(200).JSON(user)
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	initializers.DB.Db.Find(&users)
	fmt.Println(c.App().Stack(), "-------------")
	// return c.Status(200).JSON(users)
	return c.JSON(c.BaseURL())
}

func UpDateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Check if the record exists
	existingUser := new(models.User)
	result := initializers.DB.Db.First(existingUser, user.ID)
	if result.Error != nil {
		return c.Status(404).JSON("User not found")
	}

	// Update the user record
	result = initializers.DB.Db.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user)
	if result.Error != nil {
		return c.Status(500).JSON(result.Error.Error())
	}

	return c.Status(200).JSON(user)
}

func Delete(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	initializers.DB.Db.Unscoped().Where("id= ?", user.ID).Delete(&user)
	return c.Status(200).JSON("Delete success")
}

func GetSynonyms(c *fiber.Ctx) error {
	word := new(models.Word)
	searchTerm := c.Query("term")

	// Perform the database query to find the word
	result := initializers.DB.Db.Where("term = ?", searchTerm).First(word)

	// Check if the word was found
	if result.Error != nil {
		// Word not found, return an appropriate response
		return c.Status(404).JSON(fiber.Map{"error": "Word not found"})
	}
	fmt.Print(word)
	// Define a slice to hold the synonyms
	var synonyms []models.Word

	// Perform a raw SQL query to fetch synonyms of the word
	rows, err := initializers.DB.Db.Raw(`

	select * from words w 
	where w.id in (
	select word_id2 from synonyms s 
		where s.word_id1 = (	
		select s2.word_id1  from synonyms s2 
		where s2.word_id2 = ?
		)
	) and w.term != ?
    `, word.ID, searchTerm).Rows()

	// Check for errors during the query
	if err != nil {
		// Return an error response
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	// Iterate over the rows and scan them into the slice
	for rows.Next() {
		var synonym models.Word
		initializers.DB.Db.ScanRows(rows, &synonym)
		synonyms = append(synonyms, synonym)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		// Return an error response
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the list of synonyms
	return c.Status(200).JSON(synonyms)
}
