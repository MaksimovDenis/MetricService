package storage

import (
	"errors"
	"strconv"
	"strings"
)

type Repository interface {
	AddMetrics(reqKey string, reqValue any) error
}

type MemStorage struct {
	Storage map[string]any
}

func New(data map[string]any) *MemStorage {
	return &MemStorage{Storage: data}
}

func (ms *MemStorage) AddMetrics(reqKey string, reqValue any) error {
	if value, ok := ms.Storage[reqKey]; ok {
		processValue, err := TypeCaster(value, reqValue)
		if err != nil {
			return err
		}

		ms.Storage[reqKey] = processValue
	} else {
		ms.Storage[reqKey] = reqValue
	}

	return nil
}

func TypeCaster(value, newValue any) (any, error) {
	switch v := value.(type) {
	case float64:
		if n, ok := newValue.(float64); ok {
			return n, nil
		}

		if n, ok := newValue.(int64); ok {
			return float64(n), nil
		}

		return nil, errors.New("invalid type of metric")
	case int64:
		switch n := newValue.(type) {
		case int64:
			return v + n, nil
		case float64:
			return v + int64(n), nil
		default:
			return nil, errors.New("invalid type of metric")
		}
	default:
		return newValue, nil
	}
}

func StringConverter(value string) (any, error) {
	if strings.Contains(value, ".") {
		result, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	return int64(result), nil
}
