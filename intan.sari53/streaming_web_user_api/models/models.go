package models

import (
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	if err == nil {
		DB, err = pop.Connect(env)
		if err == nil {
			pop.Debug = env == "development"
		}
	}
}
