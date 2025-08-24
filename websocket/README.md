# WebSocket 示例

该目录包含使用 Gorilla WebSocket 库的 Go 语言示例。

## 文件结构

- [server/](server/) - WebSocket 服务端示例，包含一个增强的回显服务器实现
- [client/](client/) - WebSocket 客户端示例，连接到服务器并支持交互式消息发送

## 功能特性

### 服务端特性
1. 连接统计（当前连接数、消息总数）
2. 详细的调试日志输出
3. 特殊命令处理（time、stats、ping）
4. 网页界面用于浏览器测试
5. 状态接口（/status）返回JSON格式的统计信息

### 客户端特性
1. 交互式命令行界面
2. 支持手动输入消息发送
3. 自动发送测试消息功能
4. 特殊命令支持（ping、time、stats）
5. 详细的调试信息输出

## 运行示例

1. 启动服务器：
   ```bash
   cd server
   go run main.go
   ```

2. 在另一个终端中，运行客户端：
   ```bash
   cd client
   go run main.go
   ```

3. 或者打开浏览器并导航到 http://localhost:9090 使用网页界面

## 工作原理

服务器实现了一个增强的回显服务：

1. 将 HTTP 连接升级为 WebSocket 连接
2. 从客户端读取消息
3. 根据消息内容进行特殊处理或回显
4. 提供统计信息和状态监控

客户端：

1. 连接到 WebSocket 服务器
2. 支持交互式命令输入
3. 可以手动发送消息或自动发送测试消息
4. 在中断时优雅地关闭连接

网页界面提供了一种通过浏览器测试 WebSocket 连接的简单方法。

## 特殊命令

客户端和服务端都支持以下特殊命令：

- `time` - 返回服务器当前时间
- `stats` - 返回服务器统计信息
- `ping` - 发送ping消息测试连接

## 端口配置

默认情况下，服务器运行在 `localhost:9090`。可以通过 `-addr` 参数修改：

```bash
go run server/main.go -addr localhost:8080
go run client/main.go -addr localhost:8080
```