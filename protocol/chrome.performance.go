package protocol

import (
	"encoding/json"

	performance "github.com/mkenney/go-chrome/protocol/performance"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Performance is a struct that provides a namespace for the Chrome Performance protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/
*/
type Performance struct{}

/*
Disable disables collecting and reporting metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-disable
*/
func (Performance) Disable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Performance.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Enable enables collecting and reporting metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-enable
*/
func (Performance) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "Performance.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetMetrics retrieves current values of run-time metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#method-getMetrics
*/
func (Overlay) GetMetrics(
	socket *sock.Socket,
) (performance.GetMetricsResult, error) {
	command := &sock.Command{
		Method: "Performance.getMetrics",
	}
	result := performance.GetMetricsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
OnMetrics adds a handler to the Performance.metrics event. Performance.metrics returns current
values of the metrics.

https://chromedevtools.github.io/devtools-protocol/tot/Performance/#event-metrics
*/
func (Overlay) OnMetrics(
	socket *sock.Socket,
	callback func(event *performance.MetricsEvent),
) {
	handler := sock.NewEventHandler(
		"Performance.metrics",
		func(name string, params []byte) {
			event := &performance.MetricsEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
