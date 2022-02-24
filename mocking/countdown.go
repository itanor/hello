package main

import (
  "io"
  "fmt"
  "os"
  "time"
)

type ConfigurableSleeper struct {
  duration time.Duration
  sleep func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
  c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
  Sleep()
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
  time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    sleeper.Sleep()
    fmt.Fprint(out, i)
    fmt.Fprint(out, "\n")
  }

  sleeper.Sleep()
  fmt.Fprint(out, finalWord)
}

func main() {
  sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
  Countdown(os.Stdout, sleeper)
}

