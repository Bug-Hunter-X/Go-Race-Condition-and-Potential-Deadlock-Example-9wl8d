```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int)

        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := 0; i < 5; i++ {
                        ch <- i
                }
                close(ch) // Close the channel after sending all values
        }()

        wg.Wait() // Wait for the goroutine to finish before reading from the channel

        for v := range ch {
                fmt.Println(v)
        }
}
```