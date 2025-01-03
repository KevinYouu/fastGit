#!/bin/bash
# Created by KevinYouu on 2024-03-24 14:30:11
# Description: install fastGit
# Github: https://github.com/KevinYouu/fastGit

repo="KevinYouu/fastGit"
# Check if git is installed
if ! command -v git &>/dev/null; then
    echo "git is not installed. Please install git first."
    echo "https://git-scm.com/downloads"
    exit
fi

version=$(curl -s https://api.github.com/repos/"$repo"/releases/latest | grep 'tag_name' | cut -d\" -f4)
echo "latest version: $version"

package_name=""

function systemCheck() {
    system=$(uname -s)
    arch=$(uname -m)

    case "$system" in
    Darwin)
        if [ "$arch" == "x86_64" ]; then
            package_name="darwin_amd64"
        elif [ "$arch" == "arm64" ]; then
            package_name="darwin_arm64"
        else
            echo "Your system is not supported."
        fi
        ;;
    Linux)
        if [ "$arch" == "x86_64" ]; then
            package_name="linux_amd64"
        elif [ "$arch" == "aarch64" ]; then
            package_name="linux_arm64"
        else
            echo "Your system is not supported."
        fi
        ;;
    *)
        echo "Your system is not supported."
        ;;
    esac
}

systemCheck
url=https://github.com/$repo/releases/download/$version/fastGit_"$version"_"$package_name".zip
file=fastGit_"$version"_"$package_name".zip

echo "Downloading $url"

if wget "$url"; then
    echo "Download successful"
else
    echo "Download failed"
    exit 1
fi
echo "Extracting and installing $file"

if sudo unzip -o "$file" -d /usr/local/bin/; then
    echo "Extraction successful"
else
    echo "Extraction failed"
    exit 1
fi

# Set permissions
if sudo chmod +x "/usr/local/bin/fastGit"; then
    echo "Permissions set successfully"
else
    echo "Failed to set permissions"
    exit 1
fi

echo "Installation completed successfully"

rm "$file"

SHELL_TYPE=$(echo "$SHELL" | awk -F/ '{print $NF}')

case $SHELL_TYPE in
"bash")
    source "$HOME"/.bashrc
    ;;
"zsh")
    source "$HOME"/.zshrc
    ;;
"fish")
    source "$HOME/.config/fish/config.fish"
    ;;
*)
    echo "Your shell is not supported."
    ;;
esac

# Initialize config
if [ ! -d "/usr/local/bin/fastGit" ]; then
    /usr/local/bin/fastGit init
    echo "Config initialized successfully"
fi
