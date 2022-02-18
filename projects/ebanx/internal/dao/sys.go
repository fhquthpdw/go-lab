package dao

type SysDao struct {
	BaseDao
}

func NewSysDao() SysDao {
	return SysDao{
		BaseDao: BaseDao{
			db: db,
		},
	}
}

func (s SysDao) Reset() {
	accountDao := NewAccountDao()
	accountList := accountDao.List()

	for k, _ := range accountList {
		s.db.Delete(k)
	}
}
