package prom

import (
	"sync"
)

type LazyGaugeVec struct {
	name  string
	Value *GaugeVec
	once  sync.Once
}

func NewLazyGaugeVec(name string) *LazyGaugeVec {
	return &LazyGaugeVec{
		name: name,
	}
}

func (c *LazyGaugeVec) Get() (*GaugeVec, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetGaugeVec(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type GaugeVecConfig struct {
	GaugeConfig `yaml:",inline"`
	Labels      []string `yaml:"labels"`
}

func (c *GaugeVecConfig) GetLabels() []string {
	return c.Labels
}
