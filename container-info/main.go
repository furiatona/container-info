package main

import(
	"fmt"
	"time"
	"runtime"
	"log"
	"bytes"
	"container-info/models"
	"container-info/config"
	"container-info/logging"
	"container-info/connectionpool"
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
func getinfo(c *gin.Context){
	current_date := time.Now()
	os := runtime.GOOS
	container_data := map[string]interface{}{
	"Xendit - Trial - Dheny Priatna - 8 April 2022 - Current date":current_date, 
	"OS":os, 
}
	logging_stdout(container_data)
	logging_db(container_data)
	c.JSON(http.StatusOK, container_data)
}
func logging_stdout(container_data map[string]interface{}){
	fmt.Printf("Xendit - Trial - Dheny Priatna - 8 April 2022 - Current date: %v\n", container_data["current_date"])
	fmt.Printf("OS Type: %v\n", container_data["OS"])
	fmt.Printf("memory total: %d bytes\n", container_data["memory_total"])
	fmt.Printf("memory used: %d bytes\n", container_data["memory_used"])
	fmt.Printf("memory cached: %d bytes\n", container_data["memory_cached"])
	fmt.Printf("memory free: %d bytes\n", container_data["memory_free"])
	fmt.Printf("memory used percentage: %f %% \n",container_data["memory_percentage"])
	fmt.Printf("cpu user: %f %%\n", container_data["cpu_user"])
	fmt.Printf("cpu system: %f %%\n", container_data["cpu_system"])
	fmt.Printf("cpu idle: %f %%\n", container_data["cpu_idle"])

}
func logging_db(container_data map[string]interface{}){
	jsonStr, err := json.Marshal(container_data)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{}).Errorf("Unable to execute the query. error:%v.", err)
		fmt.Println(err)
	}
	var container_info_data models.ContainerInfo
	err = json.NewDecoder(bytes.NewReader(jsonStr)).Decode(&container_info_data)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{}).Errorf("Unable to execute the query. error:%v.", err)
		fmt.Println(err)
	}
	insert_db(container_info_data)
}

// function to insert into database.
func insert_db(g models.ContainerInfo) {
	querytext := fmt.Sprintf("INSERT INTO ContainerInfo (ID,Date,OS,Memory_total,Memory_used,Memory_cached,Memory_free,Memory_percentage,Cpu_user,Cpu_system,Cpu_idle) values ('%s','%s','%s','%v','%v','%v','%v','%v','%v','%v','%v')", g.ID, g.Date, g.OS,g.Memory_total, g.Memory_used, g.Memory_cached,g.Memory_free, g.Memory_percentage,g.Cpu_user,g.Cpu_system,g.Cpu_idle)
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