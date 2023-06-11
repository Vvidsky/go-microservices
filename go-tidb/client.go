package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "go-tidb/employeepb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

var client pb.EmployeeServiceClient

type Employee struct {
	EmpId    int
	EmpFName string
	EmpLName string
	Salary   float32
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client = pb.NewEmployeeServiceClient(conn)

	r := gin.Default()
	r.GET("/employees", getAllEmployee)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func getAllEmployee(c *gin.Context) {
	stream, err := client.GetEmployeeList(context.Background(), &pb.GetEmployeeListRequest{})
	if err != nil {
		log.Fatalf("client.GetEmployeeList failed: %v", err)
	}
	for {
		employee, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		emp := employee.GetEmployee()
		c.JSON(200, gin.H{
			"employee": &Employee{
				EmpId:    int(emp.GetEmpId()),
				EmpFName: emp.GetEmpFName(),
				EmpLName: emp.GetEmpLName(),
				Salary:   emp.Salary,
			},
		})
	}
}
