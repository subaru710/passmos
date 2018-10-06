package passport

type ExternalStore interface {
	Type() string
	GetPersonalData() *PersonalData
	SetPersonalData(PersonalData) (string, error)
	HasPersonalData() bool
}
