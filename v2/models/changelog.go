package models

//go:generate genny -tag generated -in=unmarshalling_generics.go -out=changelog_generated.go gen "GenericType=Changelog"

// Changelog reflects the change log of an issue
type Changelog struct {
	Histories []ChangelogHistory `json:"histories,omitempty"`
}

// Unmarshal unmarshals JSON s to Changelog.
func (x *Changelog) Unmarshal(j string) error {
	return nil
}
