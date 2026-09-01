package router

import (
	"admin/global"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Healthz 存活探针（liveness）：进程能响应即视为存活，不做外部依赖检查
func Healthz(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

// Readyz 就绪探针（readiness）：数据库（集群模式下还包括 Redis）可用才返回 200
func Readyz(ctx *gin.Context) {
	checks := make(map[string]string)
	healthy := true

	// 数据源：MySQL / MongoDB，逐一检查连通性
	if global.AppDB == nil {
		healthy = false
		checks["store"] = "not_initialized"
	} else {
		for _, name := range global.AppDB.MysqlNames() {
			checks["mysql/"+name] = "up"
		}
		for _, name := range global.AppDB.MongoNames() {
			checks["mongodb/"+name] = "up"
		}
		if len(checks) == 0 {
			healthy = false
			checks["store"] = "no_datasource"
		}
		for key := range global.AppDB.Ping(ctx.Request.Context()) {
			healthy = false
			checks[key] = "down"
		}
	}

	// 集群部署（验证码使用 Redis 存储）时检查 Redis
	if strings.EqualFold(strings.TrimSpace(global.AppConfig.Captcha.Store), "redis") {
		if global.AppRedis == nil {
			healthy = false
			checks["redis"] = "not_initialized"
		} else if err := global.AppRedis.Ping(context.Background()).Err(); err != nil {
			healthy = false
			checks["redis"] = "down"
		} else {
			checks["redis"] = "up"
		}
	}

	status := http.StatusOK
	statusText := "ok"
	if !healthy {
		status = http.StatusServiceUnavailable
		statusText = "unavailable"
	}
	ctx.JSON(status, gin.H{"status": statusText, "checks": checks})
}
