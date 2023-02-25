package cache

import "time"

type Cache interface {
	Exists(id string) (bool, error)
	Set(key, value string, expr time.Duration) error
	Get(key string) (string, error)
	GetAllKeys() ([]string, error)
	Delete(key string) (int64, error)
}

type universalClient struct {
	DataInMemory map[string]string
}

func NewCache() Cache {
	return &universalClient{
		DataInMemory: make(map[string]string),
	}
}

func (r *universalClient) Exists(key string) (bool, error) {
	_, err := r.DataInMemory[key]
	//if err != nil {
	//	return false, nil
	//}
	return err, nil
}

func (r *universalClient) Set(key, value string, expr time.Duration) error {
	r.DataInMemory[key] = value
	return nil
}

func (r *universalClient) Get(key string) (string, error) {
	return r.DataInMemory[key], nil
}

func (r *universalClient) Delete(key string) (int64, error) {
	return int64(10), nil
}

func (r *universalClient) GetAllKeys() ([]string, error) {
	var result []string
	for _, data := range r.DataInMemory {
		result = append(result, data)
	}
	return result, nil
}
