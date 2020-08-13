package main

import (
	"log"

	"github.com/ev3go/ev3dev"
)

func main() {
	a, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left large motor on a: %v", err)
	}
	a.SetPosition(0)
	a.SetPosition(100)

	b, err := ev3dev.TachoMotorFor("ev3-ports:outB", "lego-ev3-m-motor")
	if err != nil {
		log.Fatalf("failed to find medium motor on b: %v", err)
	}

	b.SetPosition(0)
	b.SetPosition(100)

	c, err := ev3dev.TachoMotorFor("ev3-ports:outC", "lego-ev3-l-motor")
	if err != nil {
		log.Fatalf("failed to find left large motor on c: %v", err)
	}

	c.SetPosition(0)
	c.SetPosition(100)

}
