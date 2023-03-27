package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//var out bytes.Buffer
	//cmd := exec.Command("php", "file2yaml.php", "istio.yaml")
	//cmd.Stdin = strings.NewReader("and old falcon")
	//cmd.Stdout = &out
	//err := cmd.Run()

	out, err := exec.Command("php", "file2yaml.php", "istio.yaml").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Yaml content: %s\n", string(out))
}
