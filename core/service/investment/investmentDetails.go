package investment

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
	"github.com/tech/core/domain"
	"github.com/tech/utilities"
	"time"
)

var mapIdAndMonth = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func (s *Investment) AddNewInvestmentDetails(ctx context.Context, user *domain.User, request *domain.InvestmentBuyRequest) domain.Response {
	funcName := "AddNewInvestmentDetails"

	totalAmount, err := s.investment.InvestmentDetailsPersistence.GetUserInvestmentDetailsByUid(user.Id)
	if err == orm.ErrNoRows {
		newEntry := &domain.InvestmentDetails{
			Uid:                     user.Id,
			TotalInvestmentAmount:   request.TotalAmount,
			PresentInvestmentAmount: request.TotalAmount,
			UpdatedOn:               time.Now(),
			CreatedOn:               time.Now(),
		}

		_, err = s.investment.InvestmentDetailsPersistence.AddYpInvestmentDetails(newEntry)
		if err != nil {
			fmt.Println(funcName, " error while inserting entry to yp_investment_details table err := ", err)
			return domain.Response{Code: "452", Msg: err.Error()}
		}
	} else if err != nil {
		fmt.Println(funcName, " error while getting user investment details", err)
	} else {
		totalAmount.TotalInvestmentAmount = cast.ToString(cast.ToFloat64(totalAmount.TotalInvestmentAmount) + cast.ToFloat64(request.TotalAmount))
		totalAmount.PresentInvestmentAmount = cast.ToString(cast.ToFloat64(totalAmount.PresentInvestmentAmount) + cast.ToFloat64(request.TotalAmount))
		updateColumns := []string{
			"total_investment_amount",
			"present_investment_amount",
		}

		err = s.investment.InvestmentDetailsPersistence.UpdateYpInvestmentDetailsByColumns(totalAmount, updateColumns...)
		if err != nil {
			fmt.Println(funcName, " error while updating entry to yp_investment_details table err := ", err)
			return domain.Response{Code: "452", Msg: err.Error()}
		}
	}

	newEntry := &domain.InvestmentBuyDetails{
		Uid:            user.Id,
		Date:           request.Date,
		Month:          mapIdAndMonth[request.Month],
		Year:           request.Year,
		Type:           request.Category,
		Name:           request.Name,
		Symbol:         request.Symbol,
		Enable:         1,
		BuyCount:       request.BuyCount,
		AmountPerBuy:   request.AmountPerBuy,
		TotalAmount:    request.TotalAmount,
		RemainingCount: request.BuyCount,
		FdInterest:     request.RateOfInterest,
		UpdatedOn:      time.Now(),
		CreatedOn:      time.Now(),
	}

	// insert to yp_investment_buy_details table
	_, err = s.investment.InvestmentBuyPersistence.AddYpInvestmentBuyDetails(newEntry)
	if err != nil {
		fmt.Println(funcName, err)
		return domain.Response{Code: "452", Msg: err.Error()}
	}

	return domain.Response{Code: "200", Msg: "success"}
}

