package tree

type BPlusNode interface {
	Setup(int order) void

	// Type of node. 0 empty, 1 internal, 2 leaf
	NodeType() int

	// Delete an item (composite)
	Delete() BPlusNode

	// Search for an item
	Search() item

	// Return tree depth
	Depth() int

	// Return width of this B+ tree
	Span() int
}
