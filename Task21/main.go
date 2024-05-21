package main

type AnimalsAdapter interface {
	Reaction()
}

type CatAdapter struct {
	*Cat
}

func NewCatAdapter(cat *Cat) AnimalsAdapter {
	return &CatAdapter{
		cat,
	}
}

func (ca *CatAdapter) Reaction() {
	ca.Meow()
}

type DogAdapter struct {
	*Dog
}

func NewDogAdapter(dog *Dog) AnimalsAdapter {
	return &DogAdapter{
		dog,
	}
}
func (da *DogAdapter) Reaction() {
	da.Woof()
}

func main() {
	dAdapter := NewDogAdapter(&Dog{})
	dAdapter.Reaction()

	cAdapter := NewCatAdapter(&Cat{})
	cAdapter.Reaction()
}
