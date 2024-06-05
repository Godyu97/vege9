package midware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/cors"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部

		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Header("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "86400")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "86400")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func GinCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginWithContextFunc: func(c *gin.Context, origin string) bool {
			return true
		},
		AllowPrivateNetwork:    true,
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "X-Requested-With", "X-Extra-Header", "X-CSRF-Token", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "If-Modified-Since", "Cache-Control", "Pragma", "FooBar", "Token", "session"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Expires", "Last-Modified", "Pragma", "FooBar"},
		MaxAge:                 24 * time.Hour,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		CustomSchemas:          nil,
		AllowWebSockets:        true,
		AllowFiles:             false,
	})
}
