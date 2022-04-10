package main

import(
	"fmt"
	"time"
	"runtime"
	"log"
	"container-info/config"
	"container-info/logging"
	"container-info/connectionpool"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
func getinfo(c *gin.Context){
	current_date := time.Now().Format("2006-01-02 15:04:05") 
	os := runtime.GOOS
	ID := "Xendit - Trial - Dheny Priatna - 2022-04-08 - " + current_date
	container_data := map[string]interface{}{
	"ID":ID,
	"current_date":current_date, 
	"OS":os, 
}
	logging_stdout(container_data)
	logging_db(container_data)
	c.JSON(http.StatusOK, container_data)
}
func logging_stdout(container_data map[string]interface{}){
	fmt.Printf("ID: %v\n", container_data["ID"])
	fmt.Printf("Current date: %v\n", container_data["current_date"])
	fmt.Printf("OS Type: %v\n", container_data["OS"])
}
func logging_db(container_data map[string]interface{}){
	insert_db(container_data)
}

// function to insert into database.
func insert_db(g map[string]interface{}) {
	querytext := fmt.Sprintf("INSERT INTO ContainerInfo (ID,Date,OS,Memory_total,Memory_used,Memory_cached,Memory_free,Memory_percentage,Cpu_user,Cpu_system,Cpu_idle) values ('%s','%s','%s','%v','%v','%v','%v','%v','%v','%v','%v')", g["ID"], g["current_date"], g["OS"],0,0,0,0,0,0,0,0)
	fmt.Println(querytext)
	_, err := connectionpool.CreateConnection.Exec(querytext)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{}).Errorf("Unable to execute the query. error:%v.", err)
		fmt.Println(err)
	}
}

func ping(c *gin.Context){
	c.JSON(http.StatusOK, "pong")
}
func main(){
	config.InitializeAppConfig()
	router := gin.Default()
	router.GET("/info", getinfo)
	router.GET("/ping", ping)
	log.Fatal(router.Run(":8080"))
	fmt.Println("Running..")
}