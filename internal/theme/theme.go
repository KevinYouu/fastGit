package theme

import (
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// 定义主题颜色 - 采用现代渐变配色方案
var (
	// 主色调 - 使用更鲜艳的渐变色
	PrimaryColor   = lipgloss.Color("#00D9FF") // 青蓝色
	SecondaryColor = lipgloss.Color("#7C3AED") // 紫色
	AccentColor    = lipgloss.Color("#F59E0B") // 金黄色

	// 渐变辅助色
	GradientStart  = lipgloss.Color("#667EEA") // 渐变起始色
	GradientEnd    = lipgloss.Color("#764BA2") // 渐变结束色
	HighlightColor = lipgloss.Color("#FF6B6B") // 高亮色

	// 状态颜色 - 更鲜明的状态指示
	SuccessColor = lipgloss.Color("#00F5A0") // 亮绿色
	ErrorColor   = lipgloss.Color("#FF5E5B") // 亮红色
	WarningColor = lipgloss.Color("#FFD23F") // 亮黄色
	InfoColor    = lipgloss.Color("#4ECDC4") // 青绿色

	// 中性颜色 - 更好的对比度
	TextColor     = lipgloss.Color("#FFFFFF") // 纯白文字
	TextSecondary = lipgloss.Color("#E2E8F0") // 浅灰文字
	TextMuted     = lipgloss.Color("#A0AEC0") // 静音文字
	TextDark      = lipgloss.Color("#4A5568") // 深色文字
	BorderColor   = lipgloss.Color("#718096") // 边框色
	BorderActive  = lipgloss.Color("#00D9FF") // 激活边框色

	// 背景颜色 - 更深的层次感
	BackgroundColor     = lipgloss.Color("#0F172A") // 深蓝背景
	BackgroundLighter   = lipgloss.Color("#1E293B") // 浅一级背景
	BackgroundHighlight = lipgloss.Color("#334155") // 高亮背景
	BackgroundActive    = lipgloss.Color("#475569") // 激活背景
)

// 样式定义 - 增强视觉效果
var (
	// 基础样式 - 添加阴影和圆角
	BaseStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 0).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Background(BackgroundLighter)

	// 标题样式 - 增加渐变效果
	TitleStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Padding(1, 2).
			MarginBottom(1).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(PrimaryColor).
			Background(BackgroundHighlight).
			Align(lipgloss.Center)

	// 子标题样式
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Bold(true).
			Padding(0, 1).
			Italic(true)

	// 描述样式 - 更好的可读性
	DescriptionStyle = lipgloss.NewStyle().
				Foreground(TextSecondary).
				Italic(true).
				Padding(0, 1).
				MarginBottom(1)

	// 状态样式组 - 带图标和背景
	ErrorStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(ErrorColor).
			Bold(true).
			Padding(0, 2)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(SuccessColor).
			Bold(true).
			Padding(0, 2)

	WarningStyle = lipgloss.NewStyle().
			Foreground(TextDark).
			Background(WarningColor).
			Bold(true).
			Padding(0, 2)

	InfoStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(InfoColor).
			Bold(true).
			Padding(0, 2)

	// 输入框样式 - 增强交互反馈
	InputStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(BackgroundLighter).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Padding(1, 2).
			MarginBottom(1)

	// 焦点输入框样式 - 添加发光效果
	FocusedInputStyle = lipgloss.NewStyle().
				Foreground(TextColor).
				Background(BackgroundHighlight).
				BorderStyle(lipgloss.ThickBorder()).
				BorderForeground(PrimaryColor).
				Padding(1, 2).
				MarginBottom(1)

	// 选择器容器样式
	SelectStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(BackgroundLighter).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Padding(1, 2).
			MarginBottom(1)

	// 选中项样式 - 更显眼的选中效果
	SelectedStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(PrimaryColor).
			Bold(true).
			Padding(0, 2).
			MarginLeft(2)

	// 未选中项样式
	UnselectedStyle = lipgloss.NewStyle().
			Foreground(TextSecondary).
			Background(BackgroundLighter).
			Padding(0, 2).
			MarginLeft(2)

	// 悬停项样式
	HoverStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(BackgroundActive).
			Padding(0, 2).
			MarginLeft(2)

	// 加载动画样式 - 更炫酷的动画
	SpinnerStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Padding(0, 1)

	// 进度条样式 - 渐变进度条
	ProgressStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Background(BackgroundLighter).
			Padding(0, 1)

	// 卡片样式 - 用于包装内容
	CardStyle = lipgloss.NewStyle().
			Background(BackgroundLighter).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Padding(2, 3).
			Margin(1, 0)

	// 按钮样式
	ButtonStyle = lipgloss.NewStyle().
			Foreground(TextColor).
			Background(PrimaryColor).
			Bold(true).
			Padding(1, 3).
			MarginRight(2)

	// 次要按钮样式
	SecondaryButtonStyle = lipgloss.NewStyle().
				Foreground(TextColor).
				Background(SecondaryColor).
				Bold(true).
				Padding(1, 3).
				MarginRight(2)
)

