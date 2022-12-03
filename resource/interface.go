package resource

type Interface interface {
	Entity() interface{}
	Collection() []interface{}
}
