func main() {
    runtime.GOMAXPROCS(1)
    wg := sync.WaitGroup{}
    wg.Add(20)
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Print(" A:", i)
            wg.Done()
        }()
    }

    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Print(" B:", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
