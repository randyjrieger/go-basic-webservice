package telemetryService

/*

var (
	telemetryClient               appinsights.TelemetryClient
	AppInsightsInstrumentationKey = "1eae044e-9de9-4c76-b775-e6e8e9414241"
)

func init() {
	flag.StringVar(&AppInsightsInstrumentationKey, "instrumentationKey", AppInsightsInstrumentationKey, "set instrumentation key from azure portal")
	telemetryClient = appinsights.NewTelemetryClient(AppInsightsInstrumentationKey)
	telemetryClient.Context().Tags.Cloud().SetRole("hello-world")
	appinsights.NewDiagnosticsMessageListener(func(msg string) error {
		log.Printf("[%s] %s\n", time.Now().Format(time.UnixDate), msg)
		return nil
	})
}

func responseToRequest(h func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UTC()
		h(w, r)
		duration := time.Now().Sub(startTime)
		request := appinsights.NewRequestTelemetry(r.Method, r.URL.Path, duration, "200")
		request.Timestamp = time.Now().UTC()
		telemetryClient.Track(request)
	})
}
func cannedResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Request completed.`))
}*/
