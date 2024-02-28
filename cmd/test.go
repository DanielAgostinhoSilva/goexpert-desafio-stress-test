/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/DanielAgostinhoSilva/goexpert-desafio-stress-test/src/infrastructure"
	"github.com/spf13/cobra"
)

var url string
var totalRequest int
var concurrentRequest int

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Sistema CLI em Go para realizar testes de carga em um serviço web",
	Long: `Sistema CLI em Go para realizar testes de carga em um serviço web. 
	O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.`,
	Run: func(cmd *cobra.Command, args []string) {
		stressTest := infrastructure.NewStressTestReport()
		stressTest.Execute(url, totalRequest, concurrentRequest)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado.")
	testCmd.Flags().IntVarP(&totalRequest, "requests", "r", 0, "Número total de requests.")
	testCmd.Flags().IntVarP(&concurrentRequest, "concurrency", "c", 0, "Número de chamadas simultâneas.")
	testCmd.MarkFlagsRequiredTogether("url", "requests", "concurrency")
}
