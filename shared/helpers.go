package shared

import (
	"context"

	"github.com/gookit/color"
)

const (
	CONTENT_TYPE_JSON        = "application/json"
	CONTENT_TYPE_URL_ENCODED = "application/x-www-form-urlencoded"
	USER_AGENT               = "aws-cli/2.5.3 Python/3.9.12"
	CONTENT_TYPE_HEADER      = "Content-Type"
)

type Scanner interface {
	EnumerateAll(ctx context.Context, region string, creds interface{}) error
	EnumerateSpecificService(ctx context.Context, region, service string, creds interface{}) error
}

func GetMessageColor(p string) string {
	switch p {
	case "error":
		return Red("●")
	case "warning":
		return Yellow("●")
	case "debug", "low":
		return Blue("●")
	case "success":
		return Green("●")
		// return Green("✓")
	default:
		return ""
	}
}

func Red(s string) string {
	return color.Red.Sprintf("%s", s)
}

func Cyan(s string) string {
	return color.Cyan.Sprintf("%s", s)
}

func Yellow(s string) string {
	return color.Yellow.Sprintf("%s", s)
}

func Green(s string) string {
	return color.Green.Sprintf("%s", s)
}

func Blue(s string) string {
	return color.Blue.Sprintf("%s", s)
}

func Magenta(s string) string {
	return color.Magenta.Sprintf("%s", s)
}

func Black(s string) string {
	return color.Black.Sprintf("%s", s)
}
