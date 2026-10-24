package main

import (
	"math"
	"testing"
)

func TestCalcStats(t *testing.T) {
	s := calcStats([]int{2, 4, 4, 4, 5, 5, 7, 9})
	want := Stats{Count: 8, Sum: 40, Mean: 5, Max: 9, Min: 2, Even: 4, Odd: 4, Range: 7, Median: 4.5, Variance: 4, StdDev: 2}
	if s != want {
		t.Errorf("calcStats = %+v\nwant        %+v", s, want)
	}
}

func TestCalcStatsNegativeAndZero(t *testing.T) {
	s := calcStats([]int{-3, 0, -8, 7, -1})
	if s.Sum != -5 || s.Max != 7 || s.Min != -8 || s.Even != 2 || s.Odd != 3 || s.Median != -1 {
		t.Errorf("calcStats = %+v", s)
	}
	if math.Abs(s.Mean-(-1)) > 1e-12 {
		t.Errorf("Mean = %v; want -1", s.Mean)
	}
}

func TestCalcStatsSingle(t *testing.T) {
	s := calcStats([]int{6})
	if s.Sum != 6 || s.Mean != 6 || s.Max != 6 || s.Min != 6 || s.Even != 1 || s.Odd != 0 || s.Variance != 0 {
		t.Errorf("calcStats = %+v", s)
	}
}

func TestCalcStatsDoesNotModifyInput(t *testing.T) {
	nums := []int{3, 1, 2}
	calcStats(nums)
	if nums[0] != 3 || nums[1] != 1 || nums[2] != 2 {
		t.Errorf("входной срез изменён: %v", nums)
	}
}
