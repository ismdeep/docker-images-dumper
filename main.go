package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	for _, image := range Config.Images {
		for _, tag := range image.Tags {
			imageLink := fmt.Sprintf("%v:%v", image.Name, tag)
			cmd := exec.Command("docker", "pull", imageLink)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}

			_ = exec.Command("/Users/ismdeep/Data/shells/docker-image-save", imageLink).Run()
			fmt.Println("Done at => ", imageLink)
		}
	}
}
