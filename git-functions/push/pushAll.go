package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/choose"
	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/input"
	"github.com/KevinYouu/fastGit/functions/log"
	"github.com/KevinYouu/fastGit/git-functions/auth"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func PushAll() {
	repoPath := "."

	suffix := choose.Choose([]string{"fix", "feat", "docs", "style", "refactor", "test", "chore", "revert"})
	data := input.Input("Enter your commit message: \n", "commit message", "\n(esc to quit)")

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

	// create the SSH authWithSSH
	authWithSSH := auth.AuthWithSSH(worktree)
	var authWithPassword *http.BasicAuth

	err = worktree.Pull(&git.PullOptions{Auth: authWithSSH})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		fmt.Println(colors.RenderColor("yellow", "SSH submission failed, Trying to push with username and password..."), err)

		authWithPassword = auth.AuthWithPassword(worktree)
		err = worktree.Pull(&git.PullOptions{Auth: authWithPassword})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Println(colors.RenderColor("red", "Failed to pull changes:"), err)
			log.CheckIfError(err)
			os.Exit(1)
		}
	}

	fmt.Println(colors.RenderColor("green", "Changes pulled successfully"))
	// get the remote
	remote, err := repo.Remote("origin")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get remote:"), err)
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

	fmt.Println(colors.RenderColor("green", "Changes added to index and committed successfully"))

	// push changes
	err = remote.Push(&git.PushOptions{
		RemoteName: remote.Config().Name,
		RemoteURL:  remote.Config().URLs[0],
		Auth:       authWithSSH,
		Progress:   os.Stdout,
	})

	if err != nil {
		fmt.Println(colors.RenderColor("yellow", "SSH submission failed: "), err)
		fmt.Println(colors.RenderColor("yellow", "Trying to push with username and password..."), err)

		err = remote.Push(&git.PushOptions{
			RemoteName: remote.Config().Name,
			RemoteURL:  remote.Config().URLs[0],
			Auth:       authWithPassword,
			Progress:   os.Stdout,
		})
		if err != nil {
			fmt.Println(colors.RenderColor("red", "Failed to push to remote repository:"), err)
			os.Exit(1)
		}
	}

	fmt.Println(colors.RenderColor("green", "Pushed changes successfully"))
}
