package uuid3

import (
	"github.com/google/uuid"
	"propellerboard/util/identifier"
)

type Uuid5Generator struct{}

func NewUuid5Generator() Uuid5Generator {
	return Uuid5Generator{}
}

func (e *Uuid5Generator) GenerateID() (identifier.GeneratedID, error) {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return identifier.GeneratedID(newUuid.String()), nil
}
