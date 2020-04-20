package multi_test

import (
	"test/test/alg/multi"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMulti(t *testing.T) {
	type args struct {
		lhs int
		rhs int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "2-2-2", args: args{lhs: 2, rhs: 2}, want: 4},
	}
	for _, tt := range tests {
		Convey("将两数相加", t, func() {
			So(multi.Multi(tt.args.lhs, tt.args.rhs), ShouldEqual, tt.want)
		})

	}
}
