# 第03章：路由基础 —— GET/POST/PUT/DELETE

## 本章目标
学完本章后，你将能够：
1. 注册 GET/POST/PUT/DELETE/PATCH 五种路由
2. 理解 `gin.Context` 的核心作用
3. 返回 String、JSON、HTML 三种格式的响应
4. 用 curl 和 Postman 测试 API
5. 理解路由方法如何映射到 CRUD 操作

## 前置知识
- 需要先学习：第01-02章（Gin入门、HTTP基础）

---

## 3.1 路由注册全览

Gin 提供了对应 HTTP 方法的路由注册函数：

```go
r := gin.Default()

r.GET("/users", handler)       // 查询
r.POST("/users", handler)      // 创建
r.PUT("/users/:id", handler)   // 完整更新
r.PATCH("/users/:id", handler) // 部分更新
r.DELETE("/users/:id", handler)// 删除

// 还有一个特殊的
r.Any("/ping", handler)        // 匹配所有 HTTP 方法
```

---

## 3.2 `gin.Context`：请求处理上下文

每个请求进入 Gin 后，框架会为其准备一个 `*gin.Context`，其中包含请求、响应写入器、路由参数和处理链状态。它通常只在当前请求处理期间使用；不要把它保存到 goroutine 或请求生命周期之外。

```go
func handler(c *gin.Context) {
	// 读：请求信息
	method := c.Request.Method   // 例如 "GET"
	path := c.Request.URL.Path // 例如 "/users"
	clientIP := c.ClientIP()
	_ = method
	_ = path
	_ = clientIP

	// 当前请求内跨中间件传值
	c.Set("key", "value")
	value, exists := c.Get("key")
	if exists {
		_ = value
	}

	// 写：一个请求通常选择一种响应格式
	c.JSON(http.StatusOK, gin.H{"status": "ok"}) // JSON
	// 或：c.String(http.StatusOK, "hello")      // 纯文本
	// 或：c.HTML(http.StatusOK, "tpl", nil)     // 先加载模板后渲染 HTML
}
```

`Next` 和 `Abort` 主要用于中间件，而不是普通业务 Handler。下面的代码片段还需要导入 `log`、`net/http`、`time` 和 `github.com/gin-gonic/gin`：

```go
func TimingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // 执行后续处理器；返回后继续执行当前中间件的后半段
		log.Printf("耗时: %s", time.Since(start))
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatus(http.StatusUnauthorized) // 终止后续处理器
			return // Abort 不会终止当前 Go 函数
		}
		c.Next()
	}
}
```

> 📌 `*gin.Context` 是 Gin 中使用频率最高的对象。后续每一章都在和它打交道。

---

## 3.3 GET 路由

### 返回字符串

```go
r.GET("/hello", func(c *gin.Context) {
	c.String(http.StatusOK, "你好，世界！")
})
// curl http://localhost:8080/hello
// 输出：你好，世界！
```

### 返回 HTML

先在项目的 `templates/welcome.html` 创建模板文件，再加载模板并注册路由：

```html
<!doctype html>
<html lang="zh-CN">
<body><h1>你好，{{.name}}</h1></body>
</html>
```

```go
r.LoadHTMLGlob("templates/*.html")
r.GET("/welcome", func(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{"name": "张三"})
})
```

这是路由片段，假设 `r` 已初始化且已导入 `net/http` 和 Gin。

### 返回 JSON

```go
// 方式一：gin.H（最常用）
r.GET("/user", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "张三",
		"age":  25,
	})
})
// curl http://localhost:8080/user
// {"age":25,"name":"张三"}

// 方式二：传入结构体
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
r.GET("/user2", func(c *gin.Context) {
	c.JSON(http.StatusOK, User{Name: "张三", Age: 25})
})

// 方式三：传入 map
r.GET("/user3", func(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"name": "张三",
		"age":  25,
	})
})
```

---

## 3.4 POST 路由

```go
// 接收 POST 请求，返回创建成功的 JSON
r.POST("/users", func(c *gin.Context) {
	// 实际项目中，这里会解析 Body、存入数据库
	c.JSON(http.StatusCreated, gin.H{
		"message": "用户创建成功",
		"id":      1,
	})
})
```

```bash
curl -X POST http://localhost:8080/users
# {"id":1,"message":"用户创建成功"}
```

> 📌 创建成功通常返回 `201 Created`，而不是 `200 OK`。

---

## 3.5 PUT 和 PATCH 路由

