# Railway Deployment Log Streamer
GO Project that stream logs from an active Railway deployment into a text file created locally

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
You can then open the file and watch it fill up with logging information about the deployment directly from the Railway Logs API.
