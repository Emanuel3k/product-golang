package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ReadJson[T any](path string) ([]*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result []*T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func WriteJson[T any](path string, data []*T) error {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
