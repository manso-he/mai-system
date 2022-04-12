package regutil

type ctxKey int

const (
	requestIdKey ctxKey = iota
	loggerKey
	storeKey
)
