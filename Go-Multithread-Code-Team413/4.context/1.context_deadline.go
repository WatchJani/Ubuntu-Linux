// Sample program to show how to use the WithDeadline function
// of the Context package.
package main

import (
    "context"
    "fmt"
    "time"
)

type data struct {
    UserID string
}

func main() {
    duration := time.Now().Add(3 * time.Second)

    // Create a context that is both manually cancellable and will signal
    // a cancel at the specified duration.
    ctx, cancel := context.WithDeadline(context.Background(), duration)
    defer cancel()

    // Create a channel to received a signal that work is done.
    ch := make(chan data, 1)
    // Ask the goroutine to do some work for us.
    go func() {

        // Simulate work.
        time.Sleep(50 * time.Millisecond)

        // Report the work is done.
        ch <- data{"123"}
    }()

    // Wait for the work to finish. If it takes too long move on.
    select {
    case d := <-ch:
        fmt.Println("work complete", d)
    case <-ctx.Done():
        fmt.Println("work cancelled")
    }
}
