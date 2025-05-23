package adapter

type RootAdapter interface {
	ParentAdapter

	Version() string
	SetVersion(string)
}
