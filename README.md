# Go Race Condition and Potential Deadlock

This repository demonstrates a subtle race condition and potential deadlock in a Go program that uses goroutines and channels.  The program attempts to send values to a channel and then receive them, but the order of operations leads to a potential issue.

## The Bug

The `bug.go` file contains the erroneous code.  The main goroutine waits for a worker goroutine to finish and close the channel.  However, it then proceeds to read from the closed channel, which in normal scenarios won't throw an immediate error, but may lead to unexpected behavior or even deadlock.