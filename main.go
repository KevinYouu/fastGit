package main

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/version"

	"github.com/KevinYouu/fastGit/git-functions/clone"
	"github.com/KevinYouu/fastGit/git-functions/merge"
	"github.com/KevinYouu/fastGit/git-functions/push"
	"github.com/KevinYouu/fastGit/git-functions/remote"
	"github.com/KevinYouu/fastGit/git-functions/reset"
	"github.com/KevinYouu/fastGit/git-functions/status"
	"github.com/KevinYouu/fastGit/git-functions/tag"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		version.GetVersion()
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
	case "rs":
		reset.Reset()
	case "t":
		tag.CreateAndPushTag()
	case "s":
		status.Status()
	case "m":
		merge.MergeIntoCurrent()
	case "v":
		version.GetVersion()
	default:
		fmt.Println("unknown command:", args[1])
		os.Exit(1)
	}
}
