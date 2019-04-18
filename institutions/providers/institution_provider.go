package providers

type InstitutionProvider interface {
	Provide() ([]string, error)
}
