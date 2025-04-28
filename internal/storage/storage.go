package storage

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
}

func (m MemStorage) UpdateCounter(name string, value float64) error {}
