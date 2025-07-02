package main

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/spinner"
	"github.com/KevinYouu/fastGit/internal/theme"
)

func main() {
	// 演示新的 UI 组件和加载动画

	// 显示标题
	fmt.Println(theme.GetHeader("FastGit UI 演示"))
	fmt.Println()

	// 演示选择表单
	demoSelectForm()

	// 演示输入表单
	demoInputForm()

	// 演示多选表单
	demoMultiSelectForm()

	// 演示确认表单
	demoConfirmForm()

	// 演示各种加载动画
	demoSpinners()

	// 演示多步骤命令
	demoMultiStepCommand()
}

func demoSelectForm() {
	fmt.Println(theme.SubtitleStyle.Render("🎯 选择表单演示"))

	options := []string{
		"选项 1 - 基础功能",
		"选项 2 - 高级功能",
		"选项 3 - 专业功能",
		"选项 4 - 企业功能",
	}

	label, value, err := form.SelectFormWithStringSlice("请选择功能级别", options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("✅ 你选择了: %s (值: %s)\n\n", label, value)
}

func demoInputForm() {
	fmt.Println(theme.SubtitleStyle.Render("✏️ 输入表单演示"))

	result, err := form.Input("请输入项目名称", "my-awesome-project")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("✅ 项目名称: %s\n\n", result)
}

func demoMultiSelectForm() {
	fmt.Println(theme.SubtitleStyle.Render("☑️ 多选表单演示"))

	options := []string{
		"Docker 支持",
		"CI/CD 集成",
		"测试框架",
		"文档生成",
		"代码检查",
		"性能监控",
	}

	values, err := form.MultiSelectForm("选择项目特性", options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("✅ 选择的特性: %v\n\n", values)
}

func demoConfirmForm() {
	fmt.Println(theme.SubtitleStyle.Render("❓ 确认表单演示"))

	confirmed := form.Confirm("是否要继续执行操作？")
	if confirmed {
		fmt.Println("✅ 用户确认继续")
	} else {
		fmt.Println("❌ 用户取消操作")
	}
}

func demoSpinners() {
	fmt.Println(theme.SubtitleStyle.Render("🎪 加载动画演示"))

	// 演示默认 spinner
	fmt.Println("🔄 默认 Spinner")
	command.RunCmdWithCustomSpinner("sleep", []string{"2"}, "执行默认动画...", theme.GetSpinnerFrames())

	// 演示脉冲 spinner
	fmt.Println("\n💗 脉冲 Spinner")
	command.RunCmdWithCustomSpinner("sleep", []string{"2"}, "执行脉冲动画...", theme.GetPulseSpinnerFrames())

	// 演示点状 spinner
	fmt.Println("\n⚫ 点状 Spinner")
	command.RunCmdWithCustomSpinner("sleep", []string{"2"}, "执行点状动画...", theme.GetDotsSpinnerFrames())

	// 演示箭头 spinner
	fmt.Println("\n🔄 箭头 Spinner")
	command.RunCmdWithCustomSpinner("sleep", []string{"2"}, "执行箭头动画...", theme.GetArrowSpinnerFrames())

	fmt.Println()
}

func demoMultiStepCommand() {
	fmt.Println(theme.SubtitleStyle.Render("📋 多步骤命令演示"))

	steps := []command.MultiStepInfo{
		{
			Name:        "初始化",
			Description: "初始化项目环境",
			Command:     "sleep",
			Args:        []string{"1"},
			LoadingMsg:  "正在初始化项目...",
		},
		{
			Name:        "安装依赖",
			Description: "安装项目依赖包",
			Command:     "sleep",
			Args:        []string{"2"},
			LoadingMsg:  "正在安装依赖...",
		},
		{
			Name:        "构建项目",
			Description: "编译和构建项目",
			Command:     "sleep",
			Args:        []string{"1"},
			LoadingMsg:  "正在构建项目...",
		},
		{
			Name:        "运行测试",
			Description: "执行单元测试",
			Command:     "sleep",
			Args:        []string{"1"},
			LoadingMsg:  "正在运行测试...",
		},
		{
			Name:        "部署应用",
			Description: "部署到目标环境",
			Command:     "sleep",
			Args:        []string{"1"},
			LoadingMsg:  "正在部署应用...",
		},
	}

	err := command.RunMultiStepCommand(steps)
	if err != nil {
		fmt.Printf("❌ 多步骤命令执行失败: %v\n", err)
	} else {
		fmt.Println("🎉 多步骤命令执行成功！")
	}
}

// 演示简单的使用案例
func simpleDemo() {
	// 简单的 Git 操作演示
	fmt.Println(theme.GetHeader("Git 操作演示"))

	// 获取 Git 状态
	_, err := command.RunCmdWithAdvancedSpinner(
		"git",
		[]string{"status", "--porcelain"},
		"检查 Git 状态...",
		"Git 状态检查完成",
		spinner.SpinnerDefault,
	)

	if err != nil {
		fmt.Printf("Git 状态检查失败: %v\n", err)
		return
	}

	// 确认是否要提交
	if form.Confirm("是否要提交当前更改？") {
		// 多步骤 Git 操作
		gitSteps := []command.MultiStepInfo{
			{
				Name:        "添加文件",
				Description: "添加所有更改的文件",
				Command:     "git",
				Args:        []string{"add", "."},
				LoadingMsg:  "正在添加文件到暂存区...",
			},
			{
				Name:        "提交更改",
				Description: "提交更改到本地仓库",
				Command:     "git",
				Args:        []string{"commit", "-m", "Auto commit via FastGit"},
				LoadingMsg:  "正在提交更改...",
			},
		}

		err := command.RunMultiStepCommand(gitSteps)
		if err != nil {
			fmt.Printf("Git 操作失败: %v\n", err)
		} else {
			fmt.Println("🎉 Git 操作完成！")
		}
	}
}
