package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"sync"
)

func main() {
	usr, _ := user.Current()
	projectsDir := filepath.Join(usr.HomeDir, "/projects")

	var wg sync.WaitGroup

	dirs, err := os.ReadDir(projectsDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		projectPath := filepath.Join(projectsDir, d.Name())
		if _, err := os.Stat(filepath.Join(projectPath, "/docker-compose.override.yml")); err != nil {
			continue
		}

		wg.Add(1)
		go down(&wg, projectPath)
	}

	wg.Wait()
}

func down(wg *sync.WaitGroup, path string) {
	defer wg.Done()

	out, err := exec.Command("docker-compose", "--project-directory="+path, "down").Output()
	if err != nil {
		fmt.Println(path, err)
	} else {
		fmt.Println(path, "Command Successfully Executed")
		output := string(out[:])
		fmt.Println(output)
	}
}
