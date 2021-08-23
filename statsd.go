package moka

import (
	"io"
	"log"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	statsdreporter "github.com/ringsaturn/moka/statsd"
	"github.com/uber-go/tally"
)

type MokaStatsd struct {
	Scope       tally.Scope
	ScopeCloser io.Closer
}

func NewMokaStatsd(statsdConfig *statsd.ClientConfig) (*MokaStatsd, error) {
	mokaStatsd := &MokaStatsd{}
	statter, err := statsd.NewClientWithConfig(statsdConfig)
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
