package push

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KevinYouu/fastGit/functions/choose"
	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/input"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
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

	// get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user's home directory:", err)
		os.Exit(1)
	}

	// read the private key
	privateKeyPath := filepath.Join(homeDir, ".ssh", "id_rsa")
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println("Failed to read private key file:", err)
		os.Exit(1)
	}

	// create the SSH auth
	auth, err := ssh.NewPublicKeys("git", []byte(privateKey), "")
	if err != nil {
		fmt.Println("Failed to create SSH auth:", err)
		os.Exit(1)
	}
	// push changes
	err = remote.Push(&git.PushOptions{
		RemoteName: remote.Config().Name,
		RemoteURL:  remote.Config().URLs[0],
		Auth:       auth,
	})
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to push to remote repository:"), err)
		os.Exit(1)
	}

	fmt.Println(colors.RenderColor("green", "Pushed changes successfully"))
}
