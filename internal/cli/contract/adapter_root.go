package contract

type RootAdapter interface {
	ParentAdapter

	Version() string
	SetVersion(string)
}
