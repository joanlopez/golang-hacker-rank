package hierarchical

type Tree interface {
	Contains(pattern interface{}) bool
	Insert(data interface{})
	Leafs() int
}