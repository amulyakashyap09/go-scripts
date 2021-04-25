package dir

import (
	"os"
	"path/filepath"
)

// GetCurrentExecutable returns the running process file
func GetCurrentExecutable() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return ex
}

// GetWorkingDirectory returns the working dir path
func GetWorkingDirectory() string {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

// GetParentDirectory returns the working dir path
func GetParentDirectory() string {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	return parent
}
