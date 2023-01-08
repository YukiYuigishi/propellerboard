package uuid3

import (
	"github.com/google/uuid"
	"propellerboard/util/identifier"
)

type Uuid3Generator struct{}

func NewUuid3Generator() identifier.Generater {
	return Uuid3Generator{}
}

func (e *Uuid3Generator) GeneratedID() (identifier.GeneratedID, error) {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return identifier.GeneratedID(newUuid.String()), nil
}
