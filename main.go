package main

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/git-functions/clone"
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
	case "c":
		clone.Clone()
	case "pa":
		push.PushAll()
	case "ps":
		push.PushSelected()
	case "ra":
		remote.Add()
	case "rv":
		remote.GetRemotes()
	case "t":
		tag.CreateAndPushTag()
	case "s":
		status.Status()
	case "v":
		version.GetVersion()
	default:
		fmt.Println("unknown command:", args[1])
		os.Exit(1)
	}
}
