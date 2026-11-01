# 第01章：Gin 框架概述与安装

## 本章目标
学完本章后，你将能够：
1. 理解什么是 Web 框架以及为什么需要它
2. 对比 Go 主流 Web 框架的差异
3. 安装 Gin 并跑通第一个 HTTP 服务
4. 理解 `gin.Default()` 和 `gin.New()` 的区别
5. 逐行理解 Gin 程序的每个部分

## 前置知识
- 需要先学习：Go 语言基础（变量、函数、结构体、指针、错误处理、`go mod`）
- 需要了解：用浏览器访问过网站（知道 URL 是什么即可）
- 不要求先学 MySQL 或 GORM；前 1–5 章先使用内存数据理解 HTTP 和路由

---

## 1.1 什么是 Web 框架

### 没有框架的世界：用 `net/http` 写服务

Go 标准库自带了 `net/http`，可以不用任何框架就写 HTTP 服务：

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册路由
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// 启动服务并处理启动/运行错误
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("服务退出:", err)
	}
}
```

这已经能跑了，但随着业务增长，你会发现很多重复工作：
- 需要自己组合路由和解析路径参数
- 获取 JSON Body 通常要自己调用 `json.Decoder` 或 `json.Unmarshal`
- 参数校验需要额外实现或引入库
- 日志、恢复、认证等横切逻辑需要自己组织
- 标准库本身不提供 Gin 风格的中间件链

### 框架帮你做了什么

```go
// 用 Gin 写同样的功能
r := gin.Default()
r.GET("/users/:id", func(c *gin.Context) {
	id := c.Param("id")         // 直接获取路径参数
	c.JSON(http.StatusOK, gin.H{"id": id}) // 一行返回 JSON
})
if err := r.Run(":8080"); err != nil {
	panic(err)
}
```

Web 框架帮你处理了：
- **路由匹配**：自动解析路径参数
- **参数解析**：自动绑定 JSON/Form/Query
- **参数校验**：一行标签搞定
- **响应封装**：一行返回 JSON/XML
- **中间件**：日志、恢复、认证、跨域...

---

## 1.2 Go Web 框架对比

| 框架 | 特点 | 适用场景 |
|------|------|---------|
| **Gin** | 轻量、路由和中间件 API 简洁、生态成熟 | 常见 REST API 和 Web 服务 |
| Echo | API 完整、组件较多 | 喜欢 Echo API 风格的项目 |
| Fiber | 接口风格接近 Express | 从 Node.js 迁移的团队 |
| Beego | 更偏全栈 MVC | 需要较完整框架约定的项目 |
| net/http | 标准库、零额外框架依赖 | 小型服务或需要完全自定义的项目 |

> Stars 会随时间变化，不能作为稳定的技术结论。选择框架时还应比较团队熟悉度、依赖、性能测试结果和维护状态。

**选择 Gin 的常见理由**：
- 路由、中间件和响应 API 较易上手
- `gin-contrib` 等生态可以覆盖常见需求
- 在 Go Web 项目中使用广泛，资料和示例较多
- 性能是否满足需求，应以自己的接口和部署环境压测为准

---

## 1.3 安装 Gin

### 第一步：创建项目

```bash
mkdir gin-tutorial
cd gin-tutorial
go mod init gin-tutorial
```

### 第二步：安装 Gin

```bash
go get github.com/gin-gonic/gin
```

`go get -u` 会尝试升级依赖，教学项目不必每次安装都强制升级；需要升级时再单独执行 `go get -u` 或通过版本管理更新。

安装完成后 `go.mod` 会记录实际解析到的版本，版本号以你的执行结果为准，不要在教程中手写固定版本。

### 验证安装

此时还没有创建 `main.go`，先验证模块是否已解析：

```bash
go list -m github.com/gin-gonic/gin
# 输出模块路径和实际安装的版本
```

完成下一节的 `main.go` 后，再运行 `go run .` 启动服务。

---

## 1.4 第一个 Gin 程序

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建 Gin 引擎（包含 Logger 和 Recovery 中间件）
	r := gin.Default()

	// 2. 注册路由：GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 3. 启动 HTTP 服务，监听 8080 端口
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
```

运行：
```bash
go run main.go
# 输出：
# [GIN-debug] GET    /ping    --> main.main.func1
# [GIN-debug] Listening and serving HTTP on :8080
```

测试：
```bash
curl http://localhost:8080/ping
# {"message":"pong"}
```

> 🎉 恭喜！你的第一个 Gin 服务已经跑起来了！

### 逐行程序解析

```go
r := gin.Default()
```
`gin.Default()` 创建一个 `*gin.Engine`，**自动附带 Logger（日志）和 Recovery（panic 恢复）中间件**。

