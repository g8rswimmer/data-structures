package binary

type mockIntCompare struct {
	data int
}

func (m *mockIntCompare) Compare(obj interface{}) Path {
	intCompare, ok := obj.(*mockIntCompare)
	if ok == false {
		return Path(-100)
	}
	switch {
	case intCompare.data < m.data:
		return Left
	case intCompare.data > m.data:
		return Right
	default:
		return Equal
	}
}
