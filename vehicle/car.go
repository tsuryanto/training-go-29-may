package vehicle

type Car struct {
	Tires [4]Tire
	Doors [2]Door
}

type Tire interface {
	CanRolling() bool
}

func (c Car) ChangeTire(index int, tire Tire) {
	if index >= 4 {
		return
	}
	c.Tires[index] = tire
}
