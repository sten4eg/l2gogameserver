package clientStates

type State byte

const (
	// Connected Дефолтный стейт
	Connected State = 1
	// Authed Клиент прошел авторизацию, но к нему еще не привязан персонаж
	Authed State = 2
	// Joining Клиент выбрал персонажа, но еще не присоединился к серверу
	Joining State = 3
	// InGame Клиент выбрал персонажа и находится в игре
	InGame State = 4
)
