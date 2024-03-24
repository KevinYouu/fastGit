English | [简体中文](README-CN.md)

fastGit is a tool that helps you quickly submit code with a command line interface. It supports Linux, Mac, and Windows. The inspiration comes from [gum](https://github.com/charmbracelet/gum)

> This project is utilizing its own features to submit code.

![](assets/fast-git.gif)

### How to use

#### Install fastGit

```bash
# Linux/macOS
curl -sSL https://github.com/KevinYouu/fastGit/install.sh | bash
```

or

```bash
wget -qO- https://github.com/KevinYouu/fastGit/install.sh | bash
```

#### Run

```bash
fastGit pa
```

### Features

- [x] `fastGit pa`, submit all changes in the working directory
- [x] `fastGit ra`, add a remote repository
- [ ] `fastGit ps`, submit some changes in the working directory
- [ ] `fastGit rb`, delete a remote repository

### Thanks to the following open source projects

[go](https://github.com/golang/go)

[go-git](https://github.com/go-git/go-git)

[bubbletea](github.com/charmbracelet/bubbletea)
