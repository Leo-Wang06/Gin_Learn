package chapter23

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key-32bytes!!!!!!!")

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-tutorial",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的Token")
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "请先登录"})
			return
		}

		// Bearer xxx → 取 xxx 部分
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "认证格式错误"})
			return
		}

		// 解析 Token
		claims, err := ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "Token无效或已过期"})
			return
		}

		// 存入 Context（后续 Handler 可获取）
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

func LearnCode() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Username != "admin" || req.Password != "123456" {
			c.JSON(401, gin.H{"code": 401, "message": "用户名或密码错误"})
			return
		}

		token, _ := GenerateToken(1, req.Username)
		c.JSON(200, gin.H{
			"code": 0,
			"message": "登录成功",
			"data": gin.H{"token": token, "expires_in": 86400},
		})
	})

	auth := r.Group("/api")
	auth.Use(JWTAuthMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			username, _ :=c.Get("username")
			c.JSON(200, gin.H{"user_id": userID, "username": username})
		})
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}