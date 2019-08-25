package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"

	"github.com/sodeve/streaming_web_user_api/models"
)

// GetVideoHandler is a default handler to serve up
// a home page.
func GetVideoHandler(c buffalo.Context) error {

	q := models.DB.RawQuery("Select * from vw_continue_watching where 1=?", 1)
	videos := &models.ContinueViews{}
	i := 0
	q.Count(&i)

	log.Printf("count: %v", i)

	q = models.DB.RawQuery("Select * from vw_continue_watching where 1=?", 1)

	if err := q.All(videos); err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("error: %v", videos)
	return c.Render(200, r.JSON(videos))
}
