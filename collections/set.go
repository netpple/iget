type void struct{}

var marking void

type Set struct {
  m map[string]void
}

func NewSet() *Set {
  return &Set{make(map[string]void)}
}

// Add, Remove, Contains, Len, Entries, String 
func (s *Set) Add(value String) {
  s.m[value] = marking
}

func (s *Set) Remove(value String) {
  delete(s.m, value)
}

func (s *Set) Contains(value String) {
  _, c := s.m[value]
  return c
}

func (s *Set) Len() int {
  return len(s.m)
}

func (s *Set) Entries() []string {
  entries := make([]string, 0, len(s.m))
  for k := range s.m {
    entries = append(entries, k)
  }

  return entries
}

func (s *Set) String() string {
  return fmt.Sprintf("[%s]", strings.Join(s.Entries(), ", "))
}
