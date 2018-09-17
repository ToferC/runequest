package runequest

// Modifier is a bonus or penalty to a Character element
type Modifier struct {
	Object
	Value int
}

// Object is an interface tfor any numerical ability
type Object interface {
	ModifyValue(int)
	SetBase(int)
}

// ModifyValue implements the Object interface to increases or decreases a value
func (m Modifier) ModifyValue() {
	m.Object.ModifyValue(m.Value)
}

// ModifyValue updates the Skill value
func (s *Skill) ModifyValue(v int) {
	s.Value += v
}

// SetBase sets a skills Base value to an integer
func (s *Skill) SetBase(i int) {
	s.Base = i
}

// ModifyValue updates the Statistic value
func (s *Statistic) ModifyValue(v int) {
	s.Value += v
}

// SetBase sets a stats Base value to an integer
func (s *Statistic) SetBase(i int) {
	s.Base = i
}

// ModifyValue updates the Ability value
func (a *Ability) ModifyValue(v int) {
	a.Value += v
}

// SetBase sets an abilities base balue to an int
func (a *Ability) SetBase(i int) {
	a.Base = i
}
