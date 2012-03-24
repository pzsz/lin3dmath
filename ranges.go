package lin3dmath

type IntRange struct {
	Start int
	End   int
}

func (s IntRange) Size() int {
	return s.End - s.Start
}

func (s IntRange) Overlap(other IntRange) bool {
	// other    S     E
	// s     S     E         
	if s.End > other.Start && s.Start <= other.Start {
		return true
	}

	if other.End >  s.Start && other.Start <= s.Start {
		return true
	}
	return false
}

func (s IntRange) OverlapAndStick(other IntRange) bool {
	if s.End >= other.Start && s.Start <= other.Start {
		return true
	}

	if other.End >= s.Start && other.Start <= s.Start {
		return true
	}
	return false
}


func (s IntRange) InsideInt(pos int) bool {
	return s.Start <= pos && pos < s.End
}

func (s *IntRange) MoreInt(pos int) bool {
	return pos < s.Start
}

// Ranges have to overlap. this is OR operation on sets
func (s IntRange) Add(other IntRange) (result IntRange) {
	if s.Start < other.Start {
		result.Start = s.Start
	} else {
		result.Start = other.Start
	}

	if s.End > other.End {
		result.End = s.End
	} else {
		result.End = other.End
	}
	return
}

// Ranges have to overlap. this is AND operation on sets
func (s IntRange) And(other IntRange) (result IntRange) {
	if s.Start >= other.Start {
		result.Start = s.Start
	} else {
		result.Start = other.Start
	}

	if s.End <= other.End {
		result.End = s.End
	} else {
		result.End = other.End
	}
	return
}


func (s *IntRange) Less(other IntRange) bool {
	return s.End < other.Start
}

type IntRanges struct {
	ranges []IntRange
}

func NewIntRanges() *IntRanges {
	return &IntRanges{}
}

func (s *IntRanges) Get() []IntRange {
	return s.ranges
}


func (s *IntRanges) GetUnion(rng *IntRanges) (result []IntRange) {
	for _, r1 := range s.ranges {
		for _, r2 := range rng.ranges {		
			if r1.Overlap(r2) {
				result = append(result, r1.And(r2))
			}
		}
	}
	return 
}

/*
func (s *IntRanges) GetUnion(rng *IntRanges) (result []IntRange) {
	left_pos := 0
	right_pos := 0

	if len(rng.ranges) == 0 || len(s.ranges) == 0 {
		return nil
	}

	grow_left = s.ranges[0].Start < rng.ranges[0].Start

	for ;left_pos < len(s.ranges) && right_pos < len(rng.ranges); {
		if s.ranges[left_pos].Overlap(rng.ranges[right_pos]) {
			result = append(result, s.ranges[left_pos].And(rng.ranges[right_pos]))
		} 
		
		if grow_left {
			left_pos += 1
		} else {
			right_pos += 1
		}
		
	}	
}*/

// Get inverted ranges in [min,max) range
func (s *IntRanges) GetInvert(min, max int) (result []IntRange) {
	prev_end := 0
	added_last_step := false
	for _, rang := range s.ranges {
		if rang.Start >= max {
			added_last_step = true
			break
		}

		if min < rang.Start {
			end := rang.Start
			if end > max {
				end = max
			}
			result = append(result, IntRange{prev_end, end})
		}
		prev_end = rang.End
	}

	if !added_last_step && prev_end < max {
		result = append(result, IntRange{prev_end, max})
	}
	
	return result
}

func (s *IntRanges) AddInt(start, end int) {	
	s.Add(IntRange{start, end})
}

func (s *IntRanges) Add(r IntRange) {	
	for i, rang := range s.ranges {
		if r.Less(rang) {
			s.ranges = append(s.ranges, r)
			copy(s.ranges[i+1:], s.ranges[i:])
			s.ranges[i] = r
			return
		}
		if r.OverlapAndStick(rang) {
			s.ranges[i] = rang.Add(r)
			return				
		}
	}

	s.ranges = append(s.ranges, r)
}
