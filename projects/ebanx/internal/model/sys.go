package model

import "ebanx/internal/dao"

type SysModel struct {
	dao dao.SysDao
}

func NewSysModel() SysModel {
	return SysModel{
		dao: dao.NewSysDao(),
	}
}

func (s SysModel) Reset() {
	s.dao.Reset()
}
