package virtualdom

type NodeType int
type Attribute int

const (
	div NodeType = iota
	span
	button
)

const (
	id Attribute = iota
	class
)

// State manage state in view
type State struct{}

// Node represent component of dom
type Node struct {
	nodeType   NodeType
	attributes map[Attribute]string
	children   []interface{}
}
