package ev3api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/ev3go/ev3dev"
)

// New returns a Handler interface to be used.
func New() *Handler {
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
	motor, err := ev3dev.TachoMotorFor(string(legoMotorName), string(legoMotorType))
	if err != nil {
		return err
	}
	h.motorMap[name] = motor
	return nil
}

// MoveMotor is a REST API that allows the client to move a registered motor by a
// name and speed.
func (h *Handler) MoveMotor(w http.ResponseWriter, r *http.Request) {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	motorName := values.Get("name")
	if motorName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("query param 'name' is required"))
		return
	}

	speedString := values.Get("speed")
	if motorName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("query param 'speed' is required"))
		return

	}

	speed, err := strconv.Atoi(speedString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("param 'speed' has to be an integer"))
		return
	}

	motor, ok := h.motorMap[motorName]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("motor not registered"))
		return

	}

	motor.SetSpeedSetpoint(speed).Command("run-forever")
	if motor.Err() != nil {
		motor.Command("stop")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// StopMotor stops the motor.
func (h *Handler) StopMotor(w http.ResponseWriter, r *http.Request) {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	motorName := values.Get("name")
	if motorName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("query param 'name' is required"))
		return
	}

	motor, ok := h.motorMap[motorName]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("motor not registered"))
		return
	}

	motor.Command("stop")
	w.WriteHeader(http.StatusOK)
	return
}
