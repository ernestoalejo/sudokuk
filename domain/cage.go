package domain

type Cage struct {
	Sum   int8
	Cells []*Cell
}

type Cell struct {
	Row, Col int8
}
