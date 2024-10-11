package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
)

func getClusterConfig() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(os.Getenv("CONNECTION_STRING"))
	cluster.Port = 9042
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASSANDRA_USER"),
		Password: os.Getenv("CASSANDRA_PASSWORD"),
	}
	cluster.Keyspace = os.Getenv("CASSANDRA_KEYSPACE")
	cluster.Consistency = gocql.Quorum
	return cluster
}

func Init() {

	cluster := getClusterConfig()
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	ctx := context.Background()

	scanner := session.Query("SELECT id, name, amount FROM Expense").WithContext(ctx).Iter().Scanner()

	var id gocql.UUID
	var name string
	var amount int

	for scanner.Next() {
		err = scanner.Scan(&id, &name, &amount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Expense:", id, name, amount)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
