package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方式: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("使用方法: task-cli add <task description>")
		} else {
			addTask(os.Args[2])
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("使用方法: task-cli update <task id> <new description>")
		} else {
			updateTask(os.Args[2], os.Args[3])
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("使用方法: task-cli delete <task id>")
		} else {
			deleteTask(os.Args[2])
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("使用方法: task-cli mark-in-progress <task id>")
		} else {
			markInProgress(os.Args[2])
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("使用方法: task-cli mark-done <task id>")
		} else {
			markDone(os.Args[2])
		}
	case "list":
		if len(os.Args) == 2 {
			listTasks("all")
		} else {
			listTasks(os.Args[2])
		}
	default:
		fmt.Println("未知命令")
	}

}
