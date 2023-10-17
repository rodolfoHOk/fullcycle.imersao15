package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/application/grpc"
	"github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
