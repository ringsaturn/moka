package moka

import (
	"io"
	"log"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/uber-go/tally"
	statsdreporter "github.com/uber-go/tally/statsd"
)

type MokaStatsd struct {
	Scope       tally.Scope
	ScopeCloser io.Closer
}

func NewMokaStatsd() (*MokaStatsd, error) {
	mokaStatsd := &MokaStatsd{}
	statter, err := statsd.NewBufferedClient("127.0.0.1:8125",
		"stats", 100*time.Millisecond, 1440)
	if err != nil {
		log.Fatalf("could not create statsd client: %v", err)
	}

	opts := statsdreporter.Options{}
	r := statsdreporter.NewReporter(statter, opts)

	scope, closer := tally.NewRootScope(tally.ScopeOptions{
		Prefix:   "my-service",
		Tags:     map[string]string{},
		Reporter: r,
	}, 1*time.Second)
	mokaStatsd.Scope = scope
	mokaStatsd.ScopeCloser = closer
	return mokaStatsd, err

	// counter := scope.Counter("test-counter")
	// gauge := scope.Gauge("test-gauge")
	// timer := scope.Timer("test-timer")
	// histogram := scope.Histogram("test-histogram", tally.DefaultBuckets)
}

func (m *MokaStatsd) Counter(name string, val int64) {
	m.Scope.Counter(name).Inc(val)
}

func (m *MokaStatsd) Gauge(name string, val float64) {
	m.Scope.Gauge(name).Update(val)
}

func (m *MokaStatsd) Timer(name string, val time.Duration) {
	m.Scope.Timer(name).Record(val)
}

func (m *MokaStatsd) Close() error {
	return m.ScopeCloser.Close()
}
