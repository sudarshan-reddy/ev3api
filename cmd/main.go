package main

import (
	"log"
	"time"

	"github.com/ev3go/ev3dev"
)

func main() {
	a, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left large motor on a: %v", err)
	}

	a.SetSpeedSetpoint(-10).Command("run-forever")
	time.Sleep(2 * time.Second)
	a.Command("stop")

	b, err := ev3dev.TachoMotorFor("ev3-ports:outB", "lego-ev3-m-motor")
	if err != nil {
		log.Fatalf("failed to find medium motor on b: %v", err)
	}

	b.SetSpeedSetpoint(-10).Command("run-forever")
	time.Sleep(2 * time.Second)
	b.Command("stop")

	c, err := ev3dev.TachoMotorFor("ev3-ports:outC", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left large motor on c: %v", err)
	}

	c.SetSpeedSetpoint(-10).Command("run-forever")
	time.Sleep(2 * time.Second)
	c.Command("stop")

}
