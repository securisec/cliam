package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

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
		logger.LogSuccess("success", service)
		return status, nil
	} else {
		logger.LogDebug("failed", status)
	}
	return status, fmt.Errorf("Bad status: %d", status)
}

func templateBuilder(t string, args map[string]string) (string, error) {
	temp := template.Must(template.New("").Parse(t))
	var tpl bytes.Buffer
	err := temp.Execute(&tpl, args)
	return tpl.String(), err
}
