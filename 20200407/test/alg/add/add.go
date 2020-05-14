package add

import (
	_ "github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"
)

var TLOG = rogger.GetLogger("")

func Add(lhs, rhs int) int {
	TLOG.Error("result:", lhs+rhs)
	return lhs + rhs
}

func Add3(lhs, mid, rhs int) int {
	TLOG.Error("result:", lhs+mid+rhs)
	return lhs + mid + rhs
}
