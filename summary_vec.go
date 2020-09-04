package prom

import (
	"sync"
)

type LazySummaryVec struct {
	name  string
	Value *SummaryVec
	once  sync.Once
}

func NewLazySummaryVec(name string) *LazySummaryVec {
	return &LazySummaryVec{
		name: name,
	}
}

func (c *LazySummaryVec) Get() (*SummaryVec, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetSummaryVec(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type SummaryVecConfig struct {
	SummaryConfig `yaml:",inline"`
	Labels        []string `yaml:"labels"`
}

func (c *SummaryVecConfig) GetLabels() []string {
	return c.Labels
}
