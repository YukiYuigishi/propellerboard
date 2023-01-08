package identifier

type GeneratedID string

type Generater interface {
	GenerateID() (GeneratedID, error)
}
