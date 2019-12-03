package Intcode

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zero", args{0, 0}, 0},
		{"Left hand nonzero", args{1, 0}, 1},
		{"Right hand nonzero", args{0, 1}, 1},
		{"Left hand negative", args{-1, 0}, -1},
		{"Right hand negative", args{0, -1}, -1},
		{"Both nonzero", args{5, 8}, 13},
		{"Both negative", args{-5, -8}, -13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zero", args{0, 0}, 0},
		{"Left hand nonzero", args{1, 0}, 0},
		{"Right hand nonzero", args{0, 1}, 0},
		{"Left hand identity", args{1, 7}, 7},
		{"Right hand identity", args{5, 1}, 5},
		{"Left hand negative", args{-2, 7}, -14},
		{"Right hand negative", args{5, -2}, -10},
		{"Both nonzero", args{5, 7}, 35},
		{"Both negative", args{-5, -7}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeAdd(t *testing.T) {
	type args struct {
		leftAddress   int
		rightAddress  int
		resultAddress int
		program       []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Add 1 to first", args{0, 0, 0, []int{1}}, []int{2}},
		{"Double second", args{1, 1, 1, []int{0, 10}}, []int{0, 20}},
		{"Add first two, put in third", args{0, 1, 2, []int{7, 9, 0}}, []int{7, 9, 16}},
		{"Overwrite left", args{0, 1, 0, []int{7, 9, 0}}, []int{16, 9, 0}},
		{"Overwrite right", args{0, 1, 1, []int{7, 9, 0}}, []int{7, 16, 0}},
		{"Zero out third left negative", args{0, 1, 2, []int{-9, 9, 1000}}, []int{-9, 9, 0}},
		{"Zero out third right negative", args{0, 1, 2, []int{9, -9, 1000}}, []int{9, -9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeAdd(tt.args.leftAddress, tt.args.rightAddress, tt.args.resultAddress, tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeMultiply(t *testing.T) {
	type args struct {
		leftAddress   int
		rightAddress  int
		resultAddress int
		program       []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Square first - identity", args{0, 0, 0, []int{1}}, []int{1}},
		{"Square square - nonidentity", args{0, 0, 0, []int{2}}, []int{4}},
		{"Double second", args{0, 1, 1, []int{2, 10}}, []int{2, 20}},
		{"Triple second", args{0, 1, 1, []int{3, 10}}, []int{3, 30}},
		{"Multiply first two, put in third", args{0, 1, 2, []int{7, 9, 0}}, []int{7, 9, 63}},
		{"Overwrite left", args{0, 1, 0, []int{7, 9, 0}}, []int{63, 9, 0}},
		{"Overwrite right", args{0, 1, 1, []int{7, 9, 0}}, []int{7, 63, 0}},
		{"Zero out third left zero", args{0, 1, 2, []int{0, 9, 1000}}, []int{0, 9, 0}},
		{"Zero out third right zero", args{0, 1, 2, []int{9, 0, 1000}}, []int{9, 0, 0}},
		{"Zero out third both zero", args{0, 1, 2, []int{0, 0, 1000}}, []int{0, 0, 0}},
		{"Negate left", args{0, 1, 0, []int{10, -1, 1000}}, []int{-10, -1, 1000}},
		{"Negate right", args{0, 1, 1, []int{-1, -10, 1000}}, []int{-1, 10, 1000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeMultiply(tt.args.leftAddress, tt.args.rightAddress, tt.args.resultAddress, tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeStep(t *testing.T) {
	type args struct {
		programCounter int
		program        []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		want1   bool
		wantErr bool
	}{
		{"Halt", args{0, []int{99, 42, 42, 42}}, []int{99, 42, 42, 42}, true, false},
		{"Halt No extra", args{0, []int{99}}, []int{99}, true, false},
		{"Halt On second step", args{1, []int{1, 0, 0, 0, 99, 42, 42, 42}}, []int{1, 0, 0, 0, 99, 42, 42, 42}, true, false},
		{"Halt On second step no extra", args{1, []int{1, 0, 0, 0, 99}}, []int{1, 0, 0, 0, 99}, true, false},
		{"Add on first step", args{0, []int{1, 0, 0, 1, 99}}, []int{1, 2, 0, 1, 99}, false, false},
		{"Add on second step", args{1, []int{3, 0, 2, 0, 1, 5, 2, 1, 99}}, []int{3, 7, 2, 0, 1, 5, 2, 1, 99}, false, false},
		{"Multiply on first step", args{0, []int{2, 0, 0, 2, 99}}, []int{2, 0, 4, 2, 99}, false, false},
		{"Multiply on second step", args{1, []int{1, 0, 1, 0, 2, 4, 5, 2, 99}}, []int{1, 0, 8, 0, 2, 4, 5, 2, 99}, false, false},
		{"Unknown opcode", args{0, []int{98, 0, 1, 0, 2, 4, 5, 2, 99}}, []int{98, 0, 1, 0, 2, 4, 5, 2, 99}, true, true},
		{"Program counter too high", args{3, []int{1, 0, 1, 0, 2, 4, 5, 2, 99}}, []int{1, 0, 1, 0, 2, 4, 5, 2, 99}, true, true},
		{"Unexpected instruction length", args{0, []int{1, 0, 1}}, []int{1, 0, 1}, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := executeStep(tt.args.programCounter, tt.args.program)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeStep() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("executeStep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	type args struct {
		program []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Halt", args{[]int{99, 42, 42, 42}}, []int{99, 42, 42, 42}, false},
		{"Halt No extra", args{[]int{99}}, []int{99}, false},
		{"Add then Halt On second step", args{[]int{1, 0, 0, 0, 99, 42, 42, 42}}, []int{2, 0, 0, 0, 99, 42, 42, 42}, false},
		{"Add then Halt On second step no extra", args{[]int{1, 0, 0, 0, 99}}, []int{2, 0, 0, 0, 99}, false},
		{"Add on first step", args{[]int{1, 0, 0, 1, 99}}, []int{1, 2, 0, 1, 99}, false},
		{"Add on second step", args{[]int{1, 0, 2, 0, 1, 5, 2, 1, 99}}, []int{3, 7, 2, 0, 1, 5, 2, 1, 99}, false},
		{"Multiply on first step", args{[]int{2, 0, 0, 2, 99}}, []int{2, 0, 4, 2, 99}, false},
		{"Multiply on second step", args{[]int{1, 0, 1, 1, 2, 4, 5, 2, 99}}, []int{1, 1, 8, 1, 2, 4, 5, 2, 99}, false},
		{"Unknown opcode", args{[]int{98, 0, 1, 0, 2, 4, 5, 2, 99}}, []int{98, 0, 1, 0, 2, 4, 5, 2, 99}, true},
		{"Unexpected instruction length", args{[]int{1, 0, 1}}, []int{1, 0, 1}, true},
		{"First example from site", args{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, false},
		{"Second example from site", args{[]int{1, 0, 0, 0, 99}}, []int{2, 0, 0, 0, 99}, false},
		{"Third example from site", args{[]int{2, 3, 0, 3, 99}}, []int{2, 3, 0, 6, 99}, false},
		{"Fourth example from site", args{[]int{2, 4, 4, 5, 99, 0}}, []int{2, 4, 4, 5, 99, 9801}, false},
		{"Fifth example from site", args{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, false},
		{"First problem", args{[]int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 5, 19, 23, 2, 6, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 5, 35, 39, 2, 6, 39, 43, 2, 6, 43, 47, 1, 5, 47, 51, 2, 9, 51, 55, 1, 5, 55, 59, 1, 10, 59, 63, 1, 63, 6, 67, 1, 9, 67, 71, 1, 71, 6, 75, 1, 75, 13, 79, 2, 79, 13, 83, 2, 9, 83, 87, 1, 87, 5, 91, 1, 9, 91, 95, 2, 10, 95, 99, 1, 5, 99, 103, 1, 103, 9, 107, 1, 13, 107, 111, 2, 111, 10, 115, 1, 115, 5, 119, 2, 13, 119, 123, 1, 9, 123, 127, 1, 5, 127, 131, 2, 131, 6, 135, 1, 135, 5, 139, 1, 139, 6, 143, 1, 143, 6, 147, 1, 2, 147, 151, 1, 151, 5, 0, 99, 2, 14, 0, 0}}, []int{4484226, 12, 2, 2, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 24, 1, 5, 19, 25, 2, 6, 23, 50, 1, 27, 5, 51, 2, 9, 31, 153, 1, 5, 35, 154, 2, 6, 39, 308, 2, 6, 43, 616, 1, 5, 47, 617, 2, 9, 51, 1851, 1, 5, 55, 1852, 1, 10, 59, 1856, 1, 63, 6, 1858, 1, 9, 67, 1861, 1, 71, 6, 1863, 1, 75, 13, 1868, 2, 79, 13, 9340, 2, 9, 83, 28020, 1, 87, 5, 28021, 1, 9, 91, 28024, 2, 10, 95, 112096, 1, 5, 99, 112097, 1, 103, 9, 112100, 1, 13, 107, 112105, 2, 111, 10, 448420, 1, 115, 5, 448421, 2, 13, 119, 2242105, 1, 9, 123, 2242108, 1, 5, 127, 2242109, 2, 131, 6, 4484218, 1, 135, 5, 4484219, 1, 139, 6, 4484221, 1, 143, 6, 4484223, 1, 2, 147, 4484225, 1, 151, 5, 0, 99, 2, 14, 0, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Execute(tt.args.program)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
