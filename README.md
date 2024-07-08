# Time Tracker CLI

A simple time tracker CLI app built with Go and [Cobra](https://github.com/spf13/cobra).

## Table of Contents

- [Demo](#Demo)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Contributing](#contributing)

## Demo

> Cli overview

![Cli view](/assets/images/1_overview.png)

> Task Commands

![Cli view](/assets/images/2_task_demo.png)

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/sagor31h2/timetracker.git
   cd timetracker
   ```

2. Ensure you have Go installed (version 1.22.4+).

3. Install dependencies and tidy up the module:

   ```sh
   go mod tidy
   ```

4. Build the project:

   ```sh
   go build -o timetracker
   ```

   > windows

   ```sh
   $env:GOOS = "windows"; $env:GOARCH = "amd64"; go build -o timetracker.exe
   ```

   > Linux

   ```sh
   export GOOS=linux
   export GOARCH=amd64
   go build -o timetracker
   ```

   > Mac Os

   ```sh
   export GOOS=darwin
   export GOARCH=amd64
   go build -o timetracker
   ```

## Usage

- You can run the CLI directly using go run:

  ```sh
  go run main.go [command]
  ```

- Or use the built executable:

  ```sh
  ./timetracker [command]
  ```

## Commands

- start: Start tracking time.

  ```sh
  ./timetracker start
  ```

  Output:

  ```sql

  Started tracking time at: <current time>
  ```

- stop: Stop tracking time.

  ```sh
  ./timetracker stop
  ```

  Output:

  ```sql
  Stopped tracking time at: <current time>
  ```

- status: View current tracking status.

  ```sh
  ./timetracker status
  ```

  Output:

  ```
  Current tracking status.
  ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.
