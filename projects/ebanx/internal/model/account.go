package model

import "ebanx/internal/dao"

type AccountModel struct {
	dao dao.AccountDao
}

func NewAccountModel() AccountModel {
	return AccountModel{
		dao: dao.NewAccountDao(),
	}
}

func (a AccountModel) Get(id string) *dao.Account {
	return a.dao.Get(id)
}

func (a AccountModel) Create(id string, amount int) (*dao.Account, error) {
	info := dao.Account{
		Id:      id,
		Balance: amount,
	}
	return &info, a.dao.Create(id, info)
}

func (a AccountModel) Deposit(account *dao.Account, amount int) (*dao.Account, error) {
	account.Balance += amount
	return account, a.dao.Update(account.Id, *account)
}

func (a AccountModel) Withdraw(account *dao.Account, amount int) (*dao.Account, error) {
	account.Balance -= amount
	return account, a.dao.Update(account.Id, *account)
}

func (a AccountModel) Transfer(originAccount *dao.Account, destAccount *dao.Account, amount int) (*dao.Account, *dao.Account, error) {
	originAccount.Balance -= amount
	err := a.dao.Update(originAccount.Id, *originAccount)

	destAccount.Balance += amount
	err = a.dao.Update(destAccount.Id, *destAccount)

	return originAccount, destAccount, err
}
