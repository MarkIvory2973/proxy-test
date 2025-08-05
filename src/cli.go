package main

import (
	"proxy-test/mihomo"

	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "Proxy Test",
	Short: "Mihomo delays tester",
	Long:  `A delays tester for Mihomo.`,
}

func initCliCmd() {
	initTestCmd()

	err := cliCmd.Execute()
	if err != nil {
		logger.Fatal(err)
	}
}

var tls bool
var addr string
var port int
var secret string
var group string
var url string
var timeout int
var num int
var interval int
var weight float64

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test delays",
	Run: func(cmd *cobra.Command, args []string) {
		api = mihomo.API{
			TLS:    tls,
			Addr:   addr,
			Port:   port,
			Secret: secret,
		}

		test()
	},
}

func initTestCmd() {
	testCmd.Flags().BoolVarP(&tls, "tls", "t", false, "Use TLS")
	testCmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1", "Mihomo API Address")
	testCmd.Flags().IntVarP(&port, "port", "p", 9090, "Mihomo API Port")
	testCmd.Flags().StringVarP(&secret, "secret", "s", "", "Mihomo API Access Key")
	testCmd.Flags().StringVarP(&group, "group", "g", "", "Group Name")
	testCmd.Flags().StringVarP(&url, "url", "u", "https://www.google.com/generate_204", "Test URL")
	testCmd.Flags().IntVarP(&timeout, "timeout", "m", 3000, "Maximum Delay")
	testCmd.Flags().IntVarP(&num, "num", "n", 10, "Number of Times")
	testCmd.Flags().IntVarP(&interval, "interval", "i", 3, "Interval of Each Test")
	testCmd.Flags().Float64VarP(&weight, "weight", "k", 0.5, "Weight ∈ [0, 1]")

	cliCmd.AddCommand(testCmd)
}
