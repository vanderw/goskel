package tracer

type Metric struct {
	Name        string
	Unit        string
	Description string
}

var MetricRequestDurationMillis = Metric{
	Name:        "request_duration_millis",
	Unit:        "ms",
	Description: "measures the latency of HTTP requests processed by the server, in milliseconds",
}

var MetricRequestsInFlight = Metric{
	Name:        "requests_in_flight",
	Unit:        "{count}",
	Description: "measures the number of requests currently being processed by the server",
}
