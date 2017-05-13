package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var ServerBind string
var MongoDBIpPort string

func init() {
	cmd := cobra.Command{}
	cmd.Flags().StringVarP(&ServerBind, "bind", "b", "0.0.0.0:5015", "bind de interface ip:porta")
	cmd.Flags().StringVarP(&MongoDBIpPort, "mongodb", "m", "127.0.0.1:27017", "conexao mongo")
	//"Kartola FC é um wrapper da API do cartolafc"

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

