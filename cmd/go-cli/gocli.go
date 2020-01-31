package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's your query IPs, CNAMEs, MX records and Name Servers!"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "tutorialedge.net",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns", // name server
			Usage: "Looks up the NameServers for a Particular Host",
			Flags: myFlags,
			// use ns command to execute it
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil { // error check up
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Print(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the Ip addresses for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
