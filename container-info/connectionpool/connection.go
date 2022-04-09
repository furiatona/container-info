package connectionpool

import (
	"container-info/models"

	_ "github.com/lib/pq"
)

var (
	CreateConnection = models.InitMySQLPool()
)

