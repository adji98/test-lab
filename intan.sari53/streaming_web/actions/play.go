package actions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/sodeve/streaming_web/models"
)

// PlayHandler is a default handler to serve up
// main page.
func PlayHandler(c buffalo.Context) error {

	res, err := http.Get("http://localhost:3001/get_video?v" + c.Param("v") + "&q=" + c.Param("q"))
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

	log.Printf("parsed JSON: %v", record)

	c.Set("videos", record)
	c.Set("hasVideo", len(record) > 0)

	c.Set("video_id", c.Param("v"))
	c.Set("quality_id", c.Param("q"))

	return c.Render(200, r.HTML("play.html"))
}
