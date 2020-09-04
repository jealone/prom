package prom

import (
	"sync"
	"time"
)

type LazySummary struct {
	name  string
	Value Histogram
	once  sync.Once
}

func NewLazySummary(name string) *LazySummary {
	return &LazySummary{
		name: name,
	}
}

func (c *LazySummary) Get() (Summary, error) {
	c.once.Do(func() {
		c.Value = GetRegistry().GetSummary(c.name)
	})
	if nil == c.Value {
		return nil, ErrorMetricNotFound
	}
	return c.Value, nil
}

type SummaryConfig struct {
	Opts SummaryOptsConfig `yaml:"opts"`
}

func (c *SummaryConfig) GetOpts() *SummaryOptsConfig {
	return &c.Opts
}

type SummaryOptsConfig struct {
	OptsConfig `yaml:",inline"`
	Objectives string `yaml:"objectives"`
	MaxAge     int    `yaml:"max_age"`
	AgeBuckets uint32 `yaml:"age_buckets"`
	BufCap     uint32 `yaml:"buf_cap"`
}

func (o *SummaryOptsConfig) GetObjectives() map[float64]float64 {

	objectives := make(map[float64]float64)
	err := json.UnmarshalFromString(o.Objectives, objectives)
	if nil != err {
		panic(err)
	}
	return objectives
}

func (o *SummaryOptsConfig) GetMaxAge() time.Duration {
	return time.Duration(o.MaxAge) * time.Millisecond
}

func (o *SummaryOptsConfig) GetAgeBuckets() uint32 {
	return o.AgeBuckets
}

func (o *SummaryOptsConfig) GetBufCap() uint32 {
	return o.BufCap
}
