package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	renderIntro()
	for {
		renderMenu()
		command := readCommand()

		switch command {
		case 1:
			initLog()
			break
		case 2:
			fmt.Println("Showing log")
			renderLogs()
			break
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
			break
		default:
			fmt.Println("unknow command")
			os.Exit(-1)
		}
	}
}

func renderIntro() {
	var name string = "Rodrigo"
	var version float32 = 1.1
	fmt.Println("hello world", name)
	fmt.Println("Version:", version)
}

func renderMenu() {
	fmt.Println("1 - Start")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)

	fmt.Println("Command", command)
	return command
}

func getUserData() (string, string, int) {
	firstName := "Rodrigo"
	lastName := "Cabral"
	age := 23

	return firstName, lastName, age
}

func initLog() {
	fmt.Println("Starting...")
	urls := readFromFile()
	for i := 0; i < monitoring; i++ {
		for i, url := range urls {
			fmt.Println("Testing URL", i+1, ":", url)
			verify(url)
		}
		time.Sleep(delay * time.Second)
	}
}

func verify(url string) {
	msg := "Result: loaded with success!"
	result, error := http.Get(url)
	if error != nil {
		fmt.Println("Error", error)
		os.Exit(-1)
	}

	if result.StatusCode != 200 {
		msg = "Result: went wrong. Status Code:" + fmt.Sprintf("%d", result.StatusCode)
	}

	writeLog(url, msg)

	fmt.Println(msg)
}

func readFromFile() []string {
	sites := []string{}
	file, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("Error", error)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		fmt.Println(line)

		if error == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func writeLog(url string, log string) {
	file, error := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if error != nil {
		fmt.Println("Error", error)
		os.Exit(-1)
	}

	dateTime := time.Now().Format("02/01/2006 15:04:05")

	file.WriteString(dateTime + " - " + url + ": " + log + "\n")
	file.Close()
}

func renderLogs() {
	file, error := os.ReadFile("logs.txt")

	if error != nil {
		fmt.Println("Error", error)
		os.Exit(-1)
	}

	fmt.Println(string(file))
}

// func showNames() {
// 	// slices
// 	names := []string{"rodrigo", "cabral", "joao"}
// 	// len(names) - returns the length of the array/slice
// 	// cap(names) - returns the capacity of the array/slice
// 	names = append(names, "lis")
// 	fmt.Println(names)
// }
