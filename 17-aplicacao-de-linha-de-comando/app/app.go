package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de enderecos na internet",
			Flags: [].cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "devbook.com.br"
				},
			},
			Action: buscarIps,
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