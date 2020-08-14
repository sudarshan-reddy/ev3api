package ev3api

import (
	"net/http"

	"github.com/ev3go/ev3dev"
)

// NewHandler returns a Handler interface to be used.
func NewHandler() *Handler {
	return &Handler{
		motorMap: make(map[string]*ev3dev.TachoMotor),
	}
}

// Handler represents the struct on which the ev3apis are written on.
type Handler struct {
	motorMap map[string]*ev3dev.TachoMotor
}

// LegoMotorName is a type indicating the name of the lego motor as
// supported by ev3go.
type LegoMotorName string

const (
	// OutA ...
	OutA LegoMotorName = "ev3-ports:outA"

	// OutB ...
	OutB LegoMotorName = "ev3-ports:outB"

	// OutC ...
	OutC LegoMotorName = "ev3-ports:outC"

	// OutD ...
	OutD LegoMotorName = "ev3-ports:outD"
)

// LegoMotorType is a type indicating the size primarily of the lego
// motor.
type LegoMotorType string

const (
	// Medium ...
	Medium LegoMotorType = "lego-ev3-m-motor"

	// Large ...
	Large LegoMotorType = "lego-ev3-l-motor"
)

// RegisterLegoMotor Lets you register a motor to the ev3api. Duplicate names will be over-written.
func (h *Handler) RegisterLegoMotor(name string, legoMotorName LegoMotorName,
	legoMotorType LegoMotorType) error {
	motor, err := ev3dev.TachoMotorFor(legoMotorName, legoMotorType)
	if err != nil {
		return err
	}
	h.motorMap[name] = motor
	return nil
}

// MoveMotor is a REST API that allows the client to move a registered motor by a
// name and speed.
func (h *Handler) MoveMotor(r *http.Request, w http.ResponseWriter) {
	r.URL.Values()
	a.SetSpeedSetpoint(-10).Command("run-forever")
}
