package merkledag

const (
	FILE = iota
	DIR
)

type Node interface {
	Size() uint64
	Name() string
	Type() int
}

type File interface {
	Node

	Bytes() []byte //存储文件内容
}

type Dir interface {
	Node

	It() DirIterator //迭代器
}

type DirIterator interface {
	Next() bool

	Node() Node
}
