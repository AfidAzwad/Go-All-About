# Concurrency vs Asynchrony in Go

You might have noticed terms like **concurrent** and **asynchronous** being used when talking about programming, particularly with goroutines in Go. While they are related concepts, they are not the same. Let's clarify the distinction between them.

## 1. Concurrency

- **Definition**: Concurrency refers to the ability of a system to make progress on multiple tasks by working on them in overlapping time periods, though not necessarily at the same instant.
  
- **Example in Go**: In Go, goroutines enable **concurrent execution**. Goroutines are lightweight threads that can run at the same time (or nearly the same time) depending on the system's ability to schedule them. The Go runtime schedules multiple goroutines on different cores, allowing them to run concurrently.

    ```go
    go addition(5, 6)
    fmt.Println("Hello World!")
    ```

   In this case, both the `addition` goroutine and the `main` function are running concurrently. They might overlap in time, but not necessarily run at exactly the same time.

## 2. Asynchrony

- **Definition**: Asynchrony means that tasks are started and allowed to run independently, and the program does not wait for them to complete before moving on. The key aspect of asynchrony is that tasks start and run **without blocking** the main flow of the program.

- **Example in Go**: When you start a goroutine, it behaves **asynchronously** in the sense that it doesn’t block the execution of the rest of the program. The `main` function keeps executing, while the goroutine is performing its task. The main program doesn't wait for the goroutine to complete unless explicitly told to (using synchronization tools like `WaitGroup`).

    ```go
    go addition(5, 6) // This is asynchronous, does not block main.
    fmt.Println("Hello World!")
    ```

## Key Differences

- **Asynchronous** execution is about **not waiting** for a task to finish before moving on.
- **Concurrent** execution is about **dealing with multiple tasks** at the same time (they may run independently or interact with each other).

### In Go:

- **Goroutines** are **concurrent** because they allow multiple tasks to run potentially in overlapping time periods.
- They are also **asynchronous** because once started, the program doesn’t wait for them to finish unless you explicitly synchronize (like using `WaitGroup`).

### Example Scenario

If you're performing two tasks:
- **Concurrent** means you are able to work on both tasks at the same time.
- **Asynchronous** means one task might be kicked off and you're free to do something else while it's happening, without waiting for it to finish.

## Conclusion

In Go, **goroutines are both concurrent and asynchronous**:
- **Concurrent** because multiple goroutines can execute simultaneously or switch between them.
- **Asynchronous** because they run independently without blocking other parts of the program.
