# miniflow 

miniflow is a simple command line task executor.

It puts tasks in a dag and resolve dependencies between them.

## Features

- find task dependency cycle
- resolve task dependencies
- execute tasks in parallel
- report failed tasks and downstream tasks

## Usage

1. define tasks in json format

```json
{
    "name": "miniflow0",
    "parallel": 2,
    "tasks": [
        {"id": 0, "cmd": "echo task0; sleep 5", "downstream":[2]},
        {"id": 1, "cmd": "echo task1; exit 1", "upstream":[0]},
        {"id": 2, "cmd": "echo task2; sleep 3"},
        {"id": 3, "cmd": "echo task3; sleep 3", "upstream": [0]},
        {"id": 4, "cmd": "echo task4", "upstream": [1]},
        {"id": 5, "cmd": "echo task5", "upstream": [4, 3]}
    ]
}
```

2. run miniflow

```bash
go build && chmod +x miniflow
./miniflow /path/to/tasks.json
```
