package storage

import (
	err "github.com/sixicebit/metricsalert/internal/errors"
)

type MemStorage struct {
	gauges   map[string]float64
	counters map[string]int64
}

// NewMemStorage создает новый экземпляр MemStorage.
func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauges:   map[string]float64{},
		counters: map[string]int64{}, // Инициализация карты
	}
}

type Storage interface {
	UpdateGauge(name string, value float64) error
	UpdateCounter(name string, value int64) error
}

func (m MemStorage) UpdateGauge(name string, value float64) error {
	if name == "" {
		return err.ErrEmptyMetricName
	}
	m.gauges[name] = value
	return nil
}

func (m MemStorage) UpdateCounter(name string, value int64) error {
	if name == "" {
		return err.ErrEmptyMetricName
	}

	m.counters[name] += value

	return nil
}
