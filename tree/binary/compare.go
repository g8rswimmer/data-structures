package binary

// Path is the type of node path
type Path int

const (
	// Left is the path of the node, which is less than the parent
	Left Path = -1
	// Right is the path of the node, which is greater than the parent
	Right Path = 1
	// Equal is the node is the same as the parent
	Equal Path = 0
	// Root is the node is the same as the root
	Root Path = 0
)

// Comparor is the interface used to compare a node to another
type Comparor interface {
	Compare(obj interface{}) Path
}
