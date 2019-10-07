package binary

type Path int

const (
	Left  Path = -1
	Right Path = 1
	Equal Path = 0
	Root  Path = 0
)

type Comparor interface {
	Compare(obj interface{}) Path
}
