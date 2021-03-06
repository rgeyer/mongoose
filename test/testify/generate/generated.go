package generate

//go:generate mongoose -mock testify -name Generated

// Generated is a test of mongoose when invoked by `go generate`.
type Generated interface {
	Foo(a string, b int) error
}

// NotGenerated should not have a mock generated for it.
type NotGenerated interface {
	Bar() error
}
