package status

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func Status() {
	// Open the repository
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println("Failed to open repository:", err)
		os.Exit(1)
	}

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Println("Failed to get worktree:", err)
		os.Exit(1)
	}

	// Get the list of untracked files
	untracked, err := worktree.Filesystem.ReadDir(".")
	if err != nil {
		fmt.Println("Failed to get untracked files:", err)
		os.Exit(1)
	}
	fmt.Println("\nUntracked files:")
	for _, file := range untracked {
		// fmt.Println(file.Name())
		if file.Name() != ".git" {
			if file.IsDir() {
				fmt.Println(file.Name() + "/")
			} else {
				fmt.Println(file.Name())
			}
		}
	}
}
