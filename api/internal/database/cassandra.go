package database

import (
	"os"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

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

func MustInit() {
	var err error
	cluster := getClusterConfig()
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}
