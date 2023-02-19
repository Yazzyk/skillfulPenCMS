package utils

//go:generate stringer -type=ENV
type ENV int

const (
	DEV ENV = iota
	TEST
	PROD
)
