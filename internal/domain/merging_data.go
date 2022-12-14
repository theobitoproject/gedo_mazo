package domain

// MergingData contains all data that
// can be merged into a document
type MergingData map[string]string

// NewMergingData creates a new instance of MergingData
func NewMergingData() (*MergingData, error) {
	return &MergingData{}, nil
}