```go
r.GET("/ping", func(c *gin.Context) { ... })
```
注册一个 GET 路由，路径是 `/ping`。当浏览器或 curl 访问 `GET /ping` 时，执行后面的函数。

```go
func(c *gin.Context)
```
这是**处理函数**（Handler）的签名。`*gin.Context` 是 Gin 的核心——它携带了这次请求的所有信息（参数、请求头、响应写入器），贯穿整个请求处理过程。

```go
c.JSON(http.StatusOK, gin.H{"message": "pong"})
```
返回一个 JSON 响应。`gin.H` 是用于构建 JSON 对象的 map 类型（当前版本等价于 `map[string]any`）：
```go
type H map[string]any
```

```go
if err := r.Run(":8080"); err != nil {
	panic(err)
}
```
启动 HTTP 服务器，监听 8080 端口。`:8080` 通常表示监听所有本地网卡的 8080 端口；开发时可用 `127.0.0.1:8080` 限制为本机访问。`Run` 返回服务器启动或运行错误，正式程序应记录并处理该错误。

---

## 1.5 `gin.Default()` vs `gin.New()`

```go
// Default：包含 Logger + Recovery 中间件
r := gin.Default()

// New：纯净引擎，没有任何中间件
r := gin.New()
```

| 方法 | 中间件 | 何时使用 |
|------|--------|---------|
| `gin.Default()` | Logger + Recovery | 大多数场景（推荐） |
| `gin.New()` | 无 | 需要完全自定义中间件时 |

> 📌 日常开发用 `gin.Default()` 就行。Recovery 能防止 panic 导致服务崩溃，Logger 帮你看到每个请求的日志。

### 验证 Recovery 的作用

下面是两个需要分别运行的对比示例。每个示例都应放在独立的 `main` 函数或独立文件中，不要把它们放在同一个 Go 作用域中：

```go
// 示例一：Default 自带 Recovery
r := gin.Default()
r.GET("/panic", func(c *gin.Context) {
	panic("出错了！")
})
if err := r.Run(":8080"); err != nil {
	panic(err)
}
// 访问 /panic → Recovery 捕获 panic，通常返回 500，当前请求结束；服务进程继续运行
```

```go
// 示例二：New 没有任何中间件
r := gin.New()
r.GET("/panic", func(c *gin.Context) {
	panic("出错了！")
})
if err := r.Run(":8080"); err != nil {
	panic(err)
}
// 没有 Recovery 时，panic 会沿调用栈向上传播；不要依赖底层服务器的隔离行为。
// 生产服务应显式注册 gin.Recovery() 或自定义恢复中间件。
```

---

## 1.6 热加载（开发利器）

每次改代码都要 `Ctrl+C` → `go run` → 再测试，很烦。用热加载工具：

```bash
# 安装 Air（常用的 Go 热加载工具）
go install github.com/air-verse/air@latest

# 在项目目录下运行
air
# 之后每次保存代码，自动重新编译运行！
```

---

## 常见错误

### 错误1：端口被占用

```
listen tcp :8080: bind: address already in use
```

**原因**：8080 端口已经被其他程序占用。

**解决**：换个端口 `r.Run(":8081")` 或者关掉占用 8080 的程序。

### 错误2：导入路径与包名混淆

```go
// ❌ 错误：Gin 的导入路径不是 "gin"
import "gin"

// ✅ Gin 模块的包名是 gin
import "github.com/gin-gonic/gin"
```

### 错误3：在错误目录运行 `go run`

```bash
# 在包含 go.mod 和 main.go 的项目目录运行
go run .
```

> 单文件入口也可以执行 `go run main.go`；如果一个包由多个 `.go` 文件组成，使用 `go run .` 可一起编译该目录下的包文件。

---

## 本章小结

- Web 框架帮你省掉路由匹配、参数解析、响应封装等模板代码
- Gin 是 Go 生态中使用广泛的 Web 框架之一
- `gin.Default()` 自动包含 Logger + Recovery
- 一个 Gin 程序三步走：创建引擎 → 注册路由 → `Run()`
- `gin.Context` 是贯穿请求处理的核心对象
- `gin.H` 是构建 JSON 对象的快捷 map 类型

## 练习题

1. 创建一个 Gin 项目，写一个 `GET /hello` 接口，返回 `{"msg": "你好，世界"}`。
2. 分别用 `gin.Default()` 和 `gin.New()` 创建引擎，访问一个会 panic 的接口，观察区别。
3. 给服务换一个端口（比如 9090），验证是否生效。
4. 安装 Air 热加载工具，体验自动重启。
5. （思考题）为什么 Web 框架要提供中间件机制？直接用函数处理请求不行吗？
