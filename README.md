# distributed-task-scheduler

## Building and Running Applications

This project contains multiple applications managed using Go workspaces. Below are the instructions to build and run each application.

### Prerequisites

- Ensure you have Go version 1.18 or later installed.
- Clone the repository and navigate into the project's root directory.

### Building Applications

```bash
#build worker
cd worker
go build -o worker

#build scheduler
cd mts
go build -o mts
```

### Running Applications

```bash
#run worker
cd worker
./worker

#run mts
cd mts
./mts
```
