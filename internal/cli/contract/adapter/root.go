package adapter

// RootAdapter defines the interface for the root CLI adapter.
type RootAdapter interface {
	ParentAdapter

	// Version returns the CLI application version.
	Version() string

	// SetVersion sets the CLI application version.
	SetVersion(version string)
}
