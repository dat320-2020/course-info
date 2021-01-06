# Exam December 2020: DAT320 Operating Systems and Systems Programming

## Description of the Scheduling Simulator

The exam assignments below builds on a *simulated* scheduling framework.
The simulator consists of separate packages for different entities, such as the CPU, Job, and the System itself.
The system receives a schedule of job entries via its `Run` method, and is responsible for scheduling these jobs according to their arrival times.
The system can be given any type of scheduler as long as it implements the `scheduler.Scheduler` interface.

```golang
type Scheduler interface {
	Add(*job.Job)
	Tick(time.Duration) int
}
```

The `job.Job` type may be augmented with additional fields and methods, if necessary to implement a specific scheduler.
It should not be necessary for this exam.
The `size` parameter of the `job.New()` function can be ignored in this exam.

The `Add` method takes a job and adds it to the scheduler's queue.

- The `Tick` method is where the main scheduling logic should be implemented.
- The `Tick` method is called by the simulator every clock tick of the system, defined via the `config.TickDuration` constant.
- The `Tick` method receives the current *system time* as input.
  That is, the number of ticks since the start of the simulation.
- The `Tick` method should tick all CPUs that has a running job, unless the scheduler specifies otherwise.
- The `Tick` method returns the number of jobs that finished in the current tick.
  That is, `Tick` may return 0 or 1 for a single-CPU system, and in a multi-CPU system it may return a value between 0 and the number of CPUs whose job finished in the current clock tick.
- The `Tick` method *may assign* new jobs to the different CPUs before returning.

The `System` for driving the scheduling will only use the `Scheduler` interface (the `Add` and `Tick` methods) to interact with your scheduler.

## Recommendations

Study the provided source code to understand how it works.

It is recommended to break up your code into separate functions if a function grows beyond 30 lines.

You should not need to change any of the provided code, except for the code in the various packages, as described in the Task section below.
However, you may find it useful to add debug statements also in the provided code.
This is perfectly fine.

## Troubleshooting

If you copy files between packages, remember to update the package name on the first line in the file, so that it matches its enclosing folder name.
Otherwise, you will get compile errors from VSCode and the Go compiler.

Open VSCode from the directory that contains the `go.mod` file.
This allows the Go plugin to compile your code correctly and not get confused by other projects you may have in your folder structure.

If you run into a problem that is taking a lot of time, remember you can often find good explanations online (use Google).

However, if you are not able to complete one of the tasks to your own satisfaction, you can create a file named `manual-review.md` in the relevant task folder.
In this file, please explain the problem you are having, and describe with words and/or pseudo code your idea for a solution.
Please limit your text explanation to 200 words, and at most 50 lines of pseudo code.

## Task: Implement a Multi-Queue Multiprocessor Scheduler

In this assignment, you are to implement several variants of the Multi-Queue Multiprocessor Scheduler (MQMS) in the packages listed below.
For reference, Chapter 10 in the textbook describes the basic concepts of a MQMS.

While the different tasks have been given different package names/folders, you are expected to reuse/copy code from one package to the other where that is appropriate.

Note that the provided `mq_state.go` file provides a few helper functions to manage the state of a CPU and its job queue.
Using these functions can save you some time.

To make testing easier, Tasks 1-3 are expected to behave deterministically.
That is, given the same input schedule for repeated test runs, your solution should give the same output.

1. In package `mqmsbasic` implement a MQMS that schedules the jobs on different CPUs, where each job runs to completion before another job runs on the same CPU.
   When a new job arrives, it should be scheduled on the CPU whose queue is the shortest.
   If multiple CPU queues have the same length, pick the queue of the lowest CPU number.
   That is, CPU0 should be picked before CPU1 and so on.

   Note: Job switching happens immediately after a job has finished, without any **Idle** time.
   That is, the `Tick` method should assign new jobs before returning, if more jobs exists.

   An example trace is shown below for `TestMQMSBasicSystem`.

