package database

type MemStorage struct {
	data map[string]any
}

func (s *MemStorage) getMertics() map[string]any {
	return s.data
}

func (s *MemStorage) getMetric(key string) (any, error) {
	return nil, nil
}

func (s *MemStorage) setMetric(name string, value any) error {
	return nil
}

func (s *MemStorage) deleteMetric(name string) error {
	return nil
}
