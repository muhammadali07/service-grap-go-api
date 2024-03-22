package repository

import (
	"time"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
)

func (r *Accountepository) TransactionCashDeposito(req models.TransactionDepositWithdraw) (err error) {
	err = r.db.Model(models.Account{}).Where("account_number = ?", req.AccountNumber).Update("balance", req.Amount).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error":      err.Error(),
			"payload":    req,
			"created_at": time.Time.GoString(time.Now()),
		}).Error("create new account data failed")
	}
	return
}

func (r *Accountepository) GetAccountBalance(req string) (response float64, err error) {
	var account models.Account
	err = r.db.Where("account_number = ?", req).First(&account).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"error":   err.Error(),
			"request": req,
		}).Error("query account balance data failed")
	}

	response = account.Balance
	return
}
