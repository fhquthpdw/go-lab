package dao

type Account struct {
	Id      string `json:"id"`
	Balance int    `json:"balance"`
}

type AccountDao struct {
	BaseDao
}

func NewAccountDao() AccountDao {
	return AccountDao{
		BaseDao: BaseDao{
			db: db,
		},
	}
}

func (b BaseDao) Get(key string) *Account {
	accountInterface := b.db.Get(key)
	if accountInterface == nil {
		return nil
	}

	account := accountInterface.(Account)
	return &account
}

func (a AccountDao) List() map[string]Account {
	r := make(map[string]Account)
	list := a.db.List()
	for id, item := range list {
		r[id] = item.(Account)
	}

	return r
}
