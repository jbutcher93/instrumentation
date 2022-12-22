package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	serviceName = "otel-practice-app"
	environment = "jaeger"
)

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("environment", "development"),
			attribute.Int64("ID", 1),
		)),
	)
	return tp, nil
}

func main() {
	var span trace.Span
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if environment == "jaeger" {
		tp, err := tracerProvider(fmt.Sprintf("http://%s:14268/api/traces", environment))
		if err != nil {
			log.Fatal(err)
		}
		// Register our TracerProvider as the global so any imported
		// instrumentation in the future will default to using it.
		otel.SetTracerProvider(tp)
		tr := tp.Tracer("component-main")
		_, span = tr.Start(ctx, "main")
		defer span.End()
	}
	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	r.GET("/ping", result)
	r.Run()
}

func add(ctx context.Context, num1, num2 int) int {
	_, span := otel.Tracer("main").Start(ctx, "add")
	defer span.End()
	return num1 + num2
}

func subtract(ctx context.Context, num1, num2 int) int {
	result := num1 - num2
	_, span := otel.Tracer("main").Start(ctx, "subtract")
	span.SetAttributes(attribute.Int("Result: ", result))
	defer span.End()
	return result
}

func result(c *gin.Context) {
	ctx := c.Request.Context()
	newCtx, span := otel.Tracer("main").Start(ctx, "result")
	defer span.End()
	result1 := add(newCtx, 2, 2)
	result2 := subtract(newCtx, result1, 2)
	c.JSON(200, result2)
}
