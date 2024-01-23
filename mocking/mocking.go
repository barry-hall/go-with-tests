package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type ConfigruableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigruableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigruableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
