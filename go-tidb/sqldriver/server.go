package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	pb "go-tidb/employeepb"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
	dsn  = "root:@tcp(127.0.0.1:4000)/test?charset=utf8mb4"
)

var (
	db    *sql.DB
	dbErr error
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
}

func (*server) GetEmployeeList(req *pb.GetEmployeeListRequest, stream pb.EmployeeService_GetEmployeeListServer) error {
	employees, err := getAllEmployee(db)
	if err != nil {
		fmt.Println(err)
	}
	for _, employee := range employees {
		stream.Send(&pb.GetEmployeeListResponse{
			Employee: &pb.Employee{
				EmpId:    int32(employee.EmpId),
				EmpFName: employee.EmpFName,
				EmpLName: employee.EmpLName,
				Salary:   employee.Salary,
			},
		})
	}
	return nil
}

func main() {
	flag.Parse()
	db, dbErr = sql.Open("mysql", dsn)
	if dbErr != nil {
		panic(dbErr)
	}
	log.Printf("server connected to TiDB")
	defer db.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func openDB(driverName, dataSourceName string, runnable func(db *sql.DB)) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	runnable(db)
}
