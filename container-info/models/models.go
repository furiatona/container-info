package models

import (
	"time"
)

type ContainerInfo struct{
	ID  	string `db:"ID" json:"ID"`
	Date 	time.Time `db:"Date" json:"current_date"`
	OS 		string `db:"OS" json:"OS"`
	Memory_total int `db:"Memory_Total" json:"memory_total"`
	Memory_used int  `db:"Memory_Used" json:"memory_used"`
	Memory_cached int `db:"Memory_cached" json:"memory_cached"`
	Memory_free int `db:"Memory_free" json:"memory_free"`
	Memory_percentage float64 `db:"Memory_percentage" json:"memory_percentage"`
	Cpu_user float64 `db:"Cpu_user" json:"cpu_user"`
	Cpu_system float64 `db:"Cpu_system" json:"cpu_system"`
	Cpu_idle float64 `db:"Cpu_idle" json:"cpu_idle"`
}