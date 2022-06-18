package tracer

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string){
	cfg, err := jaegercfg.FromEnv()
	if err != nil{
		panic(fmt.Sprintf("Could not parse Jaeger env vars: %s\n", err.Error()))
		return
	}
	cfg.ServiceName = service
	tracer, _, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil{
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
		return
	}
	opentracing.InitGlobalTracer(tracer)
}