package update

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// GitHub repo info
const (
	repoOwner = "KevinYouu"
	repoName  = "fastGit"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

// Get the latest release tag from GitHub API
func getLatestVersion() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}
	return release.TagName, nil
}

// Get platform-specific release name
func getPlatformName() (string, error) {
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return "windows_amd64", nil
		case "arm64":
			return "windows_arm64", nil
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return "linux_amd64", nil
		case "arm64":
			return "linux_arm64", nil
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			return "darwin_amd64", nil
		case "arm64":
			return "darwin_arm64", nil
		}
	}
	return "", fmt.Errorf("unsupported platform: %s/%s", runtime.GOOS, runtime.GOARCH)
}

// Self-update function
func UpdateSelf() error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", "-Command", "iwr -useb https://raw.githubusercontent.com/KevinYouu/fastGit/main/install.ps1 | iex")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to run install script: %w", err)
		}
		fmt.Println("Update script executed. Please restart fastGit manually.")
		os.Exit(0)
	}

	version, err := getLatestVersion()
	if err != nil {
		return fmt.Errorf("failed to get latest version: %w", err)
	}
	fmt.Println("Latest version:", version)

	platform, err := getPlatformName()
	if err != nil {
		return err
	}

	assetName := fmt.Sprintf("fastGit_%s_%s.zip", version, platform)
	url := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", repoOwner, repoName, version, assetName)

	fmt.Println("Downloading:", url)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download update: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return err
	}

	var newBinary []byte
	for _, f := range zipReader.File {
		if f.Name == "fastGit" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()
			newBinary, err = io.ReadAll(rc)
			if err != nil {
				return err
			}
			break
		}
	}

	if newBinary == nil {
		return fmt.Errorf("fastGit binary not found in zip")
	}

	// Find actual install path
	execPath, err := os.Executable()
	if err != nil {
		return err
	}
	resolvedPath, err := filepath.EvalSymlinks(execPath)
	if err != nil {
		return err
	}

	// Detect install dir: /opt/homebrew/bin or /usr/local/bin
	targetPath := resolvedPath
	if runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" {
		targetPath = "/opt/homebrew/bin/fastGit"
	} else if strings.HasPrefix(resolvedPath, "/usr/local/bin/") {
		targetPath = resolvedPath
	}

	tmpPath := targetPath + ".tmp"
	if err := os.WriteFile(tmpPath, newBinary, 0755); err != nil {
		return fmt.Errorf("write failed: %w (hint: try sudo)", err)
	}

	if err := os.Rename(tmpPath, targetPath); err != nil {
		return fmt.Errorf("rename failed: %w (hint: try sudo)", err)
	}

	fmt.Println("Update completed successfully.")
	return nil
}
