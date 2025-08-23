# Go WebAssembly (WASM) 原理、编译和使用

## 原理

WebAssembly (WASM) 是一种低级类汇编语言，具有紧凑的二进制格式，可以在现代网络浏览器中以接近原生的性能运行。Go 从版本 1.11 开始支持编译为 WASM 格式，使得 Go 程序能够直接在浏览器中运行。

### WASM 工作原理

1. **编译**: Go 代码被编译成 WASM 字节码，这是一种可在浏览器中运行的低级虚拟机格式。
2. **运行时**: 浏览器内置 WASM 虚拟机，可以直接执行 WASM 字节码。
3. **交互**: WASM 模块可以通过 JavaScript 与 DOM 和浏览器 API 进行交互。
4. **性能**: WASM 提供接近原生的执行性能，比 JavaScript 更快。

### Go 的 WASM 支持

Go 的 WASM 支持基于 `js` GOOS 和 `wasm` GOARCH 架构。编译后的程序通过 `wasm_exec.js` 文件与 JavaScript 环境交互。

## 编译

### 环境设置

在编译 Go 程序为 WASM 之前，需要设置以下环境变量：

```bash
GOOS=js
GOARCH=wasm
CGO_ENABLED=0  # 禁用 CGO 以避免平台特定的依赖
```

### 编译命令

使用以下命令将 Go 程序编译为 WASM：

```bash
# 在 Linux/macOS 上
GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -o main.wasm .

# 在 Windows 上 (PowerShell)
$env:GOOS="js"; $env:GOARCH="wasm"; $env:CGO_ENABLED="0"; go build -o main.wasm .
```

> 注意：虽然在示例中 Linux/macOS 的命令没有显式设置 CGO_ENABLED=0，但为了确保跨平台兼容性，建议显式设置。在某些系统上，如果未禁用 CGO，可能会导致编译错误，正如我们在 Windows 环境中遇到的问题。

### 必需文件

编译完成后，需要以下文件来运行 WASM 程序：

1. `main.wasm` - 编译后的 WASM 二进制文件
2. `wasm_exec.js` - Go 提供的 JavaScript 支持文件，位于 Go 安装目录的 `misc/wasm/wasm_exec.js`

## 使用

### HTML 设置

创建一个 HTML 文件来加载和运行 WASM 程序：

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Go WASM Example</title>
</head>
<body>
    <div id="app"></div>
    
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</body>
</html>
```

### 运行

由于浏览器的同源策略和安全限制，WASM 文件必须通过 HTTP 服务器提供，不能直接通过文件协议打开。

可以使用 Python 启动一个简单的 HTTP 服务器：

```bash
# Python 3
# 在包含 main.wasm、wasm_exec.js 和 index.html 的目录中运行
python -m http.server 8000

# 然后在浏览器中访问 http://localhost:8000
# 如果目录中有 index.html 文件，它将作为默认页面加载
```

当运行 `python -m http.server 8000` 命令时，Python 会在当前目录启动一个 HTTP 服务器，默认情况下：
1. 服务器会提供当前目录及其子目录中的所有文件
2. 如果访问目录 URL（如 http://localhost:8000/），且该目录包含 index.html 文件，则会自动加载该文件
3. 浏览器会加载 index.html，其中引用了 wasm_exec.js，并通过 JavaScript 加载 main.wasm

### 其他服务器选项

除了 Python，还可以使用其他方式启动本地服务器：

1. Node.js (如果已安装):
   ```bash
   npx serve
   ```

2. Go 本身:
   ```bash
   go run -m http.server 8000
   ```

## 常见问题及解决方案

### 中文显示乱码问题

在 Fyne WASM 应用中，中文显示乱码是一个常见问题。这是由于 WASM 环境下缺少合适的中文字体导致的。

**解决方案：**

1. 在编译前设置环境变量指定中文字体：
   ```bash
   # Linux/macOS
   export FYNE_FONT="/path/to/chinese-font.ttf"
   
   # Windows (PowerShell)
   $env:FYNE_FONT="C:\Windows\Fonts\simhei.ttf"
   ```

2. 在代码中动态设置字体（推荐）：
   ```go
   package main
   
   import (
       "os"
       "strings"
       
       "github.com/flopp/go-findfont"
   )
   
   func init() {
       // 查找系统中的中文字体
       fontPaths := findfont.List()
       for _, path := range fontPaths {
           if strings.Contains(path, "simhei.ttf") ||  // 黑体
              strings.Contains(path, "simsun.ttc") ||  // 宋体
              strings.Contains(path, "simkai.ttf") {   // 楷体
               os.Setenv("FYNE_FONT", path)
               break
           }
       }
   }
   
   func main() {
       // 你的 Fyne 应用代码
   }
   ```

3. 注意：对于某些组件（如 TextGrid），可能需要设置专门的等宽字体环境变量：
   ```bash
   export FYNE_FONT_MONOSPACE="/path/to/chinese-monospace-font.ttf"
   ```

### 点击位置不正确问题

在某些情况下，Fyne WASM 应用中的点击位置可能与实际显示位置不匹配，这通常是由以下原因导致的：

1. **高DPI显示器支持问题**：浏览器缩放或高DPI显示器可能导致坐标计算错误
2. **Canvas大小计算问题**：WASM环境中Canvas大小与实际显示大小不一致

**解决方案：**

1. 确保HTML页面中正确设置viewport：
   ```html
   <head>
       <meta charset="utf-8">
       <meta name="viewport" content="width=device-width, initial-scale=1.0">
       <title>Go WASM Example</title>
   </head>
   ```

2. 尝试在CSS中设置canvas属性：
   ```css
   canvas {
       display: block;
       width: 100vw;
       height: 100vh;
       touch-action: none;
   }
   ```

3. 更新到最新版本的Fyne，新版本修复了部分WASM相关的问题：
   ```bash
   go get -u fyne.io/fyne/v2
   ```

4. 如果问题依然存在，可以尝试在代码中手动调整缩放比例：
   ```go
   // 获取当前窗口并调整DPI设置
   myWindow := myApp.NewWindow("Fyne Test")
   // 根据实际情况调整窗口大小
   myWindow.Resize(fyne.NewSize(300/1.5, 200/1.5)) // 除以缩放因子
   ```

## 示例说明

本项目包含两个版本的 Fyne 应用程序：

1. [main.go](file:///d:/workspace/go-workspace/src/GoStudy/fyne-wasm/main.go) - 非 WASM 版本，使用构建标签 `//go:build !wasm`
2. [main_wasm.go](file:///d:/workspace/go-workspace/src/GoStudy/fyne-wasm/main_wasm.go) - WASM 版本，使用构建标签 `//go:build wasm`

这种结构允许同一代码库在不同平台上构建不同的版本。

## 注意事项

1. WASM 程序无法直接访问操作系统功能，如文件系统等
2. 所有与外部环境的交互都需要通过 JavaScript 进行
3. 由于浏览器安全限制，WASM 文件必须通过 HTTP 服务器提供
4. 某些 Go 标准库功能在 WASM 环境中可能不可用或行为不同
5. 中文显示和点击位置问题是 Fyne WASM 的已知问题，需要特殊处理