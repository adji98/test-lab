package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sodeve/streaming_web_recommendation_api/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
