package clone

import (
	"fmt"
	"github.com/KevinYouu/fastGit/functions/input"
	"os"

	"github.com/go-git/go-git/v5"
)

func Clone() {
	// repositoryUrl := input.Input("", "", "(esc to quit)")
	repositoryUrl := input.Input("Enter Repository URL", "Enter URL", "(esc to quit)")
	// Clone the given repository to the given directory
	Info(repositoryUrl)

	_, err := git.PlainClone(".", false, &git.CloneOptions{
		URL:      repositoryUrl,
		Progress: os.Stdout,
	})

	CheckIfError(err)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
