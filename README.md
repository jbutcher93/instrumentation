## Description

This is a simple repository I'll be using to practice instrumentation using OpenTelemetry and exporting to various backends. The goal is to attempt manual and automatic instrumentation successfully to address different issues I'm facing in actual production code. This repository will be shared in collaboration with team members who are working alongside me.

It is comprised of a simple development environment built using Docker and Tilt. OpenTelemetry is already implemented, as well as an auto-instrumentation library for Gin. Jaeger is currently the only export backend for visualization, but more of your choosing can be added.

## Resources
<ul>
<li>https://signoz.io/blog/opentelemetry-gin/</li>
<li>https://opentelemetry.io/docs/reference/specification/trace/api/</li>
<li>https://nitin-rohidas.medium.com/using-custom-span-attributes-in-opentelemetry-21e1ac33ec4c</li>
<li>https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go</li>
<li>https://medium.com/jaegertracing/jaeger-tracing-a-friendly-guide-for-beginners-7b53a4a568ca</li>
<li>https://dev.to/aurelievache/learning-go-by-examples-part-10-instrument-your-go-app-with-opentelemetry-and-send-traces-to-jaeger-distributed-tracing-1p4a</li>
</ul>
