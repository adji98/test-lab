package actions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/sodeve/streaming_web_user_api/models"
)

// LoginHandler is a default handler to login process

func LoginHandler(c buffalo.Context) error {
	username := c.Param("username")
	if len(username) < 1 {
		c.Flash().Add("danger", "Invalid User!")

		return c.Render(403, r.HTML("loginindex.html", "login.html"))
	}

	res, err := http.Get("http://localhost:3001/?u=" + username)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	record := models.User{}

	json.Unmarshal(body, &record)

	//log.Printf("record : %v", record)

	if record.Username == "" {
		c.Flash().Add("danger", "Invalid User!")

		return c.Render(403, r.HTML("loginindex.html", "login.html"))
	}

	c.Session().Set("user_fullname", record.Fullname)

	//log.Printf("record : %v", c.Session().Get("user_fullname"))

	return c.Redirect(302, "/main")
}
