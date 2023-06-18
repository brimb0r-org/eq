package internal_excel

import (
	"context"
	"fmt"
	"gSheets/application/pkg/eq_math"
	"gSheets/application/pkg/excel"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/sheets/v4"
	"strconv"
)

func GetDataToProcess(srv *excel.ExcelConfig) (string, error) {
	var (
		savings      interface{}
		totalSpent   interface{}
		remain       interface{}
		statusReport string
	)
	spreadSheetID := "1wqtoGe95h0s0JoXfPCFQxHUlT1oUA52NHtJEKgbBets"
	readRange := "Sheet1!C6:C23"
	resp, err := srv.NewClient().Spreadsheets.Values.Get(spreadSheetID, readRange).Do()
	if err != nil {
		log.Fatal().Msgf("Unable to retrieve data from sheet: %v", err)
		return statusReport, err
	}

	if len(resp.Values) == 0 {
		log.Print("No Data Found")
	}
	savings = 15000
	totalSpent = getSpentBudget(resp.Values)
	remain = eq_math.Sub(savings, totalSpent)
	statusReport = getStatus(remain)
	spentBudgetValueRowUpdate := &sheets.ValueRange{Values: [][]interface{}{{statusReport}}}
	_, err = srv.NewClient().Spreadsheets.Values.Update(spreadSheetID, "I3", spentBudgetValueRowUpdate).ValueInputOption("USER_ENTERED").Context(context.Background()).Do()
	if err != nil {
		log.Fatal().Msgf("Unable to post data to sheet: %v", err)
		return statusReport, err
	}
	return statusReport, err
}

func getSpentBudget(rows [][]interface{}) int64 {
	sum := 0
	for _, row := range rows {
		value, err := strconv.Atoi(fmt.Sprintf("%v", row[0]))
		if err != nil {
			fmt.Println(err)
		}
		sum += value
	}
	spentBudget := int64(sum)
	return spentBudget
}

func getStatus(spent interface{}) string {
	var spentConvert = int64(spent.(float64))
	switch s := spentConvert; {
	case s >= 10000:
		return "safe"
	case s < 1200:
		return "pull from another account"
	case s < 10000 && s >= 1200:
		return "managed"
	default:
		return "revaluate"
	}
}
