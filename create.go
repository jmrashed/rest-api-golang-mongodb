package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var folders = []string{
	"cmd",
	"internal/config",
	"internal/handler",
	"internal/middleware",
	"internal/model",
	"internal/repository",
	"internal/router",
	"internal/service",
	"internal/util",
	"migrations",
	"scripts",
}

var files = map[string]string{
	"cmd/main.go": `
package main

import "fmt"

func main() {
	fmt.Println("Hello, RESTful API!")
}
`,
	"internal/config/config.go": `
package config

// Configuration setup and management goes here
`,
	"internal/handler/handler.go": `
package handler

// Request handling logic goes here
`,
	"internal/middleware/middleware.go": `
package middleware

// Middleware components go here
`,
	"internal/model/model.go": `
package model

// Data models and structures go here
`,
	"internal/repository/repository.go": `
package repository

// Database operations go here
`,
	"internal/router/router.go": `
package router

// Router setup and route handling goes here
`,
	"internal/service/service.go": `
package service

// Business logic layer goes here
`,
	"internal/util/util.go": `
package util

// Utility functions go here
`,
	"migrations/001_init.up.sql": `
-- Initial database migration script
`,
}

func createFolders() error {
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func createFiles() error {
	for path, content := range files {
		dir := filepath.Dir(path)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		_, err = file.WriteString(content)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func main() {
	if err := createFolders(); err != nil {
		fmt.Println("Error creating folders:", err)
		return
	}

	if err := createFiles(); err != nil {
		fmt.Println("Error creating files:", err)
		return
	}

	fmt.Println("Project structure and files created successfully!")
}
