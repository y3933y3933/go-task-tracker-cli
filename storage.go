package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 定義存儲任務的 JSON 文件
const taskFile = "tasks.json"

// loadTasks 從 JSON 文件中讀取任務
func loadTasks() ([]Task, error) {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return []Task{}, nil
	}

	fileContent, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(fileContent, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// saveTasks 將任務保存到 JSON 文件
func saveTasks(tasks []Task) error {
	fileContent, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(taskFile, fileContent, 0644); err != nil {
		return err
	}

	return nil
}
