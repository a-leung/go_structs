package main

type Dice struct {
	Seed int
}

func (d *Dice) SetSeed(value int) {
	d.Seed = value
}

func (d *Dice) GetSeed() int {
	return d.Seed
}
