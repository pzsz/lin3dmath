package lin3dmath_test

import (
	. "pzs/lin3dmath"
	"testing"
)

func testOverlap(t *testing.T, r1,r2 IntRange, should bool) {
	if (should) {
		if !r1.Overlap(r2) {
			t.Errorf("IntRanges did not overlaped!")
		}
		if !r2.Overlap(r1) {
			t.Errorf("IntRanges did not overlaped! (in reverse)")
		}
	} else {
		if r1.Overlap(r2) {
			t.Errorf("IntRanges that should not, overlaped!")
		}
		if r2.Overlap(r1) {
			t.Errorf("IntRanges that should not, overlaped! (in reverse)")
		}		
	}
}


func testOverlapAndStick(t *testing.T, r1,r2 IntRange, should bool) {
	if (should) {
		if !r1.OverlapAndStick(r2) {
			t.Errorf("IntRanges did not overlaped!")
		}
		if !r2.OverlapAndStick(r1) {
			t.Errorf("IntRanges did not overlaped! (in reverse)")
		}
	} else {
		if r1.OverlapAndStick(r2) {
			t.Errorf("IntRanges that should not, overlaped!")
		}
		if r2.OverlapAndStick(r1) {
			t.Errorf("IntRanges that should not, overlaped! (in reverse)")
		}		
	}
}


func TestIntRangeOverlap(t *testing.T) {
	r1 := IntRange{5,10}
	r1_stick_after := IntRange{10, 15}
	r1_nstick_after := IntRange{11, 15}

	r1_stick_before := IntRange{0, 5}
	r1_nstick_before := IntRange{0, 4}

	r1_over_after := IntRange{7, 15}
	r1_over_before := IntRange{1, 7}
	r1_over := IntRange{1, 15}


	testOverlap(t, r1, r1_stick_after, false)
	testOverlap(t, r1, r1_stick_before, false)
	testOverlap(t, r1, r1_nstick_after, false)
	testOverlap(t, r1, r1_nstick_before, false)

	testOverlapAndStick(t, r1, r1_stick_after, true)
	testOverlapAndStick(t, r1, r1_stick_before, true)
	testOverlapAndStick(t, r1, r1_nstick_after, false)
	testOverlapAndStick(t, r1, r1_nstick_before, false)

	testOverlap(t, r1, r1_over_after, true)
	testOverlap(t, r1, r1_over_before, true)
	testOverlap(t, r1, r1_over, true)

	testOverlapAndStick(t, r1, r1_over_after, true)
	testOverlapAndStick(t, r1, r1_over_before, true)
	testOverlapAndStick(t, r1, r1_over, true)
}

func expectRanges(t *testing.T, res []IntRange, expected []IntRange) {
	if len(res) != len(expected) {
		t.Errorf("Expected %v elements, got %v", len(expected), len(res))
		t.Errorf("%v", res)
		return
	}
	for x:=0; x < len(res); x++ {
		if res[x].Start != expected[x].Start || 
			res[x].End != expected[x].End {
			t.Errorf("Expected %v, got %v", expected[x], res[x])
			return
		}
	}
}

func TestAdd(t *testing.T) {
	ranges := NewIntRanges()

	// 5-11. 15-17. 20-25, 29-35
	ranges.Add(IntRange{10, 14})
	ranges.Add(IntRange{11, 15})
	ranges.Add(IntRange{15, 20})
	ranges.Add(IntRange{05, 10})
	ranges.Add(IntRange{20, 25})

	rng := ranges.Get()
	expectRanges(t, rng, []IntRange{IntRange{5,25}})
}

func TestUnion(t *testing.T) {
	ranges1 := NewIntRanges()
	ranges2 := NewIntRanges()

	// 5-11. 15-17. 20-25, 29-35
	ranges1.Add(IntRange{5, 10})
	ranges1.Add(IntRange{14, 20})
	ranges1.Add(IntRange{21, 22})
	ranges1.Add(IntRange{23, 25})

	ranges2.Add(IntRange{0, 17})
	ranges2.Add(IntRange{19, 30})

	rng := ranges1.GetUnion(ranges2)
	expectRanges(t, rng, []IntRange{
		IntRange{5,10},
		IntRange{14,17},
		IntRange{19,20},
		IntRange{21,22},
		IntRange{23,25}})
}


func TestInvert(t *testing.T) {
	ranges := NewIntRanges()

	// 5-11. 15-17. 20-25, 29-35
	ranges.Add(IntRange{5,10})
	ranges.Add(IntRange{7,11})
	ranges.Add(IntRange{15, 17})
	ranges.Add(IntRange{20, 25})
	ranges.Add(IntRange{29, 35})

	inv := ranges.GetInvert(0,35)
	expectRanges(t, inv, []IntRange{
		IntRange{0,5},
		IntRange{11,15},
		IntRange{17,20},
		IntRange{25,29}})

	inv = ranges.GetInvert(5,17)
	expectRanges(t, inv, []IntRange{
		IntRange{11,15}})

	inv = ranges.GetInvert(11,40)
	expectRanges(t, inv, []IntRange{
			IntRange{11,15},
			IntRange{17,20},
			IntRange{25,29},
			IntRange{35,40}})

}