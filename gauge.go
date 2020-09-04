package prom

import (
	"sync"
)

type LazyGauge struct {
	name  string
	Value Gauge
	once  sync.Once
}

func NewLazyGauge(name string) *LazyGauge {
	return &LazyGauge{
		name: name,
	}
}

func (c *LazyGauge) Get() (Gauge, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetGauge(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type GaugeConfig struct {
	Opts OptsConfig `yaml:"opts"`
}

func (c *GaugeConfig) GetOpts() *OptsConfig {
	return &c.Opts
}
