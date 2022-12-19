## Description

This is a simple repository I'll be using to practice instrumentation using OpenTelemetry and exporting to various backends. The goal is to attempt manual and automatic instrumentation successfully to address different issues I'm facing in actual production code. This repository will be shared in collaboration with team members who are working alongside me.

It is comprised of a simple development environment built using Docker and Tilt. OpenTelemetry is already implemented, as well as an auto-instrumentation library for Gin. Jaeger is currently the only export backend for visualization, but more of your choosing can be added.

Environment variables:
<pre><code>SERVICE_NAME="otel-practice-app"
COLLECTOR_URL="http://jaeger:14268/api/traces"
INSECURE="true"</code></pre>

## Resources
<ul>
<li>https://signoz.io/blog/opentelemetry-gin/</li>
<li>https://nitin-rohidas.medium.com/using-custom-span-attributes-in-opentelemetry-21e1ac33ec4c</li>
<li>https://github.com/open-telemetry/opentelemetry-go/blob/main/example/jaeger/main.go</li>
</ul>
