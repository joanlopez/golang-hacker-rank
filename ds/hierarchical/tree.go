package hierarchical

type Tree interface {
	Contains(pattern interface{}) bool
	Insert(data interface{})
	Delete(data interface{})
	CountLeafs() int
	InOrder() string
	PreOrder() string
	PostOrder() string
	LevelOrder() string
}