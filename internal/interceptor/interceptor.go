package interceptor

import (
	"context"
	"github.com/MeinAJ/greeter-server/internal/metrics" // 导入你的指标定义包
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 饱和度：请求开始，计数器+1
	method := info.FullMethod
	metrics.InFlightRequests.Inc(method)
	defer metrics.InFlightRequests.Dec(method) // 请求结束，计数器-1

	// 记录开始时间
	startTime := time.Now()

	// 调用实际的业务逻辑
	resp, err = handler(ctx, req)

	// 延迟：计算耗时并上报
	duration := time.Since(startTime).Milliseconds()
	metrics.ReqDuration.Observe(duration, method)

	// 请求量：总请求数+1
	metrics.ReqTotal.Inc(method)

	// 错误率：如果有错误，错误数+1，并带上错误码
	if err != nil {
		// 获取 gRPC 状态码作为错误标签
		st, _ := status.FromError(err)
		code := st.Code().String()
		metrics.ErrTotal.Inc(method, code)
	}

	return resp, err
}
