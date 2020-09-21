package redisclient

// baseResult is the base result for redis commands
type baseResult interface {
	Err() error
}

// Result is the common result for redis commands
type Result interface {
	baseResult
	Val() interface{}
	Result() (interface{}, error)
}

// StringResult is the string result for redis commands
type StringResult interface {
	baseResult
	Val() string
	Result() (string, error)
}

// FloatResult is the float result for redis commands
type FloatResult interface {
	baseResult
	Val() float64
	Result() (float64, error)
}

// IntResult is the int result for redis commands
type IntResult interface {
	baseResult
	Val() int64
	Result() (int64, error)
}

// SliceResult is the slice result for redis commands
type SliceResult interface {
	baseResult
	Val() []interface{}
	Result() ([]interface{}, error)
}

// StatusResult is the status result for redis commands
type StatusResult interface {
	baseResult
	Val() string
	Result() (string, error)
}

// BoolResult the bool result for redis commands
type BoolResult interface {
	baseResult
	Val() bool
	Result() (bool, error)
}

// IntSliceResult is the int slice result for redis commands
type IntSliceResult interface {
	baseResult
	Val() []int64
	Result() ([]int64, error)
}

// StringSliceResult is the string slice result for redis commands
type StringSliceResult interface {
	baseResult
	Val() []string
	Result() ([]string, error)
}

// BoolSliceResult is the bool slice result for redis commands
type BoolSliceResult interface {
	baseResult
	Val() []bool
	Result() ([]bool, error)
}

// StringStringMapResult is the string string map result for redis commands
type StringStringMapResult interface {
	baseResult
	Val() map[string]string
	Result() (map[string]string, error)
}

// StringIntMapResult is the string int map result for redis commands
type StringIntMapResult interface {
	baseResult
	Val() map[string]int64
	Result() (map[string]int64, error)
}

// StringStructMapResult is the string struct map result for redis commands
type StringStructMapResult interface {
	baseResult
	Val() map[string]struct{}
	Result() (map[string]struct{}, error)
}
