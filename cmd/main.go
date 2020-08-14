package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sudarshan-reddy/ev3api"
)

func main() {
	armExample := ev3api.New()
	armExample.RegisterLegoMotor("base", ev3api.OutC, ev3api.Large)
	armExample.RegisterLegoMotor("arm", ev3api.OutB, ev3api.Large)
	armExample.RegisterLegoMotor("claw", ev3api.OutA, ev3api.Medium)

	r := chi.NewRouter()

	r.Get("/move", armExample.MoveMotor)
	r.Get("/stop", armExample.StopMotor)

	http.ListenAndServe("0.0.0.0:6000", r)
}
