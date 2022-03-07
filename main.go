package main

import (
	"io"

	"github.com/apex/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func mian() {

}

const JaegerSamplerParam = 1 // 采样所有追踪（不能再online环境使用）
const JaegerReportingHost = "121.196.11.0:6831"

type TraceHandler struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

// 将GlobalTracerHandler作为全局变量使用，这样保证代码中使用同一个tracer
var GlobalTracerHandler *TraceHandler

func init() {
	GlobalTracerHandler = InitTracer()
}

// 封装了初始化tracer的方法
func InitTracer() *TraceHandler {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: JaegerSamplerParam,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: JaegerReportingHost,
		},
	}
	// 设置服务名称
	cfg.ServiceName = "jaeger_test"
	// 创建tracer
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Errorf("NewTracer失败")
	}
	return &TraceHandler{
		Tracer: tracer,
		Closer: closer,
	}
}
