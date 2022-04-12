package main

import(
	"fmt"
	"runtime"
	"log"
	"container-info/config"
	"container-info/logging"
	"container-info/connectionpool"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	guuid "github.com/google/uuid"
)
func getinfo(c *gin.Context){
	os := runtime.GOOS
	container_data := map[string]interface{}{
	"OS":os, 
}
	logging_stdout(container_data)
	logging_db(container_data)
	c.JSON(http.StatusOK, container_data)
}
func logging_stdout(container_data map[string]interface{}){
	fmt.Printf("OS Type: %v\n", container_data["OS"])
}
func logging_db(container_data map[string]interface{}){
	insert_db(container_data)
}

// function to insert into database.
func insert_db(g map[string]interface{}) {
	ID := guuid.New()
	querytext := fmt.Sprintf("INSERT INTO ContainerInfo (ID,Date,OS,Memory_total,Memory_used,Memory_cached,Memory_free,Memory_percentage,Cpu_user,Cpu_system,Cpu_idle) values ('%s','%s','%s','%v','%v','%v','%v','%v','%v','%v','%v')", ID.String(), "2007-01-02 15:04:05", g["OS"],0,0,0,0,0,0,0,0)
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