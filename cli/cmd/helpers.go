package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/go-homedir"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
)

func expandPath(p string) string {
	h, err := homedir.Expand(p)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("failed to expand path")
	}
	ps, err := filepath.Abs(h)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("failed to get absolute path")
	}
	return ps
}

func printValidArgs(f func() []string) {
	fmt.Println(shared.Red("Valid arguments:"))
	for _, v := range f() {
		fmt.Println(v)
	}
}

func getRequest(url, service string) (int, error) {
	logger.LogDebugVerbose("url", url)
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	status := res.StatusCode
	if status == 200 {
		logger.LoggerStdErr.Info().Str("success", service).Str("url", url).Send()
		return status, nil
	} else {
		logger.LogDebug("failed", status)
	}
	return status, fmt.Errorf("bad status: %d", status)
}

func templateBuilder(t string, args map[string]string) (string, error) {
	temp := template.Must(template.New("").Parse(t))
	var tpl bytes.Buffer
	err := temp.Execute(&tpl, args)
	return tpl.String(), err
}

func ValidateJwtExpiration(token string) (isValid bool) {
	j, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		logger.LoggerStdErr.Err(err).Err(err).Msg("Failed to parse JWT")
	}
	// if exp field is not set, we want to return true
	return j.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), false)
}