// GetCustomTheme 返回优化的自定义huh主题
func GetCustomTheme() *huh.Theme {
	theme := huh.ThemeBase()

	// 设置表单基础样式 - 更现代的设计
	theme.Focused.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(PrimaryColor).
		Background(BackgroundLighter).
		Padding(2, 3).
		MarginBottom(1)

	theme.Blurred.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(BorderColor).
		Background(BackgroundColor).
		Padding(2, 3).
		MarginBottom(1)

	// 设置标题样式 - 更醒目的标题
	theme.Focused.Title = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Bold(true).
		Padding(0, 0, 1, 0).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(PrimaryColor).
		Width(50).
		Align(lipgloss.Center)

	theme.Blurred.Title = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 0, 1, 0).
		Width(50).
		Align(lipgloss.Center)

	// 设置描述样式 - 更好的可读性
	theme.Focused.Description = lipgloss.NewStyle().
		Foreground(TextSecondary).
		Italic(true).
		Padding(0, 0, 1, 0).
		Width(50).
		Align(lipgloss.Center)

	theme.Blurred.Description = lipgloss.NewStyle().
		Foreground(TextMuted).
		Italic(true).
		Padding(0, 0, 1, 0).
		Width(50).
		Align(lipgloss.Center)

	// 设置选择器样式 - 更好的视觉层次
	theme.Focused.SelectedOption = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(PrimaryColor).
		Bold(true).
		Padding(0, 2).
		MarginRight(1)

	theme.Focused.UnselectedOption = lipgloss.NewStyle().
		Foreground(TextSecondary).
		Background(BackgroundHighlight).
		Padding(0, 2).
		MarginRight(1)

	theme.Blurred.SelectedOption = lipgloss.NewStyle().
		Foreground(TextMuted).
		Background(BackgroundColor).
		Padding(0, 2).
		MarginRight(1)

	theme.Blurred.UnselectedOption = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 2).
		MarginRight(1)

	// 设置输入框样式 - 现代输入框设计
	theme.Focused.TextInput.Cursor = lipgloss.NewStyle().
		Foreground(AccentColor).
		Bold(true)

	theme.Focused.TextInput.Placeholder = lipgloss.NewStyle().
		Foreground(TextMuted).
		Italic(true)

	theme.Focused.TextInput.Prompt = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Bold(true).
		Background(BackgroundHighlight).
		Padding(0, 1)

	theme.Focused.TextInput.Text = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(BackgroundLighter).
		Padding(0, 1)

	// 设置错误样式
	theme.Focused.ErrorIndicator = lipgloss.NewStyle().
		Foreground(ErrorColor).
		Bold(true)

	theme.Focused.ErrorMessage = lipgloss.NewStyle().
		Foreground(ErrorColor).
		Italic(true).
		Padding(0, 1)

	return theme
}

