package linear

type List interface {
	Empty() bool
	Length() int
	String() string
	Contains(search interface{}) bool
	At(search interface{}) int
	Occurrences(search interface{}) int
	Prepend(data interface{})
	Append(data interface{})
	Shift() interface{}
	Pop() interface{}
	InsertAt(data interface{}, pos int) error
	GetAt(pos int) (interface{}, error)
	DeleteAt(pos int) (interface{}, error)
}