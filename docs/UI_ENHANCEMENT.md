# FastGit UI 增强说明

经过优化后，FastGit 的 TUI 组件现在具有更现代、更美观的界面和丰富的加载动画效果。

## 🎨 主要改进

### 1. 全新的视觉主题

- **现代配色方案**: 采用青蓝色渐变主题，提供更好的视觉体验
- **增强的对比度**: 更清晰的文字和背景对比
- **丰富的状态色彩**: 成功、错误、警告、信息等状态都有独特的颜色标识
- **圆角和阴影**: 现代化的 UI 元素设计

### 2. 优化的表单组件

#### 选择表单 (Select)

```go
// 带图标和中文描述的选择表单
label, value, err := form.SelectFormWithStringSlice("请选择功能级别", options)
```

- 🎯 装饰性图标标题
- 💡 友好的中文操作提示
- 更好的视觉层次和选中效果

#### 输入表单 (Input)

```go
// 带验证和占位符的输入表单
result, err := form.Input("请输入项目名称", "my-awesome-project")
```

- ✏️ 清晰的输入指示
- 输入验证和错误提示
- 默认值和占位符支持

#### 多选表单 (MultiSelect)

```go
// 多选复选框表单
values, err := form.MultiSelectForm("选择项目特性", options)
```

- ☑️ 直观的多选界面
- 清晰的选择状态指示
- 批量选择支持

#### 确认表单 (Confirm)

```go
// 美化的确认对话框
confirmed := form.Confirm("是否要继续执行操作？")
```

- ❓ 明确的确认提示
- 键盘和鼠标操作支持

### 3. 高级加载动画

#### 多种动画类型

```go
// 默认 spinner
spinner.SpinnerDefault

// 脉冲动画
spinner.SpinnerPulse

// 点状动画
spinner.SpinnerDots

// 箭头旋转动画
spinner.SpinnerArrow
```

#### 高级 Spinner 功能

```go
// 创建带进度和步骤的高级动画
advSpinner := spinner.NewAdvancedSpinner(spinner.SpinnerDefault, "执行中...")
advSpinner.SetProgress(0.5)  // 设置进度
advSpinner.SetSteps(steps)   // 设置步骤列表
```

#### 多步骤命令执行

```go
steps := []command.MultiStepInfo{
    {
        Name:        "初始化",
        Description: "初始化项目环境",
        Command:     "git",
        Args:        []string{"init"},
        LoadingMsg:  "正在初始化...",
    },
    // ... 更多步骤
}

err := command.RunMultiStepCommand(steps)
```

### 4. 新增的实用函数

#### 状态图标

```go
icon := theme.GetStatusIcon("success")  // ✅
icon := theme.GetStatusIcon("error")    // ❌
icon := theme.GetStatusIcon("warning")  // ⚠️
icon := theme.GetStatusIcon("info")     // ℹ️
```

#### 装饰性元素

```go
// 获取带装饰的标题
header := theme.GetHeader("FastGit 操作")

// 获取分隔线
rule := theme.GetHorizontalRule(50)
```

## 🚀 使用示例

### 基础 Git 操作

```go
// 使用高级 spinner 执行 Git 命令
_, err := command.RunCmdWithAdvancedSpinner(
    "git",
    []string{"status"},
    "检查 Git 状态...",
    "状态检查完成",
    spinner.SpinnerDefault,
)
```

### 多步骤 Git 工作流

```go
gitSteps := []command.MultiStepInfo{
    {
        Name: "检查状态",
        Command: "git",
        Args: []string{"status"},
        LoadingMsg: "检查仓库状态...",
    },
    {
        Name: "添加文件",
        Command: "git",
        Args: []string{"add", "."},
        LoadingMsg: "添加文件到暂存区...",
    },
    {
        Name: "提交更改",
        Command: "git",
        Args: []string{"commit", "-m", "Auto commit"},
        LoadingMsg: "提交更改...",
    },
}

err := command.RunMultiStepCommand(gitSteps)
```

## 🎪 演示程序

运行 `examples/ui_demo.go` 可以看到所有新功能的演示：

```bash
go run examples/ui_demo.go
```

演示内容包括：

- 各种表单组件的使用
- 不同类型的加载动画
- 多步骤命令执行
- 完整的 Git 操作流程

## 🎨 主题自定义

你可以通过修改 `internal/theme/theme.go` 中的颜色变量来自定义主题：

```go
// 自定义主色调
PrimaryColor = lipgloss.Color("#YOUR_COLOR")

// 自定义状态颜色
SuccessColor = lipgloss.Color("#YOUR_SUCCESS_COLOR")
ErrorColor = lipgloss.Color("#YOUR_ERROR_COLOR")
```

## ⚡ 性能优化

- 动画帧率优化，减少 CPU 占用
- 智能的进度计算和显示
- 异步命令执行，不阻塞 UI
- 内存使用优化

## 🔧 技术栈

- **UI 框架**: charmbracelet/huh + lipgloss
- **动画引擎**: charmbracelet/bubbles
- **终端控制**: charmbracelet/bubbletea
- **样式系统**: 自定义主题系统

现在你的 FastGit 拥有了更加现代和用户友好的界面！🎉
