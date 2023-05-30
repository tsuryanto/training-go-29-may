package vehicle

type WoodTire struct {
	Ukuran int
}

func (k WoodTire) CanRolling() bool {
	return true
}

type IronTire struct {
	Ukuran int
}

func (k IronTire) CanRolling() bool {
	return true
}

type RubberTire struct {
	Ukuran int
}

func (k RubberTire) CanRolling() bool {
	return true
}
