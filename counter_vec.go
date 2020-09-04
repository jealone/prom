package prom

import (
	"sync"
)

type LazyCounterVec struct {
	name  string
	Value *CounterVec
	once  sync.Once
}

func NewLazyCounterVec(name string) *LazyCounterVec {
	return &LazyCounterVec{
		name: name,
	}
}

func (c *LazyCounterVec) Get() (*CounterVec, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetCounterVec(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type CounterVecConfig struct {
	CounterConfig `yaml:",inline"`
	Labels        []string `yaml:"labels"`
}

func (c *CounterVecConfig) GetLabels() []string {
	return c.Labels
}
