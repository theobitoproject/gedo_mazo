package domain

// DefaultFileName is the name that is assigned to a document
// that does not specify a name
const DefaultFileName = "Untitled document"

// Document defines an actual document that
// contains data, can be managed and modified and
// is stored in some location or storage
type Document struct {
	id   string
	name string
}

// NewDocument creates a new instance of Document
func NewDocument(
	id string,
	name string,
) (*Document, error) {
	if name == "" {
		name = DefaultFileName
	}

	return &Document{id, name}, nil
}

// NewUnsavedDocument returns a document without an id
// since the id is defined by the location or storage
func NewUnsavedDocument(name string) (*Document, error) {
	return NewDocument("", name)
}

// IsSaved defines if a document has been saved/stored
// in the location or storage
func (d *Document) IsSaved() bool {
	return !d.IsZero() && d.id != ""
}

// Id returns the id of the document
func (d *Document) Id() string {
	return d.id
}

// Name returns the name of the document
func (d *Document) Name() string {
	return d.name
}

// IsZero defines if the document is empty
func (d *Document) IsZero() bool {
	return *d == Document{}
}

// String defines the string interface
func (d *Document) String() string {
	return d.name
}
