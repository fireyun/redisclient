package redisclient

// BaseResult is the base result for redis commands
type BaseResult interface {
	Err() error
}

// Result is the common result for redis commands
type Result interface {
	BaseResult
	Val() interface{}
	Result() (interface{}, error)
}

// StringResult is the string result for redis commands
type StringResult interface {
	BaseResult
	Val() string
	Result() (string, error)
}

// FloatResult is the float result for redis commands
type FloatResult interface {
	BaseResult
	Val() float64
	Result() (float64, error)
}

// IntResult is the int result for redis commands
type IntResult interface {
	BaseResult
	Val() int64
	Result() (int64, error)
}

// SliceResult is the slice result for redis commands
type SliceResult interface {
	BaseResult
	Val() []interface{}
	Result() ([]interface{}, error)
}

// StatusResult is the status result for redis commands
type StatusResult interface {
	BaseResult
	Val() string
	Result() (string, error)
}

// BoolResult the bool result for redis commands
type BoolResult interface {
	BaseResult
	Val() bool
	Result() (bool, error)
}

// IntSliceResult is the int slice result for redis commands
type IntSliceResult interface {
	BaseResult
	Val() []int64
	Result() ([]int64, error)
}

// StringSliceResult is the string slice result for redis commands
type StringSliceResult interface {
	BaseResult
	Val() []string
	Result() ([]string, error)
}

// BoolSliceResult is the bool slice result for redis commands
type BoolSliceResult interface {
	BaseResult
	Val() []bool
	Result() ([]bool, error)
}

// StringStringMapResult is the string string map result for redis commands
type StringStringMapResult interface {
	BaseResult
	Val() map[string]string
	Result() (map[string]string, error)
}

// StringIntMapResult is the string int map result for redis commands
type StringIntMapResult interface {
	BaseResult
	Val() map[string]int64
	Result() (map[string]int64, error)
}

// StringStructMapResult is the string struct map result for redis commands
type StringStructMapResult interface {
	BaseResult
	Val() map[string]struct{}
	Result() (map[string]struct{}, error)
}

// PipelineResult is the pipeline result for redis pipeline command
type PipelineResult interface {
	Elems() []Result
}