2. In package `mqmsrr` implement a Round Robin MQMS that schedule jobs on different CPUs.
   Each job stays on the same CPU until it is finished, but switches according to the time `quantum` parameter, between multiple jobs scheduled on the same CPU.

   Note: Round Robin job switching happens only at synchronized time intervals on all CPUs, and is determined by the system time and the `quantum` parameter.
   This means that if a job finishes before the time quantum expires, the CPU running that job should become **Idle**.

   An example trace is shown below for `TestMQMSRoundRobinSystem`.

3. In package `mqmsws` implement a simplified Work Stealing MQMS that schedules the jobs Round Robin on different CPUs, but where an idle CPU may steal a job from another CPU's queue.

   Note: Work stealing should only occur when a CPU's queue is empty, which is different from the description in Chapter 10.
   Further, an idle CPU should not steal a job already running on another CPU, only those that are queued.
   Otherwise, job switching follows the same behavior as for Round Robin.

   An example trace is shown below for `TestMQMSWorkStealingSystem`.

4. Implement **only one** of the following tasks.

   a) In package `mqms` implement the same as in Task 2 above.

   b) In package `mqms` implement the same as in Task 3 above.

   However, now the `Tick` method should start a separate goroutine for each CPU (for the same system time) and safely update the data structures.
   That is, there should be no data races or deadlocks.

   Hint: To prevent that multiple goroutines for *different system time values* occur concurrently, you may wish to use a `sync.WaitGroup` to limit concurrency to the same system time (`Tick`).
   Further, for this task, you may want to make adjustments to the `mq_state.go` file to ensure synchronized access.

   Note that when running the tests for these assignments, the output may be non-deterministic (different) each time you run the test.

## Test Scenarios

You can test your solution with these commands:

```sh
go test -v -run TestMQMSBasicSystem
go test -v -run TestMQMSRoundRobinSystem
go test -v -run TestMQMSWorkStealingSystem
go test -v -run TestMQMSConcurrentSystem
```

These tests does not actually test anything, but will print the output generated by the different schedulers, so that you can view the actual execution of the given schedule.
The expected schedules are shown below.

### Output From TestMQMSBasicSystem

```text
=== RUN   TestMQMSBasicSystem
Tick  CPU0     CPU1    CPU2     CPU3
0s    A(10ms)  B(3ms)  C(10ms)  D(3ms)
1ms   A(9ms)   B(2ms)  C(9ms)   D(2ms)
2ms   A(8ms)   B(1ms)  C(8ms)   D(1ms)
3ms   A(7ms)   F(5ms)  C(7ms)   H(8ms)
4ms   A(6ms)   F(4ms)  C(6ms)   H(7ms)
5ms   A(5ms)   F(3ms)  C(5ms)   H(6ms)
6ms   A(4ms)   F(2ms)  C(4ms)   H(5ms)
7ms   A(3ms)   F(1ms)  C(3ms)   H(4ms)
8ms   A(2ms)   I(5ms)  C(2ms)   H(3ms)
9ms   A(1ms)   I(4ms)  C(1ms)   H(2ms)
10ms  E(10ms)  I(3ms)  G(10ms)  H(1ms)
11ms  E(9ms)   I(2ms)  G(9ms)   Idle
12ms  E(8ms)   I(1ms)  G(8ms)   Idle
13ms  E(7ms)   J(8ms)  G(7ms)   Idle
14ms  E(6ms)   J(7ms)  G(6ms)   Idle
15ms  E(5ms)   J(6ms)  G(5ms)   Idle
16ms  E(4ms)   J(5ms)  G(4ms)   Idle
17ms  E(3ms)   J(4ms)  G(3ms)   Idle
18ms  E(2ms)   J(3ms)  G(2ms)   Idle
19ms  E(1ms)   J(2ms)  G(1ms)   Idle
20ms  Idle     J(1ms)  Idle     Idle
21ms  Idle     Idle    Idle     Idle
```

