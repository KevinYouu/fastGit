package remote

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/form"
)

func Add() {
	form_props := form.FormProps{
		Message:      "Enter the following information:",
		Field:        "remote name",
		Field2:       "remote url",
		FieldLength:  11,
		Field2Length: 11,
	}
	remoteName, remoteUrl, err := form.FormInput(form_props)
	if err != nil {
		fmt.Println("❌ line 38 err ➡️", err)
		os.Exit(1)
	}

	cmd := exec.Command("git", "remote", "add", remoteName, remoteUrl)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to add remote: "+err.Error()))
		return
	}
	fmt.Println(colors.RenderColor("green", "Remote added successfully."))
}
