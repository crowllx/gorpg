package components

import (
	"errors"
	"fmt"
)

type IStatus interface {
	_onHpChanged()
	_onDeath()
}
type Status struct {
	health       int
	maxHealth    int
	mana         int
	maxMana      int
	_onHpChanged func()
	_onDeath     func()
}

func NewStatus(hp, mp int, death func(), onHpChanged func()) *Status {
	return &Status{
		health:       hp,
		maxHealth:    hp,
		mana:         mp,
		maxMana:      mp,
		_onDeath:     death,
		_onHpChanged: onHpChanged,
	}
}
func (s *Status) Query(q string) (int, int, error) {
	switch q {
	case "health":
		return s.health, s.maxHealth, nil
	case "mana":
		return s.mana, s.maxMana, nil
	default:
		return 0, 0, errors.New(fmt.Sprintf("invalid query string: %s", q))
	}
}

func (s *Status) Modify(query string, value int) (int, error) {
	fmt.Println("help")
	switch query {
	case "health":
		s.health -= value
		if s._onHpChanged != nil {
			s._onHpChanged()
		}
		return s.health, nil
	case "mana":
		if value > s.mana {
			return s.mana, errors.New("not enough mana!")
		} else {
			s.mana -= value
		}
		return s.mana, nil
	default:
		return 0, errors.New(fmt.Sprintf("invalid query: %s", query))
	}
}
