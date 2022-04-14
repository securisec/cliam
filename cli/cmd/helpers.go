package cmd

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/securisec/cliam/logger"
)

// remove duplicates from a slice of strings
func removeDuplicates(slice []string) []string {
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
