package main

import (
 "fmt"
 "sync"
)

type Turnstile struct {
 count int
 mu    sync.Mutex
}

func (t *Turnstile) PassThrough() {
 t.mu.Lock()
 defer t.mu.Unlock()
 t.count++
}

func (t *Turnstile) GetCount() int {
 t.mu.Lock()
 defer t.mu.Unlock()
 return t.count
}

func main() {
 turnstile := Turnstile{}

 var wg sync.WaitGroup
 wg.Add(100)

 for i := 0; i < 100; i++ {
  go func() {
   defer wg.Done()
   turnstile.PassThrough()
  }()
 }

 wg.Wait()

 fmt.Printf("Total people passed through the turnstile: %d\n", turnstile.GetCount())
}