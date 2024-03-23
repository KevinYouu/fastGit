package remote

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/confirm"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/go-git/go-git/v5"
)

func GetRemotes() []*git.Remote {
	repoPath := "."

	// open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to open repository:"), err)
		os.Exit(1)
	}

	remotes, err := repo.Remotes()
	if err != nil {
		fmt.Println("Failed to get remotes:", err)
		os.Exit(1)
	}
	if len(remotes) == 0 {
		fmt.Println("No remotes found")
		data := confirm.Confirm("Add a remote?")
		if data {
			remoteName, remoteUrl, err := form.FormInput()
			if err != nil {
				fmt.Println("❌ line 38 err ➡️", err)
				os.Exit(1)
			}
			Add(remoteName, remoteUrl)
		} else {
			os.Exit(0)
		}
	}

	return remotes
}
