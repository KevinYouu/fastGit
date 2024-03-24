package main

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/git-functions/push"
	"github.com/KevinYouu/fastGit/git-functions/remote"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		return
	}

	switch args[1] {
	case "pa":
		push.PushAll()
	case "ra":
		remoteName, remoteUrl, err := form.FormInput()
		if err != nil {
			fmt.Println("❌ line 38 err ➡️", err)
			os.Exit(1)
		}
		remote.Add(remoteName, remoteUrl)
	default:
		fmt.Println("unknown command:", args[1])
		os.Exit(1)
	}
}
