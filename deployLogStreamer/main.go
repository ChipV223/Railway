package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cmd1 := exec.Command("railway", "up")
	fmt.Println("Railway Up Command Executed..")
	outfile, err := os.Create("/path/to/log/directory/log.txt")
	if err != nil {
		panic(err)
	}

	defer outfile.Close()
	cmd1.Stdout = outfile

	writer := bufio.NewWriter(outfile)
	defer writer.Flush()

	err = cmd1.Run()
	if err != nil {
		panic(err)
	}

	cmd2 := exec.Command("railway", "logs", "-b")
	fmt.Println("Railway Up Command Completed. Starting Railway Deployment..")
	outfile2, err := os.OpenFile("/path/to/log/directory/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	defer outfile2.Close()
	cmd2.Stdout = outfile2

	writer = bufio.NewWriter(outfile2)
	defer writer.Flush()

	err = cmd2.Run()
	if err != nil {
		panic(err)
	}

	cmd3 := exec.Command("railway", "logs", "-d")
	fmt.Println("Railway Deployment Completed. Just Finishing Up and Gatering the Logs...")
	outfile3, err := os.OpenFile("/path/to/log/directory/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}

	defer outfile3.Close()
	cmd3.Stdout = outfile3

	writer = bufio.NewWriter(outfile3)
	defer writer.Flush()

	err = cmd3.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Log Gathering Completed..")

}
