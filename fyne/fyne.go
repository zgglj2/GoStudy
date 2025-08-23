package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// 创建输入区域
func createInputArea(label *widget.Label) *fyne.Container {
	input := widget.NewEntry()
	input.SetPlaceHolder("请输入文本...")

	button := widget.NewButton("点击我", func() {
		label.SetText("你输入的是: " + input.Text)
	})

	return container.NewVBox(input, button)
}

// 创建选项区域
func createOptionsArea(label *widget.Label) *fyne.Container {
	check := widget.NewCheck("同意条款", func(value bool) {
		if value {
			label.SetText("已同意条款")
		} else {
			label.SetText("未同意条款")
		}
	})

	selectList := widget.NewSelect([]string{"选项1", "选项2", "选项3"}, func(value string) {
		label.SetText("你选择了: " + value)
	})

	return container.NewVBox(check, selectList)
}

// 创建进度区域
func createProgressArea() *fyne.Container {
	progress := widget.NewProgressBar()
	progress.SetValue(0.5)

	slider := widget.NewSlider(0, 1)
	slider.Step = 0.01
	slider.OnChanged = func(value float64) {
		progress.SetValue(value)
	}

	return container.NewVBox(progress, slider)
}

func main() {
	// 创建应用
	myApp := app.New()
	// 创建窗口
	myWindow := myApp.NewWindow("Fyne 示例程序")

	// 创建标签
	label := widget.NewLabel("欢迎使用 Fyne!")

	// 创建各个区域
	inputArea := createInputArea(label)
	optionsArea := createOptionsArea(label)
	progressArea := createProgressArea()

	// 使用垂直布局容器组织所有组件
	content := container.NewVBox(
		label,
		widget.NewSeparator(),
		inputArea,
		widget.NewSeparator(),
		optionsArea,
		widget.NewSeparator(),
		progressArea,
	)

	// 设置窗口内容
	myWindow.SetContent(content)

	// 设置窗口大小
	myWindow.Resize(fyne.NewSize(300, 400))

	// 显示并运行
	myWindow.ShowAndRun()
}
