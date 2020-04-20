package add

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
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
		if got := Add(tt.args.lhs, tt.args.rhs); got != tt.want {
			t.Errorf("%q. Suber.Sub1() = %v, want %v", tt.name, got, tt.want)
		}
		TLOG.Error("result: success")
	}
}

func TestAdd3(t *testing.T) {
	type args struct {
		lhs int
		mid int
		rhs int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "2-2-2", args: args{lhs: 2, mid: 2, rhs: 2}, want: 6},
	}
	for _, tt := range tests {
		Convey("将三数相加", t, func() {
			So(Add3(tt.args.lhs, tt.args.mid, tt.args.rhs), ShouldEqual, tt.want)
		})

		TLOG.Error("result: success")

	}
}
