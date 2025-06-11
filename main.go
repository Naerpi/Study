package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
)

// 路由处理函数
func PingHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "hellowprld",
    })
}

// 主函数
func main() {
    r := gin.Default()
    r.GET("/ping", PingHandler)
    r.Run() // 默认监听 :8080
}

// 测试函数
func TestPingHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.GET("/ping", PingHandler)

    req, _ := http.NewRequest("GET", "/ping", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Fatalf("expected 200, got %d", w.Code)
    }
    if !strings.Contains(w.Body.String(), "hellowprld") {
        t.Fatalf("response does not contain hellowprld")
    }
    fmt.Println("修改")
}