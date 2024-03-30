package main

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/git-functions/push"
	"github.com/KevinYouu/fastGit/git-functions/remote"
	"github.com/KevinYouu/fastGit/git-functions/status"
	"github.com/KevinYouu/fastGit/git-functions/tag"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		return
	}

	switch args[1] {
	case "pa":
		push.PushAll()
	case "ps":
		push.PushSelected()
	case "ra":
		remote.Add()
	case "rv":
		remote.GetRemotes()
	case "t":
		tag.IncrementTagVersion()
	case "s":
		status.Status()
	case "a":
		log, err := command.RunCommand("git", "status")
		if err != nil {
			fmt.Println("error executing git status command:", err)
			os.Exit(1)
		}
		fmt.Println(log)
	default:
		fmt.Println("unknown command:", args[1])
		os.Exit(1)
	}
}
