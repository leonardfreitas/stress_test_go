package commands

import (
	"fmt"
	"os"
	"time"

	stress "github.com/leonardfreitas/stress_test_go/pkg"
	"github.com/olekukonko/tablewriter"
)

type StressCommand struct {
	stress *stress.StressTester
}

func NewStressCommand(concurrency, requests int, url string, timeout int) *StressCommand {
	return &StressCommand{
		stress: stress.NewStressTester(concurrency, requests, url, time.Duration(timeout)*time.Second),
	}
}

func (s *StressCommand) Run() {
	testResult := s.stress.Run()

	tableData := [][]string{
		{"Total Requests", fmt.Sprintf("%d", testResult.TotalRequests)},
		{"Successfull Requests", fmt.Sprintf("%d", testResult.SuccessfullRequest)},
		{"Failed Requests", fmt.Sprintf("%d", testResult.FailedRequests)},
	}

	for status, total := range testResult.StatusCodes {
		tableData = append(tableData, []string{fmt.Sprintf("%d - Status code", status), fmt.Sprintf("%d", total)})
	}

	tableData = append(tableData, []string{"Elapsed Time", testResult.ElapsedTime.String()})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Metric", "Count"})

	for _, v := range tableData {
		table.Append(v)
	}

	table.Render()
}
