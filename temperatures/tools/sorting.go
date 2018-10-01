package tools

type tree struct {
	value int
	left, right *tree
}

// Sort sorts an array of values in place.
func Sort(arr []int) []int {
	var root *tree
	for _, item := range arr {
		root = add(root, item)
	}
	return appendValues(arr[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if t.value < value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}