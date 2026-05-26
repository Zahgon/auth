package observability

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func getChiRoutePattern(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// traceChiRoutesSafely attempts to extract the Chi RouteContext. If the
// request does not have a RouteContext it will recover from the panic and
// attempt to figure out the route from the URL's path.
func traceChiRoutesSafely(r *http.Request) { _ = "STUB: not implemented"; return }

// traceChiRouteURLParamsSafely attempts to extract the Chi RouteContext
// URLParams values for the route and assign them to the tracing span. If the
// request does not have a RouteContext it will recover from the panic and not
// set any params.
func traceChiRouteURLParamsSafely(r *http.Request) { _ = "STUB: not implemented"; return }

type interceptingResponseWriter struct {
	writer http.ResponseWriter

	statusCode int
}

func (w *interceptingResponseWriter) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (w *interceptingResponseWriter) Write(data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *interceptingResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *

	// countStatusCodesSafely counts the number of HTTP status codes per route that
	// occurred while GoTrue was running. If it is not able to identify the route
	// via chi.RouteContext(ctx).RoutePattern() it counts with a noroute attribute.
	new(http.Header)
}

func countStatusCodesSafely(w *interceptingResponseWriter, r *http.Request, counter metric.Int64Counter) {
	_ = "STUB: not implemented"
	return
}

func addMetricAttributes(r *http.Request) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// RequestTracing returns an HTTP handler that traces all HTTP requests coming
// in. Supports Chi routers, so this should be one of the first middlewares on
// the router.
func RequestTracing() func(http.Handler) http.Handler { _ = "STUB: not implemented"; return nil }

// there is a vulnerability with otelhttp where
// User-Agent strings are kept in RAM indefinitely and
// can be used as an easy way to resource exhaustion;
// so this code strips the User-Agent header before
// it's passed to be traced by otelhttp, and then is
// returned back to the middleware
// https://github.com/supabase/gotrue/security/dependabot/11
