package iface

// ContextKey is an interface for context keys
// go:generate mockgen -source=contextkey.go -destination=mocks/contextkey.go -package=mocks
type ContextKey interface {
	String() string
}
