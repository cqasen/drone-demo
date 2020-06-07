package enum

type IBoolEnum interface {
	IsTrue() bool
	ISFalse() bool
}

type BoolEnum struct {
	v int
}

const (
	False = iota
	True
)

func Bool(v int) BoolEnum {
	return BoolEnum{v: v}
}

func (b BoolEnum) IsTrue() bool {
	return b.v == True
}

func (b BoolEnum) ISFalse() bool {
	return b.v == False
}
