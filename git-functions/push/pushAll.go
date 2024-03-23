package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/choose"
	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/input"

	"github.com/go-git/go-git/v5"
)

func PushAll() {
	repoPath := "."

	suffix := choose.Choose()
	data := input.Input("Enter your commit message: ", "commit message", "(esc to quit)")

	// open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to open repository:"), err)
		os.Exit(1)
	}

	// get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get worktree:"), err)
		os.Exit(1)
	}

	// add all files to the index
	_, err = worktree.Add(".")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to add file to index:"), err)
		os.Exit(1)
	}

	// commit changes
	_, err = worktree.Commit(suffix+": "+data, &git.CommitOptions{})

	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to commit changes:"), err)
		os.Exit(1)
	}

	fmt.Println(colors.RenderColor("blue", "Changes added to index and committed successfully"))

	// get the remote
	remote, err := repo.Remote("origin")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get remote:"), err)
		os.Exit(1)
	}

	// push changes
	err = remote.Push(&git.PushOptions{
		RemoteName: remote.Config().Name,
		RemoteURL:  remote.Config().URLs[0],
	})
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to push to remote repository:"), err)
		os.Exit(1)
	}

	fmt.Println(colors.RenderColor("green", "Pushed changes successfully"))
}