// GetCompactTheme 返回适用于低高度终端的紧凑主题
func GetCompactTheme(isCompact bool) *huh.Theme {
	theme := huh.ThemeBase()

	if isCompact {
		// 紧凑模式：最小化边距和装饰，完全移除顶部空白
		theme.Focused.Base = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(PrimaryColor).
			Padding(0, 1, 1, 1). // 上, 右, 下, 左 - 顶部0边距，底部有1行内边距
			Margin(0, 0, 1, 0).  // 上, 右, 下, 左 - 顶部0边距，底部有1行外边距
			MarginTop(0)         // 明确设置顶部边距为0

		theme.Blurred.Base = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(BorderColor).
			Padding(0, 1, 1, 1). // 上, 右, 下, 左 - 顶部0边距，底部有1行内边距
			Margin(0, 0, 1, 0).  // 上, 右, 下, 左 - 顶部0边距，底部有1行外边距
			MarginTop(0)         // 明确设置顶部边距为0

		// 紧凑标题样式 - 完全无边距
		theme.Focused.Title = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Margin(0).    // 无边距
			MarginTop(0). // 明确设置顶部边距为0
			PaddingTop(0) // 明确设置顶部内边距为0

		theme.Blurred.Title = lipgloss.NewStyle().
			Foreground(TextMuted).
			Margin(0).    // 无边距
			MarginTop(0). // 明确设置顶部边距为0
			PaddingTop(0) // 明确设置顶部内边距为0

		// 紧凑描述样式
		theme.Focused.Description = lipgloss.NewStyle().
			Foreground(TextSecondary)

		theme.Blurred.Description = lipgloss.NewStyle().
			Foreground(TextMuted)
	} else {
		// 正常模式：使用标准主题
		theme.Focused.Base = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Background(BackgroundLighter).
			Padding(1, 2)

		theme.Blurred.Base = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Background(BackgroundColor).
			Padding(1, 2)

		// 标准标题样式
		theme.Focused.Title = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true).
			Padding(0, 0, 1, 0)

		theme.Blurred.Title = lipgloss.NewStyle().
			Foreground(TextMuted).
			Padding(0, 0, 1, 0)

		// 标准描述样式
		theme.Focused.Description = lipgloss.NewStyle().
			Foreground(TextSecondary).
			Italic(true).
			Padding(0, 0, 1, 0)

		theme.Blurred.Description = lipgloss.NewStyle().
			Foreground(TextMuted).
			Italic(true).
			Padding(0, 0, 1, 0)
	}

	// 选择器样式保持一致
	theme.Focused.SelectedOption = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(PrimaryColor).
		Bold(true).
		Padding(0, 1)

	theme.Focused.UnselectedOption = lipgloss.NewStyle().
		Foreground(TextSecondary).
		Padding(0, 1)

	theme.Blurred.SelectedOption = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 1)

	theme.Blurred.UnselectedOption = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 1)

	// 输入框样式
	theme.Focused.TextInput.Cursor = lipgloss.NewStyle().
		Foreground(AccentColor).
		Bold(true)

	theme.Focused.TextInput.Placeholder = lipgloss.NewStyle().
		Foreground(TextMuted).
		Italic(true)

	theme.Focused.TextInput.Prompt = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Bold(true)

	// 错误样式
	theme.Focused.ErrorIndicator = lipgloss.NewStyle().
		Foreground(ErrorColor).
		Bold(true)

	theme.Focused.ErrorMessage = lipgloss.NewStyle().
		Foreground(ErrorColor).
		Italic(true)

	// 隐藏帮助信息
	theme.Help.Ellipsis = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // 设置为透明或与背景同色
		Width(0).
		Height(0)

	theme.Help.ShortKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // 设置为透明
		Width(0).
		Height(0)

	theme.Help.ShortDesc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // 设置为透明
		Width(0).
		Height(0)

	theme.Help.ShortSeparator = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // 设置为透明
		Width(0).
		Height(0)

	theme.Help.FullKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // 设置为透明
		Width(0).
		Height(0)

	theme.Help.FullDesc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // 设置为透明
		Width(0).
		Height(0)

	theme.Help.FullSeparator = lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")). // 设置为透明
		Width(0).
		Height(0)

	return theme
}

