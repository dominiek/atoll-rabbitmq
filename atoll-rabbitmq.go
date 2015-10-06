
package main

import (
  "os"
  "log"
  "fmt"
  "github.com/codegangsta/cli"
)

func fatalError(err error) {
  log.Fatalf("Error: %v", err)
}

func main() {
  app := cli.NewApp()
  app.Name = "atoll-rabbitmq"
  app.Usage = "RabbitMQ monitoring plugin for Atoll"
  app.Version = "0.1.0"
  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "host",
      Value: "localhost",
      Usage: "RabbitMQ admin host",
    },
    cli.IntFlag{
      Name: "port",
      Value: 15672,
      Usage: "RabbitMQ admin port",
    },
    cli.StringFlag{
      Name: "user",
      Value: "guest",
      Usage: "RabbitMQ username",
    },
    cli.StringFlag{
      Name: "pass",
      Value: "guest",
      Usage: "RabbitMQ password",
    },
  }
  app.Action = func(c *cli.Context) {
    rabbitmq := RabbitMQ {
      c.String("host"),
      uint16(c.Int("port")),
      c.String("user"),
      c.String("pass"),
      };
    data, err := rabbitmq.Monitor();
    if err != nil {
      fatalError(err);
    } else {
      fmt.Printf("%s", data);
    }
  }

  app.Run(os.Args)
}
