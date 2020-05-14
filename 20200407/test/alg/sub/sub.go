package sub

type AlgSub interface {
	Sub(lhs, rhs int) int
	Sub1(lhs, mid, rhs int) int
}

type Suber struct{}

func (Suber) Sub(lhs, rhs int) int {
	return lhs - rhs
}

func (Suber) Sub1(lhs, mid, rhs int) int {
	return lhs - rhs + mid
}
