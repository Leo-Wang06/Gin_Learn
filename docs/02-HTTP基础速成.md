# 第02章：HTTP 基础速成（为 Gin 学习准备）

## 本章目标
学完本章后，你将能够：
1. 画出 HTTP 请求-响应模型的全流程
2. 区分 GET/POST/PUT/PATCH/DELETE 的用途
3. 根据状态码（2xx/3xx/4xx/5xx）判断请求结果
4. 理解 URL 各部分（协议、主机、端口、路径、查询参数）
5. 区分三种最常见的 Content-Type
6. 用浏览器 DevTools 和 curl 观察 HTTP 请求

## 前置知识
- 需要先学习：第01章（Gin 入门）
- 无 HTTP 基础要求，本章从零讲起
- 本章只介绍 HTTP 语义；“HTTP 方法对应数据库操作”是常见 API 约定，不是协议强制要求

---

## 2.1 HTTP 请求-响应模型

HTTP（超文本传输协议）是 Web 的基石。它的工作方式很简单：**客户端发请求，服务端返响应**。

```
  客户端（浏览器/curl）                  服务端（Gin 程序）
         │                                    │
         │  ──── GET /users HTTP/1.1 ────>    │
         │       Host: localhost:8080          │
         │       Accept: application/json      │
         │                                    │
         │                                    │  处理请求
         │                                    │  查数据库...
         │                                    │
         │  <──── HTTP/1.1 200 OK ────        │
         │       Content-Type: application/json│
         │       {"name":"张三","age":25}      │
         │                                    │
```

一次 HTTP 交互 = 一个请求（Request）+ 一个响应（Response）

---

## 2.2 HTTP 方法（Method）

HTTP 方法告诉服务端"我要做什么"：

| 方法 | 含义 | 常见后端实现 | Gin 注册方式 | 示例 |
|------|------|-------------|------------|------|
| **GET** | 获取资源 | 常见为查询数据 | `r.GET()` | 查看用户列表 |
| **POST** | 创建资源或触发一次处理 | 常见为新增数据 | `r.POST()` | 注册新用户 |
| **PUT** | 通常表示用请求内容替换/创建目标资源的当前表示 | 常见为全量更新 | `r.PUT()` | 替换用户全部信息 |
| **PATCH** | 对目标资源做部分修改 | 常见为部分更新 | `r.PATCH()` | 只改用户昵称 |
| **DELETE** | 请求删除资源 | 常见为删除或标记删除 | `r.DELETE()` | 删除用户 |
| OPTIONS | 询问目标资源支持的通信选项；浏览器跨域时常用于预检 | 通常返回允许的方法/请求头 | `r.OPTIONS()` 或中间件 | CORS 预检 |

> 📌 常见 API 约定：GET 查、POST 增、PUT 通常全量替换、PATCH 通常部分修改、DELETE 删；具体语义由接口契约定义。

```bash
# 用 curl 发送不同方法的请求
curl -X GET    http://localhost:8080/users        # 查
curl -X POST   http://localhost:8080/users        # 增
curl -X PUT    http://localhost:8080/users/1      # 改
curl -X DELETE http://localhost:8080/users/1      # 删
```

---

## 2.3 HTTP 状态码

状态码告诉客户端"结果怎么样"：

| 状态码 | 分类 | 含义 | 示例 |
|--------|------|------|------|
| **200** | 2xx 成功 | OK，请求成功 | 查询成功 |
| **201** | 2xx 成功 | Created，创建成功 | 注册成功 |
| **204** | 2xx 成功 | No Content，成功但无返回体 | 删除成功 |
| **301** | 3xx 重定向 | 永久重定向 | 旧URL搬到新URL |
| **302** | 3xx 重定向 | 临时重定向 | 登录后跳转 |
| **400** | 4xx 客户端错误 | Bad Request，请求格式或参数有误 | JSON 格式错误、缺少必填字段 |
| **401** | 4xx 客户端错误 | Unauthorized，未认证 | 没登录 |
| **403** | 4xx 客户端错误 | Forbidden，无权限 | 权限不足 |
| **404** | 4xx 客户端错误 | Not Found，资源不存在 | 查的用户不存在 |
| **409** | 4xx 客户端错误 | Conflict，冲突 | 用户名已存在 |
| **422** | 4xx 客户端错误 | Unprocessable Content（旧称 Unprocessable Entity） | 语法正确但语义无法处理 |
| **500** | 5xx 服务端错误 | Internal Server Error | 代码 panic |

> 📌 便于记忆：2xx 表示成功，3xx 表示重定向，4xx 通常表示请求侧问题，5xx 通常表示服务端处理失败。这是概括，不代表所有具体场景都能简单归责给客户端或服务器。

在 Gin 中使用：

```go
// data/msg 代表你准备返回的 JSON 值；这里用 gin.H 举例。
c.JSON(http.StatusOK, gin.H{"data": []string{}})                 // 200
c.JSON(http.StatusCreated, gin.H{"message": "created"})          // 201
c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"}) // 400
c.JSON(http.StatusNotFound, gin.H{"error": "not found"})         // 404
c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"}) // 500
```

