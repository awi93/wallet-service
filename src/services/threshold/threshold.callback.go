package threshold

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/awi93/wallet-service/src/models/deposit"
	"github.com/awi93/wallet-service/src/models/threshold"
	"github.com/lovoo/goka"
)

func (t *service) ProcessCallback(ctx goka.Context, msg interface{}) {
	log.Println("Receiving New Message")
	var s *threshold.Threshold
	var deposit = msg.(*deposit.Deposit)
	if v := ctx.Value(); v == nil {
		s = t.composeThreshold(deposit)
		ctx.SetValue(s)
	} else {
		var latest = v.(*threshold.Threshold)
		if !latest.IsAboveThreshold {
			s = t.calculateThreshold(latest, deposit)
			ctx.SetValue(s)
		}
	}
}

func (t *service) composeThreshold(deposit *deposit.Deposit) *threshold.Threshold {
	var threshold = &threshold.Threshold{
		WalletId:            deposit.WalletId,
		TotalDeposit:        deposit.Amount,
		LatestDepositPeriod: t.getPeriodKey(deposit.DepositTime.AsTime()),
		IsAboveThreshold:    deposit.Amount > float32(t.config.GetFloat64("threshold.limit")),
	}
	return threshold
}

func (s *service) getPeriodWindowIndex(minute int, window int) int {
	return int(math.Floor(float64(minute)/float64(window))) + 1
}

func (t *service) getPeriodKey(depoTime time.Time) string {
	return fmt.Sprintf("%d-%d-%d-%d-%d", depoTime.Year(), depoTime.Month(), depoTime.Day(), depoTime.Hour(), t.getPeriodWindowIndex(depoTime.Minute(), t.config.GetInt("threshold.period_window")))
}

func (t *service) calculateThreshold(latestThreshold *threshold.Threshold, deposit *deposit.Deposit) *threshold.Threshold {
	depoTime := deposit.DepositTime.AsTime()
	depoPeriod := t.getPeriodKey(depoTime)
	if latestThreshold.LatestDepositPeriod == depoPeriod {
		latestThreshold.TotalDeposit += deposit.Amount
	} else {
		latestThreshold.TotalDeposit = deposit.Amount
	}
	latestThreshold.LatestDepositPeriod = depoPeriod
	latestThreshold.IsAboveThreshold = latestThreshold.TotalDeposit > float32(t.config.GetFloat64("threshold.limit"))
	return latestThreshold
}
