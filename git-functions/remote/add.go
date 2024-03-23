package remote

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func Add(remoteName, remoteUrl string) {
	repoPath := "."

	// open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Println("Failed to open repository:", err)
		os.Exit(1)
	}

	// create the remote
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{remoteUrl},
	})
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to add remote repository:"), err)
		os.Exit(1)
	}

	fmt.Println("Remote repository " + colors.RenderColor("green", remoteName+" ") + colors.RenderColor("green", remoteUrl) + " added successfully")
}
