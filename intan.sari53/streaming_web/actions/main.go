package actions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/sodeve/streaming_web_user_api/models"
)

// MainHandler is a default handler to serve up
// main page.
func MainHandler(c buffalo.Context) error {

	res, err := http.Get("http://localhost:3001/last_view")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	//log.Printf("API Res: %v", body)

	if err != nil {
		log.Fatal(err)
	}

	record := models.ContinueViews{}

	json.Unmarshal(body, &record)

	fullname := c.Value("user_fullname")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Session: %v", fullname)

	c.Set("videos", record)
	c.Set("hasVideo", len(record) > 0)

	return c.Render(200, r.HTML("main.html"))
}
