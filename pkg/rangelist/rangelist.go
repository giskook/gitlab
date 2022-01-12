package rangelist

type RangeList interface {
	Add(r Range) error
	Remove(r Range) error
	Print()
}
