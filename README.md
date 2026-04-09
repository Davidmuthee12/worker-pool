# Go Worker Pool File Processor

A small concurrency project built to practice worker pools, goroutines, and channels.

This program creates a fixed number of workers that process a list of file names concurrently.

## Concepts Practiced

- Goroutines
- Channels
- Worker pools
- Buffered channels
- Concurrent job distribution

---

## How It Works

1. A list of file names is created.
2. A `jobs` channel is used to send file names to workers.
3. Three worker goroutines are started.
4. Each worker:
   - waits for a file name from the `jobs` channel
   - simulates processing for 2 seconds
   - sends the result into the `results` channel
5. The main goroutine reads all results and prints them.

---

## Project Structure

```text
.
├── main.go
└── README.md
```

---

## Code Overview

### Worker Function

```
func worker(id int, jobs <-chan string, results chan<- string)
```

Parameters:

- `id` → the worker number
- `jobs` → receives file names to process
- `results` → sends processed file messages back

Example worker flow:

```text
Worker 1 receives "image1.jpg"
→ processes it
→ sends "image1.jpg processed by worker 1"
```

---

## Input Files

```go id="h8z9xv"
files := []string{
	"image1.jpg",
	"image2.png",
	"document.pdf",
	"video.mp4",
	"notes.txt",
}
```

---

## Running the Program

Run:

```bash id="0xtkdy"
go run .
```

---

## Expected Behavior

Since there are 3 workers and 5 files:

- The first 3 files are picked up immediately by the 3 workers.
- When a worker finishes, it automatically takes the next available file.
- Files may finish in a different order than they started.

The exact order may change each time because goroutines run concurrently.

Example output:

```text id="3v4n5f"
Worker 3 started processing image1.jpg
Worker 2 started processing document.pdf
Worker 1 started processing image2.png

Worker 1 finished processing image2.png
Worker 1 started processing video.mp4

Worker 2 finished processing document.pdf
Worker 3 finished processing image1.jpg

image2.png processed by worker 1
document.pdf processed by worker 2
image1.jpg processed by worker 3

Worker 2 started processing notes.txt

Worker 1 finished processing video.mp4
video.mp4 processed by worker 1

Worker 2 finished processing notes.txt
notes.txt processed by worker 2
```

---

## Why the Order Changes

Even though the files are sent in this order:

```text id="gjjlwm"
image1.jpg
image2.png
document.pdf
video.mp4
notes.txt
```

the workers run independently, so whichever worker finishes first grabs the next file.

For example:

- Worker 1 may finish first and start `video.mp4`
- Worker 2 may then pick up `notes.txt`

This is what makes a worker pool efficient.

---

## Buffered Channels Used

```go id="7g01qq"
jobs := make(chan string, len(files))
results := make(chan string, len(files))
```

These are buffered channels.

Why?

- `jobs` can hold all file names before workers receive them.
- `results` can store processed results until the main goroutine reads them.

Without buffering, the sender would pause until another goroutine was ready to receive.
