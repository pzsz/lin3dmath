package lin3dmath

import "sort"

type DistSortVector2f struct {
	Points []Vector2f
	DistPoint Vector2f
}

func (s *DistSortVector2f) Sort() {
	sort.Sort(s)
}

func (s *DistSortVector2f) Len() int {
	return len(s.Points)
}

func (s *DistSortVector2f) Less(i, j int) bool {
	d1 := s.Points[i].Sub(s.DistPoint)
	d1l := d1.Len2()
	d2 := s.Points[i].Sub(s.DistPoint)
	d2l := d2.Len2()

	return d1l < d2l;
}

func (s *DistSortVector2f) Swap(i, j int) {
	t := s.Points[i]
	s.Points[i] = s.Points[j]
	s.Points[j] = t
}