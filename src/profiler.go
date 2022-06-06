package main

import (
	"time"
	"log"
)

type Profiler struct {
	Data map[string]time.Time
}

func create_profiler() Profiler {
	p := Profiler{Data:make(map[string]time.Time)}
	return p
}

func (p *Profiler) Start(name string) {
	p.Data[name] = time.Now()
}

func (p *Profiler) End(name string) {
	elapsed := time.Since(p.Data[name])
    log.Printf("\033[32m%s took %s\033[97m", name, elapsed)
}