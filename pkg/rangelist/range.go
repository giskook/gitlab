package rangelist

import "gitlab/pkg/rangelist/base"

type Range [2]int

func (r Range) Valid() bool {
	return r[0] < r[1]
}

func (r Range) Equal(dst Range) bool{
	return r.LowerBoundary() == dst.LowerBoundary() &&
		r.UpperBoundary() == dst.UpperBoundary()
}

func (r Range) LowerBoundary() int {
	return r[0]
}

func (r Range) UpperBoundary() int {
	return r[1]
}

// @Insect 包括包含关系
func (r Range) Insect(dst Range) bool {
	return (r.LowerBoundary() <= dst.LowerBoundary() &&
		r.UpperBoundary() >= dst.LowerBoundary()) ||
		(r.LowerBoundary() <= dst.UpperBoundary() &&
			r.UpperBoundary() >= dst.UpperBoundary()) ||
		(dst.LowerBoundary() <= r.LowerBoundary() &&
			dst.UpperBoundary() >= r.LowerBoundary()) ||
		(dst.LowerBoundary() <= r.UpperBoundary() &&
			dst.UpperBoundary() >= r.UpperBoundary())
}

// @ Left r 在 dst的左侧
func (r Range) Left(dst Range) bool {
	return r.UpperBoundary() < dst.LowerBoundary()
}

// @ Right r 在 dst的右侧
func (r Range) Right(dst Range) bool {
	return r.LowerBoundary() > dst.UpperBoundary()
}

func (r Range) Contain(dst Range) bool {
	return r.LowerBoundary() <= dst.LowerBoundary() &&
		r.UpperBoundary() >= dst.UpperBoundary()
}

func (r Range) Union(dst Range) Range {
	return Range{base.Min(r.LowerBoundary(), dst.LowerBoundary()),
		base.Max(r.UpperBoundary(), dst.UpperBoundary())}
}

func (r Range) Inter(dst Range) Range {
	return Range{base.Max(r.LowerBoundary(), dst.LowerBoundary()),
		base.Min(r.UpperBoundary(), dst.UpperBoundary())}
}

func (r Range) LeftDiff(dst Range) Range {
	return Range{r.LowerBoundary(), base.Min(r.UpperBoundary(), dst.LowerBoundary())}
}

func (r Range) RightDiff(dst Range) Range {
	return Range{base.Max(r.LowerBoundary(), dst.UpperBoundary()), r.UpperBoundary()}
}
