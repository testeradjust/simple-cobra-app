package internal

type EzJsonCity struct {
	Name       string         `json:"name"`
	Population int            `json:"population"`
	People     []EzJsonPerson `json:"people"`
}

func NewEzJsonCity(name string, pop int) EzJsonCity {
	return EzJsonCity{
		Name:       name,
		Population: pop,
		People:     []EzJsonPerson{},
	}
}
