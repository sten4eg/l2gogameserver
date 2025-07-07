package models

type Account struct {
	Char     []*Character
	CharSlot int32
	Login    string
}

func (a *Account) AddCharacter(character *Character) {
	a.Char = append(a.Char, character)
}
func (a *Account) Len(character *Character) int {
	return len(a.Char)
}

func (a *Account) GetLogin() string {
	return a.Login
}

func (a *Account) GetChar() *Character {
	return &Character{}
}
