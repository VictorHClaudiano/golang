package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "kabum.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: 	"ip",                                   
			Usage: 	"Busca IPS de endereços na internet",
			Flags: 	flags,
			Action: buscarIps,
		},
		//go run main.go servidores --host kabum.com.br
		{
			Name: "servidores",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidor,
		},

	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips { 
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //Name server
	if erro != nil {
		log.Fatal(erro)
	}
	
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

//go biuld para compilar as alterações
// ./linha-de-comando ip --host amazon.com.br
// ./linha-de-comando servidores --host amazon.com.br