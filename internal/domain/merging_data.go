package domain

import "fmt"

// MergingData contains all data that
// can be merged into a document
type MergingData map[string]string

// NewMergingData creates a new instance of MergingData
func NewMergingData() (*MergingData, error) {
	return &MergingData{}, nil
}

// SetKey sets a value for a specific key
func (md *MergingData) SetKey(key string, value string) error {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	d := *md
	d[key] = value

	return nil
}
