# Time Tracker CLI

A simple time tracker CLI app built with Go and Cobra.

## Table of Contents

- [Time Tracker CLI](#time-tracker-cli)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Commands](#commands)
  - [Contributing](#contributing)

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/sagor31h2/time-tracker.git
   cd time-tracker
   ```

2. Ensure you have Go installed (version 1.16+).

3. Install dependencies and tidy up the module:

   ```sh
   go mod tidy
   ```

4. Build the project:

   ```sh
   go build -o time-tracker
   ```

   > windows

```sh
$env:GOOS = "windows"; $env:GOARCH = "amd64"; go build -o timetracker.exe
```

## Usage

You can run the CLI directly using go run:

```sh
go run main.go [command]
```

Or use the built executable:

```sh
./time-tracker [command]
```

## Commands

- start: Start tracking time.

  ```sh
  ./time-tracker start
  ```

  Output:

  ```sql

  Started tracking time at: <current time>
  ```

- stop: Stop tracking time.

  ```sh
  ./time-tracker stop
  ```

  Output:

  ```sql
  Stopped tracking time at: <current time>
  ```

- status: View current tracking status.

  ```sh
  ./time-tracker status
  ```

  Output:

  ```
  Current tracking status.
  ```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.
