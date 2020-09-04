package prom

import (
	"sync"
)

type LazyHistogramVec struct {
	name  string
	Value *HistogramVec
	once  sync.Once
}

func NewLazyHistogramVec(name string) *LazyHistogramVec {
	return &LazyHistogramVec{
		name: name,
	}
}

func (c *LazyHistogramVec) Get() (*HistogramVec, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetHistogramVec(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type HistogramVecConfig struct {
	HistogramConfig `yaml:",inline"`
	Labels          []string `yaml:"labels"`
}

func (c *HistogramVecConfig) GetLabels() []string {
	return c.Labels
}
