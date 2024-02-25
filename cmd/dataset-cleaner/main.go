package main

import (
	"fmt"
	"infra-task-solution/pkg/gcs"
	"infra-task-solution/pkg/processing"
	"infra-task-solution/pkg/verification"
	"sync"
	"log"
)

func main() {
	log.Println("This is a log message from Jeremy Wang")
	datasetFolder := "assets/reddit_dataset"
	processed := make(chan processing.FileContent)

	var waitGroup sync.WaitGroup

	go processing.FilterPII(datasetFolder, processed)
	log.Println("in func main after FilterPII")

	waitGroup.Add(1)
	go func() {
		log.Println("in func main go func()")
		defer waitGroup.Done()
		// waitGroup.Wait()
		log.Println("in func main go func() before uploadProcessedFiles")
		uploadProcessedFiles(processed)
	}()

	waitGroup.Wait()

	log.Println("in func main before verification.VerifyFiles")

	allVerified, err := verification.VerifyFiles(datasetFolder)
	if err != nil {
		fmt.Printf("Verification failed: %v\n", err)
		log.Println("Verification failed:", err)
	} else if allVerified {
		fmt.Println("Verification completed successfully")
		log.Println("Verification completed successfully")
	} else {
		fmt.Println("Not all files are verified successfully")
		log.Println("Not all files are verified successfully")
	}
	log.Println("No error func main")
}

func uploadProcessedFiles(processed <-chan processing.FileContent) {
	log.Println("Entered func uploadProcessedFiles")
	for file := range processed {
		if file.OriginalErr != nil {
			fmt.Printf("Error processing file %s: %v\n", file.FileName, file.OriginalErr)
			log.Println("Error processing file", file.FileName, ":", file.OriginalErr)
			continue
		}

		err := gcs.UploadFile(file.Content, file.FileName)
		if err != nil {
			fmt.Printf("Could not upload file %s: %v\n", file.FileName, err)
			log.Println("Could not upload file", file.FileName, ":", err)
			return
		}
	}
}
