package components

type Direction struct {
	X, Y int
}
type Movement struct {
	X, Y     float64
	Dir      Direction
	Cardinal int
}
