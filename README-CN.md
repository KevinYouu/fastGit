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
curl -sSL https://raw.githubusercontent.com/KevinYouu/fastGit/main/install.sh | bash
```

或者

```bash
wget -qO- https://raw.githubusercontent.com/KevinYouu/fastGit/main/install.sh | bash
```

#### 3. 运行

```bash
# 提交工作区全部已更改的文件
fastGit pa
```

```bash
# 提交已选择的文件
fastGit ps
```

### 功能

- [x] `fastGit pa`, 提交工作区全部已更改的代码

- [x] `fastGit ps`, 提交部分已修改的代码

- [x] `fastGit t`, 创建和推送 tag

- [x] `fastGit m`, 将选择的分支合并到当前分支

- [x] `fastGit rs`, 重置到选择的 hash 版本

- [x] `fastGit init`, 初始化 fastGit 配置

- [x] `fastGit s`, 查看工作区状态

- [x] `fastGit rv`, 获取所有远程仓库

更多功能正在开发中.....

### 感谢以下开源项目

[go](https://github.com/golang/go)

[cobra](https://github.com/spf13/cobra)

[bubbletea](https://github.com/charmbracelet/bubbletea)

[huh](https://github.com/charmbracelet/huh)