---

## 2.4 URL 结构

```
https://api.example.com:8080/users/1?name=张三&age=25#section
└─┬─┘ └──────┬──────┘└─┬┘ └────┬────┘ └──────┬──────┘ └──┬──┘
 协议      主机       端口  路径          查询参数       锚点
(scheme)  (host)    (port) (path)      (query)       (fragment)
```

| 部分 | 说明 | 示例 |
|------|------|------|
| 协议 | http 或 https | `https://` |
| 主机 | 域名或 IP | `localhost`、`api.example.com` |
| 端口 | 服务端口（HTTP默认80，HTTPS默认443） | `:8080` |
| 路径 | 资源位置 | `/users/1` |
| 查询参数 | `key=value&key2=value2`；会发送给服务器 | `?name=张三&age=25` |
| 片段（fragment） | 浏览器端定位信息，不会包含在 HTTP 请求中 | `#section` |

在 Gin 中获取各部分：
```go
c.Request.Host        // "localhost:8080"
c.Request.URL.Path    // "/users/1"
c.Request.URL.Query() // map[name:[张三] age:[25]]
c.Param("id")        // "1" （路径参数）
c.Query("name")      // "张三" （查询参数）
```

---

## 2.5 Content-Type：数据的格式标签

Content-Type 告诉服务端"我发的是什么格式的数据"：

| Content-Type | 场景 | Gin 获取方式 |
|-------------|------|------------|
| `application/json` | 前后端分离 API（最常见） | `c.ShouldBindJSON()` |
| `application/x-www-form-urlencoded` | HTML 表单提交 | `c.PostForm()` |
| `multipart/form-data` | 文件上传 | `c.FormFile()` |

### JSON 请求示例

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","age":25}'
```

### 表单请求示例

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=zhangsan&password=123456"
```

---

## 2.6 RESTful API 设计原则速览

| 原则 | 说明 | 示例 |
|------|------|------|
| 用名词不用动词 | URL 表示资源 | ✅ `/users` ❌ `/getUsers` |
| 用 HTTP 方法表示操作 | GET查/POST增/PUT改/DELETE删 | ✅ `GET /users/1` |
| 资源命名保持一致 | 集合端点常用复数 | ✅ `/users`（团队内保持一致即可） |
| 嵌套资源表示关系 | 层级表示从属 | `GET /users/1/orders` |
| 版本控制 | API 版本号 | `/api/v1/users` |

---

## 2.7 用 curl 观察 HTTP

```bash
# -v 参数显示完整的请求和响应细节
curl -v http://localhost:8080/ping
```

输出：
```
> GET /ping HTTP/1.1          ← 请求行
> Host: localhost:8080        ← 请求头
> User-Agent: curl/8.0
>
< HTTP/1.1 200 OK             ← 响应行
< Content-Type: application/json  ← 响应头
< Date: Mon, 15 Jan 2024 ...
<
{"message":"pong"}            ← 响应体
```

---

## 常见错误

### 错误1：GET 请求带了 Body

```bash
# ⚠️ 不建议把业务参数放在 GET Body 中
curl -X GET http://localhost:8080/users -d '{"name":"张三"}'
# HTTP 语义没有为 GET Body 定义通用处理方式；不同客户端、代理和服务端可能忽略或拒绝它。
# 查询条件应优先放在 URL 查询参数中：/users?name=张三
```

### 错误2：Content-Type 忘记设置

```bash
# ❌ 发送 JSON 但没声明 Content-Type
curl -X POST http://localhost:8080/users -d '{"name":"张三"}'
# 服务端可能无法正确解析

# ✅ 正确
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三"}'
```

### 错误3：状态码乱用

```go
// ❌ 查询不到就返回 500？
c.JSON(500, gin.H{"error": "用户不存在"})
// 用户不存在是你说的，不是服务器内部错误！

// ✅ 正确：客户端请求了不存在的资源 → 404
c.JSON(404, gin.H{"error": "用户不存在"})
```

---

## 本章小结

- HTTP 模型：客户端发 Request → 服务端返 Response
- 常见 API 方法约定：GET(查)、POST(增)、PUT(通常全量替换)、PATCH(通常部分修改)、DELETE(删)
- 状态码概括：2xx 成功、3xx 重定向、4xx 请求侧问题、5xx 服务端处理失败
- URL = 协议 + 主机 + 端口 + 路径 + 查询参数
- JSON 是前后端分离的最常用数据格式
- RESTful API：名词表示资源，HTTP 方法表示操作

## 练习题

1. 用 curl -v 访问一个网站（如 `curl -v https://www.baidu.com`），观察请求和响应的完整内容。
2. 指出以下 API 设计的错误并修正：
   - `POST /getUserById`
   - `GET /deleteUser?id=1`
   - `POST /createNewOrder`
3. 写出以下场景应该使用的 HTTP 方法和状态码：
   - 查询用户列表成功
   - 创建订单成功
   - 访问不存在的文章
   - 未登录访问需要登录的接口
4. （思考题）为什么 RESTful 推荐用复数（`/users`）而不是单数（`/user`）？
