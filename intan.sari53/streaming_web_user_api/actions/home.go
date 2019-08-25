package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/sodeve/streaming_web_user_api/models"
)

func HomeHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	username := c.Param("u")

	q := models.DB.RawQuery("Select * from Users where username=?", username)

	user := &models.User{}

	if err := q.Find(user, username); err != nil {
		log.Printf("error: %v", err)
	}

	return c.Render(200, r.JSON(user))

}