// GetUltraCompactTheme 返回针对极低高度终端的超紧凑主题
// 将描述信息移到右侧，最大化纵向空间利用
func GetUltraCompactTheme() *huh.Theme {
	theme := huh.ThemeBase()

	// 超紧凑基础样式：无边框，完全移除顶部空白
	theme.Focused.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.HiddenBorder()).
		Padding(0, 0, 1, 0). // 顶部0边距，底部有1行内边距
		Margin(0, 0, 1, 0).  // 顶部0边距，底部有1行外边距
		MarginTop(0)         // 明确设置顶部边距为0

	theme.Blurred.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.HiddenBorder()).
		Padding(0, 0, 1, 0). // 顶部0边距，底部有1行内边距
		Margin(0, 0, 1, 0).  // 顶部0边距，底部有1行外边距
		MarginTop(0)         // 明确设置顶部边距为0

	// 超紧凑标题：单行，无装饰，完全无边距
	theme.Focused.Title = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Bold(true).
		Padding(0).
		Margin(0).
		MarginTop(0). // 明确设置顶部边距为0
		PaddingTop(0) // 明确设置顶部内边距为0

	theme.Blurred.Title = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0).
		Margin(0).
		MarginTop(0). // 明确设置顶部边距为0
		PaddingTop(0) // 明确设置顶部内边距为0

	// 隐藏描述文本，因为要移到右侧显示
	theme.Focused.Description = lipgloss.NewStyle().
		Foreground(lipgloss.Color("")).
		Height(0).
		Padding(0).
		Margin(0)

	theme.Blurred.Description = lipgloss.NewStyle().
		Foreground(lipgloss.Color("")).
		Height(0).
		Padding(0).
		Margin(0)

	// 选择器样式：紧凑但清晰
	theme.Focused.SelectedOption = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(PrimaryColor).
		Bold(true).
		Padding(0, 1)

	theme.Focused.UnselectedOption = lipgloss.NewStyle().
		Foreground(TextSecondary).
		Padding(0, 1)

	theme.Blurred.SelectedOption = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 1)

	theme.Blurred.UnselectedOption = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 1)

	// 确认按钮样式
	theme.Focused.FocusedButton = lipgloss.NewStyle().
		Foreground(TextColor).
		Background(PrimaryColor).
		Bold(true).
		Padding(0, 2)

	theme.Focused.BlurredButton = lipgloss.NewStyle().
		Foreground(TextSecondary).
		Padding(0, 2)

	return theme
}

// GetProgressBarStyle 获取现代化进度条样式
func GetProgressBarStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Background(BackgroundLighter).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor)
}

// GetSpinnerFrames 获取多样化的加载动画帧
func GetSpinnerFrames() []string {
	return []string{
		"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏",
	}
}

// GetPulseSpinnerFrames 获取脉冲动画帧
func GetPulseSpinnerFrames() []string {
	return []string{
		"●", "◐", "◑", "◒", "◓", "◔", "◕", "◖", "◗", "◘",
	}
}

// GetDotsSpinnerFrames 获取点状动画帧
func GetDotsSpinnerFrames() []string {
	return []string{
		"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷",
	}
}

// GetArrowSpinnerFrames 获取箭头旋转动画帧
func GetArrowSpinnerFrames() []string {
	return []string{
		"→", "↘", "↓", "↙", "←", "↖", "↑", "↗",
	}
}

// GetSpinnerStyle 获取加载动画样式
func GetSpinnerStyle() lipgloss.Style {
	return SpinnerStyle
}

// GetHorizontalRule 获取分隔线样式
func GetHorizontalRule(width int) string {
	rule := lipgloss.NewStyle().
		Foreground(BorderColor).
		Width(width).
		Align(lipgloss.Center).
		Render(strings.Repeat("─", width))
	return rule
}

// GetHeader 获取带装饰的标题
func GetHeader(title string) string {
	headerStyle := lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Background(BackgroundHighlight).
		Bold(true).
		Padding(1, 3).
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(PrimaryColor).
		Width(60).
		Align(lipgloss.Center)

	return headerStyle.Render(title)
}

// GetStatusIcon 根据状态获取图标
func GetStatusIcon(status string) string {
	icons := map[string]string{
		"success":  "✅",
		"error":    "❌",
		"warning":  "⚠️",
		"info":     "ℹ️",
		"loading":  "⏳",
		"pending":  "⏸️",
		"complete": "✨",
	}

	if icon, exists := icons[status]; exists {
		return icon
	}
	return "•"
}
