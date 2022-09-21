package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var ctx = context.Background()
var arguments = os.Args
var redisHost string
var redisPort int
var redisPassword string

func init() {
	const (
		defaultPassword = ""
		usagePassword   = "Redis-Server Auth"
		defaultHost     = "127.0.0.1"
		usageHost       = "Redis-server listening IP"
		defaultPort     = 6379
		usagePort       = "Redis-server listening port"
	)
	flag.StringVar(&redisPassword, "password", defaultPassword, usagePassword)
	flag.StringVar(&redisPassword, "P", defaultPassword, usagePassword)
	flag.StringVar(&redisHost, "host", defaultHost, usageHost)
	flag.StringVar(&redisHost, "h", defaultHost, usageHost)
	flag.IntVar(&redisPort, "port", defaultPort, usagePort)
	flag.IntVar(&redisPort, "p", defaultPort, usagePort)
}

func rClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		PoolSize: 1,
	})

	return client
}

func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)

	return nil
}

func checkhealthz(c echo.Context) error {

	client := rClient()
	//measure latency
	err := ping(client)
	client.Close()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Redis ERROR\n")
	}

	return c.String(http.StatusOK, "Redis OK\n")
}

func main() {
	if len(arguments) == 2 {
		fmt.Println("Please provide host:port <password>.")
		os.Exit(1)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/healtz", checkhealthz)

	e.Logger.Fatal(e.Start(":7550"))
}
