package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func AuthWithPassword(worktree *git.Worktree) *http.BasicAuth {
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

	return authWithPassword
}

func AuthWithSSH(worktree *git.Worktree) *ssh.PublicKeys {
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
	return auth
}
