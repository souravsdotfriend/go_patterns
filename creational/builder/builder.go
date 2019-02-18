package builder

import "strconv"

type Color string

const (
	BLUE Color = "blue"
	RED        = "red"
)

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Car interface {
	Drive() string
	Stop() string
}

type CarBuilder interface {
	TopSpeed(float64) CarBuilder
	Paint(Color) CarBuilder
	Build() Car
}

type carBuilder struct {
	speedOption float64
	color       Color
}

func (cb *carBuilder) TopSpeed(speed float64) CarBuilder {
	cb.speedOption = speed
	return cb
}

func (cb *carBuilder) Paint(color Color) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		topSpeed: cb.speedOption,
		color:    cb.color,
	}
}

func NewBuilder() CarBuilder {
	return &carBuilder{}
}

type car struct {
	topSpeed float64
	color    Color
}

func (c *car) Drive() string {
	return "Driving at speed: " + strconv.FormatFloat(c.topSpeed, 'f', 6, 64)
}

func (c *car) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}