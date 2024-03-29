package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// Jaeger returns a middleware for the Gin http server
// This middleware creates spans for requests and reports context errors
func Jaeger(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {

		carrier := opentracing.HTTPHeadersCarrier(c.Request.Header)
		ctx, err := tracer.Extract(opentracing.HTTPHeaders, carrier)

		var span opentracing.Span
		if err != nil {
			span = tracer.StartSpan(c.Request.Method + ":" + c.Request.URL.Path)
		} else {
			span = tracer.StartSpan(c.Request.Method+":"+c.Request.URL.Path, ext.RPCServerOption(ctx))
		}

		defer span.Finish()

		span.SetTag("method", c.Request.Method)
		span.SetTag("path", c.Request.URL.Path)
		span.SetTag("client-ip", c.ClientIP())

		c.Next()

		if len(c.Errors) > 0 {
			errors := getErrorString(c.Errors)
			span.SetTag("error", true)
			span.SetBaggageItem("error", errors)
		}
	}
}

// getErrorString takes an array of gin errors and returns it as a trimmed string.
// Without the trimming, there is a trailing newline
func getErrorString(errors interface{}) string {
	return strings.TrimSpace(fmt.Sprint(errors))
}
