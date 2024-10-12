Sure! Here is the English version of your README:

---

# Task CLI

[Roadmap](https://roadmap.sh/projects/task-tracker)

Task CLI is a simple command-line application for managing your to-do items and tasks. Users can add, update, delete tasks, and mark tasks as "todo", "in-progress", or "done". The tasks are stored in a JSON file for future reading and updates.

## Features
- **Add Task**: Add a new task to the task list
- **Update Task**: Update the description of an existing task
- **Delete Task**: Delete a specific task
- **Mark Task Status**: Mark tasks as "in-progress" or "done"
- **List Tasks**: List all tasks or filter by status

## Directory Structure

```
task-cli/
│
├── main.go          # Main entry point, handles command-line input
├── task.go          # Task structure and operations
└── storage.go       # Logic for reading and saving JSON files
```

## Installation and Usage

### 1. Clone the repository

```bash
git clone <your-repository-url>
cd task-cli
```

### 2. Build the program

Run the following command to compile the executable:

```bash
go build -o task-cli
```

### 3. Run the program

Use the compiled `task-cli` file to run commands:

```bash
./task-cli <command> [arguments]
```

## Commands

### Add a new task

```bash
./task-cli add "Buy groceries"
```

Output:
```
Task added successfully (ID: 1)
```

### Update an existing task

```bash
./task-cli update 1 "Buy groceries and cook dinner"
```

Output:
```
Task updated successfully (ID: 1)
```

### Delete a task

```bash
./task-cli delete 1
```

Output:
```
Task deleted successfully (ID: 1)
```

### Mark a task as in-progress

```bash
./task-cli mark-in-progress 1
```

Output:
```
Task marked as in-progress (ID: 1)
```

### Mark a task as done

```bash
./task-cli mark-done 1
```

Output:
```
Task marked as done (ID: 1)
```

### List all tasks

```bash
./task-cli list
```

Output:
```
Task list:
[1] Buy groceries (todo)
```

### List tasks by status

List all tasks in progress:

```bash
./task-cli list in-progress
```

Output:
```
Task list:
[1] Buy groceries and cook dinner (in-progress)
```

List all completed tasks:

```bash
./task-cli list done
```

Output:
```
Task list:
[1] Buy groceries and cook dinner (done)
```

## Notes
- **File Storage**: Tasks are stored in a `tasks.json` file in the current directory where the program is executed.
- **Error Handling**: Ensure that the task ID provided is a valid integer, otherwise an invalid task ID message will be displayed.

## Contributions

Contributions are welcome! You can submit Pull Requests or report issues and suggestions.

---

