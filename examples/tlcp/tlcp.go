package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	certDir := "/home/unvdb/.unvdb"
	fmt.Println("certDir:", certDir)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		//dbURL = "postgres://unvdb:unvdb@192.168.4.134:5678/?ssltlcp=1&sslrootcert=./cacert.pem&sslkey=./client.key&sslenccert=./client_enc.crt&sslenckey=./client_enc.key"
		dbURL = "postgres://unvdb:unvdb@192.168.4.134:5678/?ssltlcp=1&sslrootcert=/home/unvdb/unvdb_cert/cacert.pem&sslcert=/home/unvdb/unvdb_cert/client.crt&sslkey=/home/unvdb/unvdb_cert/client.key&sslenccert=/home/unvdb/unvdb_cert/client_enc.crt&sslenckey=/home/unvdb/unvdb_cert/client_enc.key"
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello, pool!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting) // 输出: Hello, pool!
}
