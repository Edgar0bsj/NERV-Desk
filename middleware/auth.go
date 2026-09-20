package middleware

import (
	"net/http"
	"strings"

	"github.com/edgar0bsj/nerv-desk/module/user/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 1. Extrair o token do header Authorization.
		authHeader := c.GetHeader("Authorization")

		// 2. Caso o token não exista, retornar Unauthorized.
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token ausente"})
			c.Abort()
			return
		}

		// 3. Validar e decodificar o JWT.
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Validar a assinatura e a validade do token.
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("nead_desk_secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		// 5. Extrair o ID e a role do usuário.
		claims := token.Claims.(jwt.MapClaims)

		// 6. Armazenar as informações do usuário no contexto.
		c.Set("user_id", claims["user_id"])
		c.Set("user_role", claims["user_role"])

		// 7. Permitir que a requisição continue.
		c.Next()
	}
}

func RequireUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extrair o token do header Authorization.
		authHeader := c.GetHeader("Authorization")

		// 2. Caso o token não exista, retornar Unauthorized.
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token ausente"})
			c.Abort()
			return
		}

		// 3. Validar e decodificar o JWT.
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Validar a assinatura e a validade do token.
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("nead_desk_secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		// 5. Extrair o ID e a role do usuário.
		claims := token.Claims.(jwt.MapClaims)

		// 6. Armazenar as informações do usuário no contexto.
		c.Set("user_id", claims["user_id"])
		c.Set("user_role", claims["user_role"])

		// 7. Caso a role seja diferente de "USER", impedir o acesso e retornar Forbidden.
		if claims["user_role"] != model.RoleUser {
			c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			c.Abort()
		}

		// 8. Permitir que a requisição continue.
		c.Next()
	}
}

func RequireAttendantRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extrair o token do header Authorization.
		authHeader := c.GetHeader("Authorization")

		// 2. Caso o token não exista, retornar Unauthorized.
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token ausente"})
			c.Abort()
			return
		}

		// 3. Validar e decodificar o JWT.
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Validar a assinatura e a validade do token.
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("nead_desk_secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		// 5. Extrair o ID e a role do usuário.
		claims := token.Claims.(jwt.MapClaims)

		// 6. Armazenar as informações do usuário no contexto.
		c.Set("user_id", claims["user_id"])
		c.Set("user_role", claims["user_role"])

		// 7. Caso a role seja diferente de "USER", impedir o acesso e retornar Forbidden.
		if claims["user_role"] != model.RoleAttendant {
			c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			c.Abort()
		}

		// 8. Permitir que a requisição continue.
		c.Next()
	}
}

func RequireAdminRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extrair o token do header Authorization.
		authHeader := c.GetHeader("Authorization")

		// 2. Caso o token não exista, retornar Unauthorized.
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token ausente"})
			c.Abort()
			return
		}

		// 3. Validar e decodificar o JWT.
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Validar a assinatura e a validade do token.
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("nead_desk_secret"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		// 5. Extrair o ID e a role do usuário.
		claims := token.Claims.(jwt.MapClaims)

		// 6. Armazenar as informações do usuário no contexto.
		c.Set("user_id", claims["user_id"])
		c.Set("user_role", claims["user_role"])

		// 7. Caso a role seja diferente de "USER", impedir o acesso e retornar Forbidden.
		if claims["user_role"] != model.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			c.Abort()
		}

		// 8. Permitir que a requisição continue.
		c.Next()
	}
}
