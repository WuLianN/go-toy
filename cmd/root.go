package main

import (
	"github.com/spf13/cobra"
	"log"
	sql "github.com/WuLianN/go-toy/cmd/utils/sql"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(sql.SqlCmd)
}

func main() {
	err := Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}