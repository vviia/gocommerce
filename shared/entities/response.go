package entities

type BaseResposne[T interface{}] struct {
	CommonResult
	Message string `json:"message"`
	Data    T
}
