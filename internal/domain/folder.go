package domain

import "fmt"

// Folder defines an actual folder that
// contains data, can be managed and modified and
// is stored in some location or storage
type Folder struct {
	id string
}

// NewFolder creates a new instance of Folder
func NewFolder(id string) (*Folder, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}

	return &Folder{id}, nil
}

// Id returns the id of the folder
func (d *Folder) Id() string {
	return d.id
}
