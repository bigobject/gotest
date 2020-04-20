package sub_test

import (
	"test/test/alg/sub"
	"test/test/alg/sub/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSuber_Sub(t *testing.T) {
	type args struct {
		lhs int
		rhs int
	}
	tests := []struct {
		name string
		s    sub.Suber
		args args
		want int
	}{
		{name: "1-1", s: sub.Suber{}, args: args{lhs: 1, rhs: 1}, want: 0},
		{name: "1-2", s: sub.Suber{}, args: args{lhs: 1, rhs: 2}, want: -1},
	}
	for _, tt := range tests {
		s := sub.Suber{}
		if got := s.Sub(tt.args.lhs, tt.args.rhs); got != tt.want {
			t.Errorf("%q. Suber.Sub() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSuber_Sub2(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSub := mock.NewMockAlgSub(mockCtl)
	mockSub.EXPECT().Sub(1, 1).Return(0)
	mockSub.EXPECT().Sub(1, 2).Return(-1)
	mockSub.EXPECT().Sub1(2, 2, 2).Return(2)

	type args struct {
		lhs int
		rhs int
	}
	tests := []struct {
		name string
		s    sub.AlgSub
		args args
		want int
	}{
		{name: "1-1", s: mockSub, args: args{lhs: 1, rhs: 1}, want: 0},
		{name: "1-2", s: mockSub, args: args{lhs: 1, rhs: 2}, want: -1},
	}
	for _, tt := range tests {
		if got := tt.s.Sub(tt.args.lhs, tt.args.rhs); got != tt.want {
			t.Errorf("%q. Suber.Sub() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSuber_Sub1(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockSub := mock.NewMockAlgSub(mockCtl)
	mockSub.EXPECT().Sub1(2, 2, 2).Return(2)
	type args struct {
		lhs int
		mid int
		rhs int
	}
	tests := []struct {
		name string
		s    sub.AlgSub
		args args
		want int
	}{
		{name: "2-2-2", s: mockSub, args: args{lhs: 2, mid: 2, rhs: 2}, want: 2},
	}
	for _, tt := range tests {
		if got := tt.s.Sub1(tt.args.lhs, tt.args.mid, tt.args.rhs); got != tt.want {
			t.Errorf("%q. Suber.Sub1() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
