package push

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KevinYouu/fastGit/functions/choose"
	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/functions/input"
	"github.com/KevinYouu/fastGit/functions/log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func PushAll() {
	repoPath := "."

	suffix := choose.Choose()
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

	err = worktree.Pull(&git.PullOptions{Auth: auth})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		fmt.Println(colors.RenderColor("yellow", "SSH submission failed, Trying to push with username and password..."), err)

		formData := form.FormProps{
			Message:      "Enter the following information:",
			Field:        "username",
			Field2:       "password",
			FieldLength:  8,
			Field2Length: 8,
		}
		username, password, err := form.FormInput(formData)
		if err != nil {
			fmt.Println(colors.RenderColor("red", "Failed to get username and password: "), err)
			os.Exit(1)
		}
		if username == "" || password == "" {
			fmt.Println("GIT_USERNAME and/or GIT_PASSWORD environment variables are not set")
			os.Exit(1)
		}
		// create the basic auth
		authWithPassword := &http.BasicAuth{
			Username: username,
			Password: password,
		}
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
		Auth:       auth,
		Progress:   os.Stdout,
	})

	if err != nil {
		fmt.Println(colors.RenderColor("yellow", "SSH submission failed: "), err)
		fmt.Println(colors.RenderColor("yellow", "Trying to push with username and password..."), err)
		formData := form.FormProps{
			Message:      "Enter the following information:",
			Field:        "username",
			Field2:       "password",
			FieldLength:  8,
			Field2Length: 8,
		}
		username, password, err := form.FormInput(formData)
		if err != nil {
			fmt.Println(colors.RenderColor("red", "Failed to get username and password: "), err)
			os.Exit(1)
		}
		if username == "" || password == "" {
			fmt.Println("GIT_USERNAME and/or GIT_PASSWORD environment variables are not set")
			os.Exit(1)
		}
		// create the basic auth
		authWithPassword := &http.BasicAuth{
			Username: username,
			Password: password,
		}

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
