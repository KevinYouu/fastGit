[English](README.md) | 简体中文

fastGit 是一个帮助你快速提交代码的命令行工具,支持 Linux、MacOS、Windows。灵感来自 [gum](https://github.com/charmbracelet/gum)

> 这个项目正在使用它自己来提交代码

> ![](assets/fast-git.gif)

### 如何使用

#### 1. 安装 Git

> 项目依赖于 Git, 请先安装 Git

```bash
# Debian/Ubuntu
sudo apt install git
# macOS
brew install git
```

#### 2. 安装 fastGit

```bash
# Linux/macOS
curl -sSL https://github.com/KevinYouu/fastGit/install.sh | bash
```

或者

```bash
wget -qO- https://github.com/KevinYouu/fastGit/install.sh | bash
```

#### 3. 运行

```bash
fastGit pa
```

### 功能

- [x] `fastGit pa`, 提交工作区全部已更改的代码

- [x] `fastGit ps`, 提交部分已修改的代码

- [x] `fastGit c`, 用于克隆远程仓库

- [x] `fastGit t`, 用于创建和推送 tag

- [x] `fastGit s`, 用于查看工作区状态

- [x] `fastGit ra`，用于添加远程仓库

更多功能正在开发中.....

### 感谢以下开源项目

[go](https://github.com/golang/go)

[go-git](https://github.com/go-git/go-git)

[bubbletea](github.com/charmbracelet/bubbletea)
