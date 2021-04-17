package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const projectsDir = "/home/maz/projects"

func main() {
	c := make(chan string)

	dirs, err := os.ReadDir(projectsDir)
	if err != nil {
		log.Fatal(err)
	}

	num := 0
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		projectPath := filepath.Join(projectsDir, d.Name())
		if _, err := os.Stat(projectPath + "/docker-compose.override.yml"); err != nil {
			continue
		}

		go down(projectPath, c)
		num++
	}

	for i := 0; i < num; i++ {
		fmt.Println(<-c)
	}
}

func down(path string, c chan string) {
	_, err := exec.Command("docker-compose", "--project-directory="+path, "down").Output()
	if err != nil {
		c <- path + " " + err.Error()
	} else {
		// output := string(out[:])
		c <- path + " Command Successfully Executed"
	}
}
