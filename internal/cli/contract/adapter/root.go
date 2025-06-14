package adapter

// RootAdapter defines the interface for the top-level CLI adapter.
type RootAdapter interface {
	ParentAdapter

	// Version returns the application version string.
	Version() string

	// SetVersion sets the application version string.
	SetVersion(version string)
}
