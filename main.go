package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	for file := range jobs {
		fmt.Printf("Worker %d started processing %s\n", id, file)

		time.Sleep(2 * time.Second)

		fmt.Printf("Worker %d finished processing %s\n", id, file)

		results <- fmt.Sprintf("%s processed by worker %d", file, id)
	}
}


func main() {
	files := []string {
		"image1.jpg",
		"image2.png",
		"document.pdf",
		"video.mp4",
		"notes.txt",
	}

	jobs := make(chan string, len(files))
	results := make(chan string, len(files))

	// Here we create the workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send the jobs to the workers
	for _, file := range files {
		jobs <- file
	}

	close(jobs)

	// Here we read the results
	for a := 1; a <= len(files); a++ {
		fmt.Println(<-results)
	}
}