package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/c12s/meridian/internal/handlers"
	"github.com/c12s/meridian/internal/store"
	"github.com/c12s/meridian/pkg/api"
	oortapi "github.com/c12s/oort/pkg/api"
	pulsar_api "github.com/c12s/pulsar/model/protobuf"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	neo4jAddress := os.Getenv("NEO4J_ADDRESS")
	driver, err := neo4j.NewDriver(fmt.Sprintf("bolt://%s", neo4jAddress), neo4j.NoAuth())
	if err != nil {
		log.Fatal(err)
	}
	dbName := os.Getenv("NEO4J_DB_NAME")

	quotas := store.NewResourceQuotaNeo4jStore(driver, dbName)
	apps := store.NewAppNeo4jStore(driver, dbName, quotas)
	namespaces := store.NewNamespaceNeo4jStore(driver, dbName, quotas, apps)
	conn, err := grpc.NewClient(os.Getenv("PULSAR_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	pulsar := pulsar_api.NewSeccompServiceClient(conn)
	administrator, err := oortapi.NewAdministrationAsyncClient(os.Getenv("NATS_ADDRESS"))
	if err != nil {
		log.Fatalln(err)
	}
	meridian := handlers.NewMeridianGrpcHandler(namespaces, apps, pulsar, quotas, administrator)

	s := grpc.NewServer()
	api.RegisterMeridianServer(s, meridian)
	reflection.Register(s)

	lis, err := net.Listen("tcp", os.Getenv("LISTEN_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-shutdown

	s.GracefulStop()
}
