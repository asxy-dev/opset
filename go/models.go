package main

// Server and Task model definitions

type Server struct {
	Name string `json:"name"`
	Host string `json:"host"`
	User string `json:"user"`
	Port int    `json:"port"`
}

type Task struct {
	Name    string   `json:"name"`
	Command string   `json:"command"`
	Targets []string `json:"targets"`
	Cron    string   `json:"cron,omitempty"`
}
