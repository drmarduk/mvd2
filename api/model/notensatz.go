package model

// NotenSatz represents a whole song with metadata and available instruments
type NotenSatz struct {
	ID   int    `json:"id" sql:"id"`
	Name string `json:"name" sql:"name"`
}

// OpenNotenSatz returns a NotenSatz from the database
func OpenNotenSatz(id int) (*NotenSatz, error) {
	result := &NotenSatz{
		ID:   id,
		Name: "Die Vogelwiese",
	}
	return result, nil
}

// NewNotenSatz inserts a new notensatz in the database
func NewNotenSatz(name string) (*NotenSatz, error) {
	result := &NotenSatz{
		ID:   42,
		Name: name,
	}
	return result, nil
}

// Save edit an notensatz
func (n *NotenSatz) Save() error {
	return nil
}

// Delete deletes the entry from the database
func (n *NotenSatz) Delete() error {
	return nil
}
