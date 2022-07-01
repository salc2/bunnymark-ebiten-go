package bunny

type Bunny struct {
	PositionX float64
	PositionY float64
	SpeedX    float64
	SpeedY    float64
	Theme     int
}

func (b *Bunny) Update(maxX, maxY float64, timeDelta, gravity float64) error {
	b.PositionX += (b.SpeedX * timeDelta)
	b.PositionY += (b.SpeedY * timeDelta)
	b.SpeedY += gravity

	if b.PositionX > maxX {
		b.SpeedX *= -1
		b.PositionX = maxX
	}
	if b.PositionX < 0 {
		b.SpeedX *= -1
		b.PositionX = 0
	}

	if b.PositionY > maxY {
		b.SpeedY *= -0.85
		b.PositionY = maxY

	}
	if b.PositionY < 0 {
		b.SpeedY = 0
		b.PositionY = 0
	}
	return nil
}
