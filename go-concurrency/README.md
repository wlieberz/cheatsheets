# About

This mini-project was me experiementing with simple concurrency in Go.

It may serve as a cheatsheet to myself in the future. Maybe it will help you 
as well. :)

# Walkthrough

### Task Scripts

We start with three bash scripts:

1. task-1.sh
2. task-2.sh
3. task-3.sh

Each is meant to simulate a task which will take some time to complete.
Each script does: `sleep` 10 seconds and then echo `$scriptName complete`

### Bash Blocking

Next, we have `all-tasks.sh` which stores the three task scripts in an array
and loops over the array, executing each script sending stderr to stdout and
uses tee to write output to the console as well as to the `output-bash` 
directory, with one file per script, named `$scriptName_output.json`. 
Nevermind that the output isn't actually json. :)

Executing `./all-tasks.sh` runs each task script sequentially, with each task
blocking further execution until it completes. If you run it with `time` you'll
see that it takes 30 seconds to complete, as you would expect. 

While `time ./all-tasks.sh` is running, in a second terminal 
you can run:

```bash
ps aux | head -1 && ps aux | grep bash | grep task
```

Run the above `ps` command a few times and you'll see that there is never more 
than one task-script running at a time.

### Go Blocking

Next, we have `all-tasks_blocking.go` which is an implementation of 
`all-tasks.sh` but in Go rather than Bash.

In `all-tasks_blocking.go` we make a slice of strings with a capacity of 3 
called `scriptsToRun` and then assign the three task-scripts into the slice.

We loop over `scriptsToRun` and call the `execScript` function on each script.

The `execScript` function (pretty much) replicates the behavior of 
`all-tasks.sh` in that it prints stdout and stderr to both the console and 
writes the output (one out-file per script) to an output directory. 
This time we write to just `output` rather than `output-bash`, just so we can
save the output in both cases and verify the output is the same.

If you want, you can run it prefixed with `time` and also run the `ps` test
command and verify that the behavior is the same as `all-tasks.sh`.

### Go Concurrent

Finally, what we have been building up to, the concurrent version, in Go:
`all-tasks.go`.

First let's see what it does and test that it works as expected.

```bash
# Build it:
go build all-tasks.go

# Run it:
time ./all-tasks

# Or, build and run without saving the resultant binary:
time go run all-tasks.go
```

You should see output like:

```bash
./task-2.sh complete
./task-1.sh complete
./task-3.sh complete

real    0m10.012s
user    0m0.017s
sys     0m0.007s
```
While it is running we can run the `ps` test command:

```bash
ps aux | head -1 && ps aux | grep bash | grep task
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
wlieberz  109614  0.0  0.0   9496  3268 pts/16   S+   16:30   0:00 /bin/bash -c ./task-3.sh | tee ./output/task-3.sh_output.json
wlieberz  109615  0.0  0.0   9496  3248 pts/16   S+   16:30   0:00 /bin/bash -c ./task-1.sh | tee ./output/task-1.sh_output.json
wlieberz  109617  0.0  0.0   9496  3316 pts/16   S+   16:30   0:00 /bin/bash -c ./task-2.sh | tee ./output/task-2.sh_output.json
wlieberz  109618  0.0  0.0   9492  3320 pts/16   S+   16:30   0:00 /bin/bash ./task-3.sh
wlieberz  109620  0.0  0.0   9492  3196 pts/16   S+   16:30   0:00 /bin/bash ./task-1.sh
wlieberz  109622  0.0  0.0   9492  3316 pts/16   S+   16:30   0:00 /bin/bash ./task-2.sh
```

From the above output notice that three instances of bash were created all at
once and, as such, it only took the duration of one script to finish the work
of all three: 10 seconds instead of 30 seconds.

We can also validate that the output is correct: we didn't get any threading
issues where the output of one thread wrote to a wrong location. 

In other words:

```bash
cd output

cat task-1.sh_output.json 
./task-1.sh complete

cat task-2.sh_output.json 
./task-2.sh complete

cat task-3.sh_output.json 
./task-3.sh complete
```

### Quick codewalk - all-tasks.go

Interestingly `all-tasks.go` differs very little from `all-tasks_blocking.go`.

We import the `sync` package to get access WaitGroups.

We initialize a new WaitGroup as `wg`.

Within each loop over `scriptsToRun` we increment the `wg` counter by one.

We need to copy the `script` var each time through the loop so the go routine 
closure gets it's own copy of the var. This would be akin to cloning the var
in Rust. As I understand it, this is done to prevent the anonymous function
or closure from sending the same stale variable to the inner `execScript` call.

This is explained better and in more detail here:
https://go.dev/doc/faq#closures_and_goroutines

Finally, outside the for loop, at the bottom of `main` we call `wg.Wait()`.
Without this, Go's parent process (running main) would see that it has reached
the end of `main` and halt execution, almost certainly before the go routine 
processes that were spawned for the task-scripts would have a chance to 
finish executing. By waiting for the WaitGroup, we ensure that all three
processes finish before halting main.