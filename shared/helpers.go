package shared

import (
	"context"
	"regexp"
	"strings"

	"github.com/gookit/color"
)

const (
	CONTENT_TYPE_JSON        = "application/json"
	CONTENT_TYPE_URL_ENCODED = "application/x-www-form-urlencoded"
	USER_AGENT               = "aws-cli/2.5.3 Python/3.9.12"
	CONTENT_TYPE_HEADER      = "Content-Type"
)

var TemplatePropertyRegex = regexp.MustCompile(`\{.(\w+)\}`)

type Scanner interface {
	EnumerateAll(ctx context.Context, region string, creds interface{}) error
	EnumerateSpecificService(ctx context.Context, region, service string, creds interface{}) error
}

func GetMessageColor(p string) string {
	switch p {
	case "error":
		return Red("●")
	case "timeout":
		return Red("⊗")
	case "warning", "maybe":
		return Yellow("●")
	case "debug", "low":
		return Cyan("●")
	case "info":
		return Magenta("●")
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

// remove duplicates from a slice of strings
func RemoveDuplicates(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	m := make(map[string]bool)
	for _, s := range slice {
		m[s] = true
	}
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}

func KebabToCamelCase(kebab string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}