func (s *Investment) GetInvestmentDetails(ctx context.Context, user *domain.User) domain.Response {
	funcName := "GetInvestmentDetails"
	response := domain.InvestmentDetailsResponse{}
	response.UserInvested = true

	investmentCategory := make(map[string]float64)
	var investmentCategoryLabels []string
	var investmentCategoriesExpenses []string
	var investmentCategoryTable []domain.InvestmentDetailsToTable

	stockInvestment := make(map[string]float64)
	var stockCategoryLabels []string
	var stockCategoriesExpenses []string
	var stockCategoryTable []domain.InvestmentDetailsToTable

	mutualFundsInvestment := make(map[string]float64)
	var mutualFundsCategoryLabels []string
	var mutualFundsCategoryExpenses []string
	var mutualFundsCategoryTable []domain.InvestmentDetailsToTable

	fdInvestment := make(map[string]float64)
	var fdInvestmentLabels []string
	var fdInvestmentExpenses []string
	var fdCategoryTable []domain.InvestmentDetailsToTable

	totalAmount, err := s.investment.InvestmentDetailsPersistence.GetUserInvestmentDetailsByUid(user.Id)
	if err == orm.ErrNoRows {
		response.UserInvested = false
	} else if err != nil {
		fmt.Println(funcName, " error while getting user investment details", err)
	}

	response.TotalInvestedAmount = totalAmount.TotalInvestmentAmount
	response.PresentInvestedAmount = totalAmount.PresentInvestmentAmount
	response.ProfitAfter = totalAmount.ProfitAfterSellAmount
	response.LossAfter = totalAmount.LossAfterSellAmount

	getUserAllInvestmentDetails, err := s.investment.InvestmentBuyPersistence.GetAllYpInvestmentBuyDetailsByUid(user.Id)
	if err != nil {
		fmt.Println(funcName, " error while getting user all investment details", err)
	}

	if getUserAllInvestmentDetails == nil {
		fmt.Println(funcName, "no investment details of user")
		return domain.Response{Code: "200", Msg: "success", Model: response}
	}

	for _, details := range getUserAllInvestmentDetails {
		switch details.Type {
		case utilities.ConstStock:
			investmentCategory[details.Type] += cast.ToFloat64(details.RemainingCount) * cast.ToFloat64(details.AmountPerBuy)
			stockInvestment[details.Name] += cast.ToFloat64(details.RemainingCount) * cast.ToFloat64(details.AmountPerBuy)
		case utilities.ConstMutualFund:
			investmentCategory[details.Type] += cast.ToFloat64(details.RemainingCount) * cast.ToFloat64(details.AmountPerBuy)
			mutualFundsInvestment[details.Name] += cast.ToFloat64(details.RemainingCount) * cast.ToFloat64(details.AmountPerBuy)
		case utilities.ConstFd:
			investmentCategory[details.Type] += cast.ToFloat64(details.TotalAmount)
			fdInvestment[details.Name] += cast.ToFloat64(details.TotalAmount)
		}
	}

	for category, amount := range investmentCategory {
		investmentCategoryLabels = append(investmentCategoryLabels, category)
		investmentCategoriesExpenses = append(investmentCategoriesExpenses, cast.ToString(amount))
		percent := (amount / cast.ToFloat64(totalAmount.PresentInvestmentAmount)) * 100
		investmentCategoryTable = append(investmentCategoryTable, domain.InvestmentDetailsToTable{
			Name:       category,
			Amount:     cast.ToString(amount),
			Percentage: cast.ToString(percent),
		})
	}
	response.InvestmentCategoryLabels = investmentCategoryLabels
	response.InvestmentCategoriesExpenses = investmentCategoriesExpenses
	response.InvestmentCategoriesTable = investmentCategoryTable

	for category, amount := range stockInvestment {
		stockCategoryLabels = append(stockCategoryLabels, category)
		stockCategoriesExpenses = append(stockCategoriesExpenses, cast.ToString(amount))
		percent := (cast.ToInt(amount) / cast.ToInt(investmentCategory[utilities.ConstStock])) * 100
		stockCategoryTable = append(stockCategoryTable, domain.InvestmentDetailsToTable{
			Name:       category,
			Amount:     cast.ToString(amount),
			Percentage: cast.ToString(percent),
		})
	}

	response.StockInvestmentLabels = stockCategoryLabels
	response.StockInvestmentExpense = stockCategoriesExpenses
	response.StockInvestmentTable = stockCategoryTable

	for category, amount := range mutualFundsInvestment {
		mutualFundsCategoryLabels = append(mutualFundsCategoryLabels, category)
		mutualFundsCategoryExpenses = append(mutualFundsCategoryExpenses, cast.ToString(amount))
		percent := (cast.ToInt(amount) / cast.ToInt(investmentCategory[utilities.ConstMutualFund])) * 100
		mutualFundsCategoryTable = append(mutualFundsCategoryTable, domain.InvestmentDetailsToTable{
			Name:       category,
			Amount:     cast.ToString(amount),
			Percentage: cast.ToString(percent),
		})
	}

	response.MutualFundsInvestmentLabels = mutualFundsCategoryLabels
	response.MutualFundsInvestmentExpenses = mutualFundsCategoryExpenses
	response.MutualFundsInvestmentTable = mutualFundsCategoryTable

	for category, amount := range fdInvestment {
		fdInvestmentLabels = append(fdInvestmentLabels, category)
		fdInvestmentExpenses = append(fdInvestmentExpenses, cast.ToString(amount))
		percent := (cast.ToInt(amount) / cast.ToInt(investmentCategory[utilities.ConstFd])) * 100
		fdCategoryTable = append(fdCategoryTable, domain.InvestmentDetailsToTable{
			Name:       category,
			Amount:     cast.ToString(amount),
			Percentage: cast.ToString(percent),
		})
	}

	response.FDInvestmentLabels = fdInvestmentLabels
	response.FDInvestmentExpense = fdInvestmentExpenses
	response.FDInvestmentTable = fdCategoryTable

	return domain.Response{Code: "200", Msg: "success", Model: response}
}