```go
// PUT：完整更新（替换资源的所有字段）
r.PUT("/users/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("用户 %s 已完整更新", id),
	})
})

// PATCH：部分更新（只更新传入的字段）
r.PATCH("/users/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("用户 %s 已部分更新", id),
	})
})
```

---

## 3.6 DELETE 路由

```go
r.DELETE("/users/:id", func(c *gin.Context) {
	id := c.Param("id")
	// 删除成功后通常返回 200 + 消息，或 204 No Content
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("用户 %s 已删除", id),
	})
})
```

---

## 3.7 完整示例：迷你 CRUD API

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// GET     /users     → 用户列表
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": []gin.H{
				{"id": 1, "name": "张三"},
				{"id": 2, "name": "李四"},
			},
		})
	})

	// GET     /users/:id → 用户详情
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": fmt.Sprintf("用户%s", id),
		})
	})

	// POST    /users     → 创建用户
	r.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"id":      3,
			"message": "用户创建成功",
		})
	})

	// PUT     /users/:id → 完整更新用户
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("用户%s完整更新", id),
		})
	})

	// PATCH   /users/:id → 部分更新用户
	r.PATCH("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("用户%s部分更新", id),
		})
	})

	// DELETE  /users/:id → 删除用户
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("用户%s已删除", id),
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
```

测试所有接口：
```bash
curl http://localhost:8080/users              # 列表
curl http://localhost:8080/users/1            # 详情
curl -X POST http://localhost:8080/users      # 创建
curl -X PUT http://localhost:8080/users/1     # 更新
curl -X PATCH http://localhost:8080/users/1  # 部分更新
curl -X DELETE http://localhost:8080/users/1  # 删除
```

---

## 3.8 路由方法与 CRUD 的对应关系

下面是常见 Web API 的映射，不是 HTTP 协议对数据库实现的强制规定：

```
HTTP方法  →  常见数据操作   →  业务含义
────────────────────────────────────────
GET       →  SELECT        →  查询/读取
POST      →  INSERT        →  创建/新增
PUT       →  UPDATE        →  通常全量替换
PATCH     →  UPDATE        →  通常部分修改
DELETE    →  DELETE        →  删除或标记删除
```

> 实际项目中，POST 也可以触发更新，PUT 也可以用于创建可幂等资源；关键是遵守接口契约，并让客户端能够稳定理解语义。

---

## 常见错误

### 错误1：用 `r.GET` 注册 POST 路由

```go
// ❌ GET 方法注册了创建操作
r.GET("/users", createUser)  // 创建应该用 POST！

// ✅ 创建用 POST
r.POST("/users", createUser)
```

### 建议：优先使用状态码常量

```go
// 数字可以工作，但常量更容易读懂
c.JSON(http.StatusOK, data)
```

### 错误3：没有按接口契约选择状态码

```go
// 创建资源成功时常用 201
c.JSON(http.StatusCreated, data)

// 200 也可能是合理选择，例如接口返回的是一个已经完成的操作结果；
// 关键是客户端和服务端要约定一致。
```

### 错误4：把 Handler 和长时间任务混在一起

```go
// Handler 返回后，net/http 会结束响应；如果没有写响应，通常是空响应（默认状态码 200），不是一直挂起。
r.GET("/users", func(c *gin.Context) {
	return
})

// ✅ 明确写出响应内容和状态码
r.GET("/users", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
})
```

> 如果在 Handler 中执行耗时任务，客户端会等待任务完成。应设置超时、使用异步任务或把任务移出请求链，并注意请求上下文取消。

---

## 本章小结

| 方法 | Gin 注册 | 用途 | 成功状态码 |
|------|---------|------|----------|
| GET | `r.GET()` | 查询 | 200 |
| POST | `r.POST()` | 创建 | 201 |
| PUT | `r.PUT()` | 全量更新 | 200 |
| PATCH | `r.PATCH()` | 部分更新 | 200 |
| DELETE | `r.DELETE()` | 删除 | 200/204 |

- `*gin.Context` 是请求处理的核心对象
- `gin.H` 是构建 JSON 的快捷方式
- `r.Run(":8080")` 启动服务并返回运行错误；`:8080` 通常监听所有本地网卡的 8080 端口

## 练习题

1. 编写一个完整的 CRUD API 骨架（用户资源），包含 5 个接口。
2. 用 curl 测试每一个接口，确认返回正确的状态码。
3. 用 `http.StatusOK` 等常量替换所有硬编码的状态码数字。
4. 写一个 `GET /hello?name=张三` 接口，返回 `{"greeting": "你好，张三"}`。
5. （思考题）`PATCH` 和 `PUT` 的区别是什么？为什么实际项目中 PATCH 用得更多？
