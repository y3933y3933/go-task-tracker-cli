package main

import (
	"fmt"
	"strconv"
	"time"
)

// Task 結構體，代表一個任務
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`    // "todo", "in-progress", "done"
	CreatedAt   time.Time `json:"createdAt"` // 任務創建時間
	UpdatedAt   time.Time `json:"updatedAt"` // 任務最後更新時間
}

// addTask 新增任務
func addTask(description string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("讀取任務時出錯:", err)
		return
	}

	newID := len(tasks) + 1
	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("保存任務時出錯:", err)
	} else {
		fmt.Printf("成功新增任務 (ID: %d)\n", newID)
	}
}

// updateTask 更新任務描述
func updateTask(idStr, newDescription string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("讀取任務時出錯:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("無效的任務 ID:", idStr)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()

			if err := saveTasks(tasks); err != nil {
				fmt.Println("保存任務時出錯:", err)
			} else {
				fmt.Printf("成功更新任務 (ID: %d)\n", id)
			}
			return
		}
	}

	fmt.Printf("未找到 ID 為 %d 的任務\n", id)
}

// deleteTask 刪除任務
func deleteTask(idStr string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("讀取任務時出錯:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("無效的任務 ID:", idStr)
		return
	}

	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}

	if err := saveTasks(newTasks); err != nil {
		fmt.Println("保存任務時出錯:", err)
	} else {
		fmt.Printf("成功刪除任務 (ID: %d)\n", id)
	}
}

// markInProgress 標記任務為進行中
func markInProgress(idStr string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("讀取任務時出錯:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("無效的任務 ID:", idStr)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "in-progress"
			tasks[i].UpdatedAt = time.Now()

			if err := saveTasks(tasks); err != nil {
				fmt.Println("保存任務時出錯:", err)
			} else {
				fmt.Printf("成功將任務標記為進行中 (ID: %d)\n", id)
			}
			return
		}
	}

	fmt.Printf("未找到 ID 為 %d 的任務\n", id)
}

// markDone 標記任務為已完成
func markDone(idStr string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("讀取任務時出錯:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("無效的任務 ID:", idStr)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now()

			if err := saveTasks(tasks); err != nil {
				fmt.Println("保存任務時出錯:", err)
			} else {
				fmt.Printf("成功將任務標記為已完成 (ID: %d)\n", id)
			}
			return
		}
	}

	fmt.Printf("未找到 ID 為 %d 的任務\n", id)
}

// listTasks 列出所有或按狀態過濾的任務
func listTasks(status string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("讀取任務時出錯:", err)
		return
	}

	fmt.Println("任務列表:")
	for _, task := range tasks {
		if status == "all" || task.Status == status {
			fmt.Printf("[%d] %s (%s)\n", task.ID, task.Description, task.Status)
		}
	}
}
