package runequest

// Modifier is a bonus or penalty to a Character element
type Modifier struct {
	Object
	Value   int
	Set     bool
	Modify  bool
	Base    bool
	Subject string
}

// Object is an interface tfor any numerical ability
type Object interface {
	ModifyValue(int)
	SetBase(int)
	SetValue(int)
}

// ModifyValue implements the Object interface to increases or decreases a value
func (m Modifier) ModifyValue() {
	m.Object.ModifyValue(m.Value)
}

// SetBase implements the Object interface to set a base value
func (m Modifier) SetBase() {
	m.Object.SetBase(m.Value)
}

// SetValue implements the Object interface to set a base value
func (m Modifier) SetValue() {
	m.Object.SetValue(m.Value)
}

// ModifyValue updates the Skill value
func (s *Skill) ModifyValue(v int) {
	s.Value += v
}

// SetBase sets a skills Base value to an integer
func (s *Skill) SetBase(i int) {
	s.Base = i
}

// SetValue sets a skills Base value to an integer
func (s *Skill) SetValue(i int) {
	s.Value = i
}

// ModifyValue updates the Statistic value
func (s *Statistic) ModifyValue(v int) {
	s.Value += v
}

// SetBase sets a stats Base value to an integer
func (s *Statistic) SetBase(i int) {
	s.Base = i
}

// SetValue sets a Stats Base value to an integer
func (s *Statistic) SetValue(i int) {
	s.Value = i
}

// ModifyValue updates the Ability value
func (a *Ability) ModifyValue(v int) {
	a.Value += v
}

// SetBase sets an abilities base balue to an int
func (a *Ability) SetBase(i int) {
	a.Base = i
}

// SetValue sets an Ability Base value to an integer
func (a *Ability) SetValue(i int) {
	a.Value = i
}
