package jaeger

import (
	"fmt"
	"io"

	"github.com/gobuffalo/envy"
	opentracing "github.com/opentracing/opentracing-go"
	jaegerClient "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// Tracer is the opentracing instance
var Tracer opentracing.Tracer

// JaegerHost is the address of the jaeger server, default localhost
var JaegerHost = envy.Get("JAEGER_AGENT_HOST", "127.0.0.1")

// JaegerPort is the UDP port of the jaeger server, default 6831
var JaegerPort = envy.Get("JAEGER_AGENT_PORT", "6831")

// JaegerService is the name of this process that will be reported to jaeger
var JaegerService = envy.Get("JAEGER_SERVICE_NAME", "kalicoin")

// Init creates a new jaeger instance and returns an io.Closer or an error
func Init() (io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           envy.Get("ENVIRONMENT", "development") == "development",
			LocalAgentHostPort: fmt.Sprintf("%v:%v", JaegerHost, JaegerPort),
		},
	}

	var err error
	var closer io.Closer
	Tracer, closer, err = cfg.New(JaegerService, jaegercfg.Logger(jaegerClient.StdLogger))

	if err != nil {
		return nil, err
	}

	return closer, nil
}
