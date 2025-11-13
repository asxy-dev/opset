package main
type Server struct {
	Name string `json:"name"` //name of server
	Host string `json:"host"` //ip or domain
	User string `json:"user"` //ssh username (ubuntu etc)
	KeyPath string `json:"key_path"` // path to ssh private key

	//server struct that hold information about remote machine
}
type Task struct {
	Command string `json:"command"` // the shell command to execute
	Interval string `json:"interval"` //time
	
	//task struct defines a schedule comand
}
type Result struct{
	ServerName string //which server ran it
	Command string //which command was executed
	Output string //what came back (stdout)
	Error string //any error message
	TimeStamp string //when it ran
}
//complete