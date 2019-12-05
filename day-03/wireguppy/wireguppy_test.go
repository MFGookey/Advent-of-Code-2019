package wireguppy

import (
	"reflect"
	"testing"
)

func TestCompassDirection_String(t *testing.T) {
	tests := []struct {
		name      string
		direction CompassDirection
		want      string
	}{
		{"Up", Up, "Up"},
		{"Right", Right, "Right"},
		{"Down", Down, "Down"},
		{"Left", Left, "Left"},
		{"Unknown low", -1, "Unknown"},
		{"Unknown high", 4, "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.direction.String(); got != tt.want {
				t.Errorf("CompassDirection.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseDirection(t *testing.T) {
	type args struct {
		direction string
	}
	tests := []struct {
		name    string
		args    args
		want    CompassDirection
		wantErr bool
	}{
		{"Up lower", args{"u"}, Up, false},
		{"Up upper", args{"U"}, Up, false},
		{"Right lower", args{"r"}, Right, false},
		{"Right upper", args{"R"}, Right, false},
		{"Down lower", args{"d"}, Down, false},
		{"Down upper", args{"D"}, Down, false},
		{"Left lower", args{"l"}, Left, false},
		{"Left upper", args{"L"}, Left, false},
		{"Nonexistant", args{"Q"}, Up, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDirection(tt.args.direction)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDirection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		instruction string
	}
	tests := []struct {
		name    string
		args    args
		want    WireInstruction
		wantErr bool
	}{
		{"Up 0", args{"U0"}, WireInstruction{Up, 0}, false},
		{"Right 10", args{"R10"}, WireInstruction{Right, 10}, false},
		{"Left 5", args{"L5"}, WireInstruction{Left, 5}, false},
		{"Down 7", args{"D7"}, WireInstruction{Down, 7}, false},
		{"Bad Direction", args{"Q7"}, WireInstruction{Up, 0}, true},
		{"Bad Distance", args{"UQ23"}, WireInstruction{Up, 0}, true},
		{"Empty wire", args{""}, WireInstruction{Up, 0}, true},
		{"Bad direction and distance", args{"QQ"}, WireInstruction{Up, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.instruction)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseWire(t *testing.T) {
	type args struct {
		wireArray []string
	}
	tests := []struct {
		name    string
		args    args
		want    []WireInstruction
		wantErr bool
	}{
		{"Single Step Wire", args{[]string{"U5"}}, []WireInstruction{{Up, 5}}, false},
		{"Multi Step Wire", args{[]string{"U5", "D4", "L23", "R56"}}, []WireInstruction{{Up, 5}, {Down, 4}, {Left, 23}, {Right, 56}}, false},
		{"Empty Wire", args{[]string{}}, nil, true},
		{"Bad direction wire", args{[]string{"Q5"}}, nil, true},
		{"Bad distance wire", args{[]string{"UQ5"}}, nil, true},
		{"Bad direction and distance wire", args{[]string{"QQ"}}, nil, true},
		{"Second step bad direction wire", args{[]string{"U1 Q5"}}, nil, true},
		{"Second step bad distance wire", args{[]string{"U1 UQ5"}}, nil, true},
		{"Second step bad direction and distance wire", args{[]string{"U1 QQ"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseWire(tt.args.wireArray)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseWire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseWire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseWires(t *testing.T) {
	type args struct {
		wireArrays [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]WireInstruction
		wantErr bool
	}{
		// Single wire tests
		{"Single Single Step Wire", args{[][]string{{"U5"}}}, [][]WireInstruction{{{Up, 5}}}, false},
		{"Single Multi Step Wire", args{[][]string{{"U5", "D4", "L23", "R56"}}}, [][]WireInstruction{{{Up, 5}, {Down, 4}, {Left, 23}, {Right, 56}}}, false},
		{"Single Empty Wire", args{[][]string{{}}}, nil, true},
		{"Single Bad direction wire", args{[][]string{{"Q5"}}}, nil, true},
		{"Single Bad distance wire", args{[][]string{{"UQ5"}}}, nil, true},
		{"Single Bad direction and distance wire", args{[][]string{{"QQ"}}}, nil, true},
		{"Single Second step bad direction wire", args{[][]string{{"U1 Q5"}}}, nil, true},
		{"Single Second step bad distance wire", args{[][]string{{"U1 UQ5"}}}, nil, true},
		{"Single Second step bad direction and distance wire", args{[][]string{{"U1 QQ"}}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseWires(tt.args.wireArrays)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseWires() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseWires() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Positive number", args{1}, 1},
		{"Negative number", args{-1}, 1},
		{"Zero", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateManhattanFromOrigin(t *testing.T) {
	type args struct {
		point Coordinate
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"X only", args{Coordinate{5, 0}}, 5},
		{"Negative X only", args{Coordinate{-7, 0}}, 7},
		{"Y only", args{Coordinate{0, 98}}, 98},
		{"Negative Y only", args{Coordinate{0, -985}}, 985},
		{"Positive X, Positive Y", args{Coordinate{10, 7}}, 17},
		{"Positive X, Negative Y", args{Coordinate{9, -4}}, 13},
		{"Negative X, Positive Y", args{Coordinate{-2, 3}}, 5},
		{"Negative X, Negative Y", args{Coordinate{-9, -40}}, 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateManhattanFromOrigin(tt.args.point); got != tt.want {
				t.Errorf("CalculateManhattanFromOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_Add(t *testing.T) {
	type args struct {
		offset Coordinate
	}
	tests := []struct {
		name  string
		start Coordinate
		args  args
		want  Coordinate
	}{
		{"Up 2", Coordinate{5, 7}, args{Coordinate{2, 0}}, Coordinate{7, 7}},
		{"Right 5", Coordinate{5, 2}, args{Coordinate{0, 5}}, Coordinate{5, 7}},
		{"Down 9", Coordinate{5, 7}, args{Coordinate{0, -9}}, Coordinate{5, -2}},
		{"Left 4", Coordinate{5, 2}, args{Coordinate{-4, 0}}, Coordinate{1, 2}},
		{"Knight's move 1,2", Coordinate{5, 2}, args{Coordinate{1, 2}}, Coordinate{6, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.start.Add(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coordinate.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWireCoordinates(t *testing.T) {
	type args struct {
		wire []WireInstruction
	}
	tests := []struct {
		name string
		args args
		want []Coordinate
	}{
		{"Up only", args{[]WireInstruction{WireInstruction{Up, 3}}}, []Coordinate{Coordinate{0, 1}, Coordinate{0, 2}, Coordinate{0, 3}}},
		{"Down only", args{[]WireInstruction{WireInstruction{Down, 3}}}, []Coordinate{Coordinate{0, -1}, Coordinate{0, -2}, Coordinate{0, -3}}},
		{"Right only", args{[]WireInstruction{WireInstruction{Right, 3}}}, []Coordinate{Coordinate{1, 0}, Coordinate{2, 0}, Coordinate{3, 0}}},
		{"Left only", args{[]WireInstruction{WireInstruction{Left, 3}}}, []Coordinate{Coordinate{-1, 0}, Coordinate{-2, 0}, Coordinate{-3, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WireCoordinates(tt.args.wire); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WireCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