func (s *Investment) GetUserHoldings(ctx context.Context, user *domain.User) domain.Response {
	//funcName := "GetUserHoldings"
	holdingsDetails := domain.InvestmentHoldingResponse{}

	holdingsDetails.StockInvestment, holdingsDetails.StockCurrent, holdingsDetails.StockTotalReturn, holdingsDetails.AllShareDetails = s.GetHoldingPrice(user, utilities.ConstStock)
	return domain.Response{Code: "200", Msg: "success", Model: holdingsDetails}
}

func (s *Investment) GetHoldingPrice(user *domain.User, category string) (invested string, current string, totalReturn string, holdingDetails []domain.HoldingDetails) {
	funcName := "GetHoldingPrice"
	checkName := make(map[string]bool)
	var allDetails = make(map[string]map[string]string)

	var investeds int
	var currents int
	var totalReturns int

	getDetails, err := s.investment.InvestmentBuyPersistence.GetInvestmentBuyDetailsByType(category, user.Id)
	if err == orm.ErrNoRows {
		return
	} else if err != nil {
		fmt.Println(funcName, err)
		return
	}

	for _, details := range getDetails {
		if checkName[details.Name] {
			allDetails[details.Name]["AvgCount"] = cast.ToString(cast.ToInt(allDetails[details.Name]["AvgCount"]) + 1)
			allDetails[details.Name]["Quantity"] = cast.ToString(cast.ToInt(allDetails[details.Name]["Quantity"]) + cast.ToInt(details.RemainingCount))
			allDetails[details.Name]["InvestedAmount"] = cast.ToString(cast.ToInt(allDetails[details.Name]["InvestedAmount"]) + (cast.ToInt(details.RemainingCount) * cast.ToInt(details.AmountPerBuy)))
			allDetails[details.Name]["AvgPrice"] = cast.ToString((cast.ToInt(allDetails[details.Name]["AvgPrice"]) + cast.ToInt(details.AmountPerBuy)) / cast.ToInt(allDetails[details.Name]["AvgCount"]))
			if allDetails[details.Name]["Symbol"] == "" {
				allDetails[details.Name]["Symbol"] = details.Symbol
			}
		} else {
			checkName[details.Name] = true
			allDetails[details.Name] = make(map[string]string)
			allDetails[details.Name]["Quantity"] = details.RemainingCount
			allDetails[details.Name]["AvgPrice"] = details.AmountPerBuy
			allDetails[details.Name]["InvestedAmount"] = cast.ToString(cast.ToInt(details.RemainingCount) * cast.ToInt(details.AmountPerBuy))
			allDetails[details.Name]["Symbol"] = details.Symbol
			allDetails[details.Name]["AvgCount"] = "1"
		}
	}

	for name, details := range allDetails {
		singleDetails := domain.HoldingDetails{}
		var symbol string
		singleDetails.Name = name
		for key, value := range details {
			switch key {
			case "Quantity":
				singleDetails.Quantity = value
			case "AvgPrice":
				singleDetails.AvgPrice = value
			case "InvestedAmount":
				singleDetails.InvestedAmount = value
				investeds += cast.ToInt(value)
			case "Symbol":
				symbol = value
			}
		}
		if symbol != "" {
			switch category {
			case utilities.ConstStock:
				stockDetails, err := s.investment.StockNamePersistence.GetYpStockNameBySymbol(symbol)
				if err != nil {
					return
				}
				singleDetails.MktPrice = cast.ToString(stockDetails.Price)
				singleDetails.CurrentAmount = cast.ToString(cast.ToInt(stockDetails.Price) * cast.ToInt(singleDetails.Quantity))
				singleDetails.TotalReturns = cast.ToString(cast.ToInt(singleDetails.CurrentAmount) - cast.ToInt(singleDetails.InvestedAmount))
				currents = currents + (cast.ToInt(stockDetails.Price) * cast.ToInt(singleDetails.Quantity))
				totalReturns = totalReturns + (cast.ToInt(singleDetails.CurrentAmount) - cast.ToInt(singleDetails.InvestedAmount))
			}
		}
		holdingDetails = append(holdingDetails, singleDetails)
	}

	if investeds > 0 {
		invested = cast.ToString(investeds)
	}
	if currents > 0 {
		currents = cast.ToInt(currents)
	}
	if totalReturns > 0 {
		totalReturn = cast.ToString(totalReturns)
	}
	return
}
