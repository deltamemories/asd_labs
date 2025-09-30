package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	// arrange
	var tests = []struct {
		name    string
		s       string
		want    float64
		wantErr bool
	}{
		{"normal_1", "(21+33)*3.2-4+(5*2-6)*2=", 176.8, false},
		{"normal_2", "-4(2*4+5-1)2(2+2)=", -384, false},
		{"two_dots_in_float", "1.2.3*3=", 0.0, true},
		{"only_operators", "+-*=", 0.0, true},
		{"space_between_digits", "1 2+3=", 0.0, true},
		{"incorrect_float", ".*5=", 0.0, true},
		{"empty_string", "", 0.0, true},
		{"only_equal_sign", "=", 0.0, false},
		{"random_text_with_equal_sign", "wiesfknds5+*-8=", 0.0, true},
		{"two_equal_signs", "1+3==", 0.0, true},
		{"without_equal_sign", "5+9-3*2", 0.0, true},
		{"divide_by_zero_1", "55/0=", 0.0, true},
		{"divide_by_zero_2", "4/(10-2*5)=", 0.0, true},
		{"unary_minus_sign", "-3+5=", 2, false},
		{"operator_near_of_unary_minus_sign", "2*-3=", -6, false},
		{"multiply_by_brackets_1", "(5+2)(6-2)=", 28, false},
		{"multiply_by_brackets_2", "5(2+3)=", 25, false},
		{"nested_brackets", "((3+9)*2+7)*4=", 124, false},
		{"incorrect_brackets_1", ")(2+7)*3=", 0.0, true},
		{"incorrect_brackets_2", "(2+7*3=", 0.0, true},
	}

	// act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act, err := Calc(tt.s)

			// assert
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error: expected wantErr=%v, got err=%v", tt.wantErr, err)
				return
			}
			if !tt.wantErr && act != tt.want {
				t.Errorf("Calc() unexpected result: expected: %f, got: %f", tt.want, act)
			}

		})
	}

}
