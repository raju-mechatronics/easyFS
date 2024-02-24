package gofs

import (
	"os"
	"path/filepath"
	"regexp"
)

func IsValidDirName(name string) bool {
	// Regular expression to match invalid characters
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)

	// Check if name ends with space or period
	if !(name == "." || name == "..") && (name[len(name)-1] == ' ' || name[len(name)-1] == '.') {
		return false
	}

	// Check if name contains invalid characters
	if invalidCharsRegex.MatchString(name) {
		return false
	}

	return true
}

func IsValidFileName(name string) bool {
	// Regular expression to match invalid characters
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)

	// Check if name contains invalid characters
	return !invalidCharsRegex.MatchString(name)

}

func Join(path string, names ...string) PathHandler {
	return PathHandler(filepath.Join(append([]string{path}, names...)...))
}

func IsWindows() bool {
	return os.PathSeparator == '\\'
}
func IsUnix() bool {
	return os.PathSeparator == '/'
}

func GetSeparator() rune {
	return os.PathSeparator
}
