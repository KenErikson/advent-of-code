package common

type RunMode int

const (
	NormalRunMode RunMode = 1 + iota
	InitRunMode
)

type InputMode int

const (
	ExampleInputMode InputMode = 1 + iota
	FullInputMode
)
