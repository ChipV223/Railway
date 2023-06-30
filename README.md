# Railway Deployment Log Streamer
GO Project that stream logs from an active Railway deployment into a text file created & saved locally

## Introduction

This program allows you to launch a Railway deployment from your local machine and creates a file with the deployment logging which you can view in real time

## Prerequisites

+ Both [Railway CLI](https://docs.railway.app/develop/cli) and [Go](https://go.dev/dl/) are installed


## Instructions

1. In a terminal window, clone the logStreamer project repository 
   - HTTPS: `git clone https://github.com/ChipV223/Railway.git`
   - SSH: `git clone git@github.com:ChipV223/Railway.git`

1. In the same terminal window, naviage to your project's root directory 

1. Run `railway login` so that you're connected to your Railway account

1. Run `railway link` to connect to your Railway project. To ensure that you're connected to the correct project, run `railway status`

1. While still in the project root directory, execute the logStreamer program by specifying the path(i.e. `go run /path/to/main.go`)

Answer the prompt at the start of the program and you should now see a new file in the location that you provided in the prompt.
You can then open the file and watch it fill up with logging information about the deployment pulled directly from the Railway Logs API.

## Code Breakdown

At the start of our GO program and before the Railway deployment starts, there's first a prompt requesting the user to provide the directory location and name of the file that will 
save the logs from the upcoming deployment.

```
fmt.Println("Provide the name & save location for your Railway deployment logs: ")
var logFile string
fmt.Scanln(&logFile)
```

Then the program executes the first Railway CLI command, `railway up`, which starts the Railway deployment process and in turn creates the log file based on the answer from the 
afromnentioned prompt and adds the logging from the command to the new file. There's error handling for the command and logfile creation to insure that if the command and/or the 
creating of the log file fails for whatever reason, the program terminates and the error displays on the terminal window. In addition, there's code to make sure that the logfile 
is closed for any I/O and the underlying writer buffer is flushed once the stdout from the CLI command has been added.

```
cmd1 := exec.Command("railway", "up")
fmt.Println("Railway Up Command Executed..")
outfile, err := os.Create(logFile)
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
```

Afterwards, the program executes the second Railway CLI command, `railway logs -b` so that the build logs of the deployment can be captured. Since the logfile is already created 
from the last step, we tell GO to open the file and append the output from the CLI command to the file.

```
cmd2 := exec.Command("railway", "logs", "-b")
fmt.Println("Railway Up Command Completed. Starting Railway Deployment..")
outfile2, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
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
```


Lastly, we execute the third Railway CLI command, `railway logs -d`, to capture the deployment logs after the project has been built in Railway

```
cmd3 := exec.Command("railway", "logs", "-d")
fmt.Println("Railway Deployment Completed. Just Finishing Up and Gatering the Logs...")
outfile3, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
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
```