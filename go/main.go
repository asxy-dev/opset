package main
import {
	"fmt"
	"encoding/json"
	"os"
}
type Server struct {
	Name string `json:"name"`
	Host string `json:"host"`
	User string `json:"user"`
	KeyPath string `json:"key_path"`
}
type Task struct {
	Command string `json:"command"`
	Interval string `json:"interval"`
}
func main(){
	if len(os.Args) < 2 { // to check user haave tyoer at least one argument
		fmt.Println("Usage: opset <command> [options]") //show usage
		return //exit
	}
	cmd := os.Args[1] //get the first argument 
	switch cmd {
		case "list-servers":
			listServers() //call list server function
		case "list-tasks":
			listTasks() //call lsit task function
		default:
			fmt.Println("Unknown command:", cmd) //unknown command
		}
}
func listServers() {
	data, err := os.ReadFile("data/servers.json")
	if err != nil {	
		fmt.Println("Error reading servers file:", err)
		return //end the program 
	}

var servers []Server //create slice to hold servers
	json.Unmarshal(data, &servers)
	fmt.Println("Servers:") 
	for _, server := range servers {
	fmt.Println("- %s (%s)/n", s.Name, s.Host) //print server details
	}
}
func listTasks() {	
	data , err := os.ReadFile("data/tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return //end the program
	}

var tasks []Task //create slice to hold tasks
	json.Unmarshal(data, &tasks)
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Println("- %s (every %s)\n", task.Command, task.Interval) //print task details
	}
}
