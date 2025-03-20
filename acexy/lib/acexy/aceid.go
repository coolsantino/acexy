// Helper utility to help finding the AceStream ID

package acexy

import (
	"errors"
	"fmt"
)

type AceID struct {
	content_id string
	infohash string
}

// Type referencing which ID is set
type AceIDType string

// Create a new `AceID` object
func NewAceID(content_id, infohash string) (AceID, error) {
	if content_id == "" && infohash == "" {
		return AceID{}, errors.New("one of `content_id` or `infohash` must have a value")
	}
	if content_id != "" && infohash != "" {
		return AceID{}, errors.New("only one of `content_id` or `infohash` can have a value")
	}
	return AceID{content_id: content_id, infohash: infohash}, nil
}

// Get the valid AceStream ID. If the `infohash` is set, it will be returned,
// otherwise the `id`.
func (a AceID) ID() (AceIDType, string) {
	if a.infohash != "" {
		return "infohash", a.infohash
	}
	return "content_id", a.content_id
}

// Get the AceStream ID as a string
func (a AceID) String() string {
	idType, content_id := a.ID()
	return fmt.Sprintf("{%s: %s}", idType, content_id)
}
