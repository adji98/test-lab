package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sodeve/streaming_web/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
