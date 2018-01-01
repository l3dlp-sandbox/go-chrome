package tracing

/*
GetCategoriesResult represents the result of calls to Tracing.getCategories.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-getCategories
*/
type GetCategoriesResult struct {
	// A list of supported tracing categories.
	Categories []string `json:"categories"`
}

/*
RecordClockSyncMarkerParams represents Tracing.recordClockSyncMarker parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-recordClockSyncMarker
*/
type RecordClockSyncMarkerParams struct {
	// The ID of this clock sync marker
	SyncID string `json:"syncId"`
}

/*
RequestMemoryDumpResult represents the result of calls to Tracing.requestMemoryDump.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-requestMemoryDump
*/
type RequestMemoryDumpResult struct {
	// GUID of the resulting global memory dump.
	DumpGUID string `json:"dumpGuid"`

	// True if the global memory dump succeeded.
	Success bool `json:"success"`
}

/*
StartParams represents Tracing.start parameters.

https://chromedevtools.github.io/devtools-protocol/tot/Tracing/#method-start
*/
type StartParams struct {
	// Optional. Category/tag filter.
	Categories string `json:"categories,omitempty"`

	// Optional. Tracing options.
	Options string `json:"options,omitempty"`

	// Optional. If set, the agent will issue bufferUsage events at this
	// interval, specified in milliseconds.
	BufferUsageReportingInterval float64 `json:"bufferUsageReportingInterval,omitempty"`

	// Optional. Whether to report trace events as series of dataCollected
	// events or to save trace to a stream (defaults to `ReportEvents`).
	//
	// Allowed values:
	//	- ReportEvents
	//	- ReturnAsStream
	TransferMode string `json:"transferMode,omitempty"`

	// Optional. Trace config.
	TraceConfig *TraceConfig `json:"traceConfig,omitempty"`
}