### Output From TestMQMSRoundRobinSystem

```text
=== RUN   TestMQMSRoundRobinSystem
Tick  CPU0     CPU1    CPU2     CPU3
0s    A(10ms)  B(3ms)  C(10ms)  D(3ms)
1ms   A(9ms)   B(2ms)  C(9ms)   D(2ms)
2ms   A(8ms)   B(1ms)  C(8ms)   D(1ms)
3ms   A(7ms)   Idle    C(7ms)   Idle
4ms   A(6ms)   Idle    C(6ms)   Idle
5ms   E(10ms)  F(5ms)  G(10ms)  H(8ms)
6ms   E(9ms)   F(4ms)  G(9ms)   H(7ms)
7ms   E(8ms)   F(3ms)  G(8ms)   H(6ms)
8ms   E(7ms)   F(2ms)  G(7ms)   H(5ms)
9ms   E(6ms)   F(1ms)  G(6ms)   H(4ms)
10ms  A(5ms)   I(5ms)  C(5ms)   J(8ms)
11ms  A(4ms)   I(4ms)  C(4ms)   J(7ms)
12ms  A(3ms)   I(3ms)  C(3ms)   J(6ms)
13ms  A(2ms)   I(2ms)  C(2ms)   J(5ms)
14ms  A(1ms)   I(1ms)  C(1ms)   J(4ms)
15ms  E(5ms)   Idle    G(5ms)   H(3ms)
16ms  E(4ms)   Idle    G(4ms)   H(2ms)
17ms  E(3ms)   Idle    G(3ms)   H(1ms)
18ms  E(2ms)   Idle    G(2ms)   Idle
19ms  E(1ms)   Idle    G(1ms)   Idle
20ms  Idle     Idle    Idle     J(3ms)
21ms  Idle     Idle    Idle     J(2ms)
22ms  Idle     Idle    Idle     J(1ms)
23ms  Idle     Idle    Idle     Idle
```

### Output From TestMQMSWorkStealingSystem

```text
=== RUN   TestMQMSWorkStealingSystem
15ms: CPU1 stealing job from CPU2: G(5ms)
15ms: CPU2 stealing job from CPU3: H(3ms)
Tick  CPU0     CPU1    CPU2     CPU3
0s    A(10ms)  B(3ms)  C(10ms)  D(3ms)
1ms   A(9ms)   B(2ms)  C(9ms)   D(2ms)
2ms   A(8ms)   B(1ms)  C(8ms)   D(1ms)
3ms   A(7ms)   Idle    C(7ms)   Idle
4ms   A(6ms)   Idle    C(6ms)   Idle
5ms   E(10ms)  F(5ms)  G(10ms)  H(8ms)
6ms   E(9ms)   F(4ms)  G(9ms)   H(7ms)
7ms   E(8ms)   F(3ms)  G(8ms)   H(6ms)
8ms   E(7ms)   F(2ms)  G(7ms)   H(5ms)
9ms   E(6ms)   F(1ms)  G(6ms)   H(4ms)
10ms  A(5ms)   I(5ms)  C(5ms)   J(8ms)
11ms  A(4ms)   I(4ms)  C(4ms)   J(7ms)
12ms  A(3ms)   I(3ms)  C(3ms)   J(6ms)
13ms  A(2ms)   I(2ms)  C(2ms)   J(5ms)
14ms  A(1ms)   I(1ms)  C(1ms)   J(4ms)
15ms  E(5ms)   G(5ms)  H(3ms)   J(3ms)
16ms  E(4ms)   G(4ms)  H(2ms)   J(2ms)
17ms  E(3ms)   G(3ms)  H(1ms)   J(1ms)
18ms  E(2ms)   G(2ms)  Idle     Idle
19ms  E(1ms)   G(1ms)  Idle     Idle
20ms  Idle     Idle    Idle     Idle
```
