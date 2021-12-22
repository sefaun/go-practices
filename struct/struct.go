package struct_practice

type General struct {
	id      int
	correct bool
}

type Sefa struct {
	id      int
	name    string
	general General
}

type AllRepo interface {
	Fetch(deger int) Sefa
}

func (sefa *Sefa) Fetch(deger int) *Sefa {
	sefa.id = deger
	sefa.name = "Sefa"
	sefa.general.id = deger + 1
	sefa.general.correct = false

	return sefa
}
