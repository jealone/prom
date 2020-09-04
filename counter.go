package prom

import (
	"sync"
)

type LazyCounter struct {
	name  string
	Value Counter
	once  sync.Once
}

func NewLazyCounter(name string) *LazyCounter {
	return &LazyCounter{
		name: name,
	}
}

func (c *LazyCounter) Get() (Counter, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetCounter(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type CounterConfig struct {
	Opts OptsConfig `yaml:"opts"`
}

func (c *CounterConfig) GetOpts() *OptsConfig {
	return &c.Opts
}
