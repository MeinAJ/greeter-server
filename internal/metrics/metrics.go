package metrics

import "github.com/zeromicro/go-zero/core/metric"

const (
	// Namespace 定义指标 namespace 和 subsystem
	Namespace = "rpc_server"
	Subsystem = "greeter" // 替换为你的服务名
)

// ReqTotal 1. 定义一个 CounterVec 来统计按方法和方法名分组的请求总数
var ReqTotal = metric.NewCounterVec(&metric.CounterVecOpts{
	Namespace: Namespace,
	Subsystem: Subsystem,
	Name:      "requests_total",
	Help:      "RPC 总请求数",
	Labels:    []string{"method"}, // 标签：接口方法名
})

// ErrTotal 2. 定义一个 CounterVec 来统计按方法和方法名分组的错误数
var ErrTotal = metric.NewCounterVec(&metric.CounterVecOpts{
	Namespace: Namespace,
	Subsystem: Subsystem,
	Name:      "errors_total",
	Help:      "RPC 错误总数",
	Labels:    []string{"method", "code"}, // 标签：接口方法名和错误码
})

// ReqDuration 3. 定义一个 HistogramVec 来统计请求耗时
var ReqDuration = metric.NewHistogramVec(&metric.HistogramVecOpts{
	Namespace: Namespace,
	Subsystem: Subsystem,
	Name:      "request_duration_milliseconds",
	Help:      "RPC 请求耗时（毫秒）",
	Labels:    []string{"method"},
	Buckets:   []float64{1, 5, 10, 25, 50, 100, 250, 500, 1000, 2000, 5000}, // 自定义耗时桶
})

// InFlightRequests 4. (可选) 定义一个 GaugeVec 来统计当前正在处理的请求数，用于饱和度监控
var InFlightRequests = metric.NewGaugeVec(&metric.GaugeVecOpts{
	Namespace: Namespace,
	Subsystem: Subsystem,
	Name:      "in_flight_requests",
	Help:      "当前正在处理的请求数",
	Labels:    []string{"method"},
})
