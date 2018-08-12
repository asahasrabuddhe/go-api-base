package request

type Request interface {
	Validate() error
}
