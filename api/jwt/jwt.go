package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var jwtKey = []byte("ywuieqwi1238799jsjkadk!@#") // 这是你的密钥，用于签名和验证 JWT

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求未包含 token"})
			c.Abort()
			return
		}

		// 解析并验证 JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 将 claims 存储到上下文中，以便在处理函数中使用
			c.Set("claims", claims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无法解析 claims"})
			c.Abort()
		}

		c.Next()
	}
}

func GenerateToken(userId int64) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //有效时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
		NotBefore: jwt.NewNumericDate(time.Now()),                     //生效时间
		Issuer:    "root",                                             //签发人
		Subject:   "web",                                              //主题
		ID:        "1",                                                //JWT ID用于标识该JWT
		Audience:  []string{"web_root"},                               //用户
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}
