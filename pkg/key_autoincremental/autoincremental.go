package key_autoincremental

type Incremental interface {
	Next() int64
}

type incremental struct {
}

var KEY int64 = 0

func New() Incremental {
	KEY = 0
	return &incremental{}
}

func (i incremental) Next() int64 {
	KEY++
	return KEY
}
