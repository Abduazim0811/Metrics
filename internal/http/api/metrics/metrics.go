package metrics

import (
	"expvar"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var version = "1.0.0"

func Metrics() gin.HandlerFunc {
	expvar.NewString("version").Set(version)
	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return runtime.NumGoroutine()
	}))
	expvar.Publish("timestamp", expvar.Func(func() interface{} {
		return time.Now().Unix()
	}))

	totalRequest := expvar.NewInt("total_request_received")
	totalResponse := expvar.NewInt("total_response_sent")
	totalProcessingTime := expvar.NewInt("total_processing_time_us")
	totalResponsesByStatus := expvar.NewMap("total_responses_sent_by_status")

	return func(c *gin.Context) {
		start := time.Now()
		totalRequest.Add(1)

		c.Next()

		duration := time.Since(start).Microseconds()
		totalProcessingTime.Add(duration)
		totalResponse.Add(1)
		totalResponsesByStatus.Add(strconv.Itoa(c.Writer.Status()), 1)
	}
}
