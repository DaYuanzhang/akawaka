package utils

import (
	"github.com/gookit/color"
	"strings"
)

type Color struct {
	Green func(a ...any) string
	White func(a ...any) string
	Red   func(a ...any) string
}

var LogColor *Color

func init() {
	if LogColor == nil {
		LogColor = NewColor()
	}
}

func NewColor() *Color {
	return &Color{
		Green: color.FgLightGreen.Render,
		White: color.FgLightWhite.Render,
		Red:   color.FgLightRed.Render,
	}
}

func (c *Color) GetColor(color string, msg string) string {
	color = strings.ToLower(color)
	switch color {
	case "green":
		return c.Green(msg)
	case "white":
		return c.White(msg)
	case "red":
		return c.Red(msg)
	default:
		return c.White(msg)
	}
}
