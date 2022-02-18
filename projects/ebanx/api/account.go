package api

import (
	"ebanx/internal/dao"
	"ebanx/internal/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const (
	EventDeposit  = "deposit"
	EventWithdraw = "withdraw"
	EventTransfer = "transfer"
)

type Account struct {
	Base
}

type eventType struct {
	Type string `json:"type"`
}

type eventFunc func(*gin.Context, []byte)

func (a Account) Event(c *gin.Context) {
	bodyByte, _ := ioutil.ReadAll(c.Request.Body)
	eType := eventType{}
	if err := json.Unmarshal(bodyByte, &eType); err != nil {
		c.String(http.StatusInternalServerError, "invalid request body")
		return
	}

	eFuncMap := map[string]eventFunc{
		EventDeposit:  a.doDeposit,
		EventWithdraw: a.doWithdraw,
		EventTransfer: a.doTransfer,
	}

	if f, ok := eFuncMap[eType.Type]; !ok {
		c.String(http.StatusInternalServerError, "invalid event type")
	} else {
		f(c, bodyByte)
	}
}

func (a Account) Balance(c *gin.Context) {
	modelIns := model.NewAccountModel()
	id := c.Query("account_id")
	info := modelIns.Get(id)

	if info == nil {
		c.String(http.StatusNotFound, "0")
	} else {
		c.String(http.StatusOK, "%d", info.Balance)
	}
}

// deposit
type depositDto struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
}
type depositResponse struct {
	Destination dao.Account `json:"destination"`
}

func (d depositResponse) ToStr() string {
	str, _ := json.Marshal(d)
	return string(str)
}

func (a Account) doDeposit(c *gin.Context, reqBody []byte) {
	var err error
	dto := depositDto{}
	if err = json.Unmarshal(reqBody, &dto); err != nil {
		c.String(http.StatusInternalServerError, "invalid request body")
		return
	}

	modelIns := model.NewAccountModel()
	info := modelIns.Get(dto.Destination)
	if info == nil {
		info, err = modelIns.Create(dto.Destination, dto.Amount)
	} else {
		info, err = modelIns.Deposit(info, dto.Amount)
	}

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	response := depositResponse{
		Destination: *info,
	}
	c.String(http.StatusCreated, response.ToStr())
}

// withdrawDto
type withdrawDto struct {
	Type   string `json:"type"`
	Origin string `json:"origin"`
	Amount int    `json:"amount"`
}
type withdrawResponse struct {
	Origin dao.Account `json:"origin"`
}

func (d withdrawResponse) ToStr() string {
	str, _ := json.Marshal(d)
	return string(str)
}

func (a Account) doWithdraw(c *gin.Context, reqBody []byte) {
	var err error
	dto := withdrawDto{}
	if err = json.Unmarshal(reqBody, &dto); err != nil {
		c.String(http.StatusInternalServerError, "invalid request body")
		return
	}

	modelIns := model.NewAccountModel()
	info := modelIns.Get(dto.Origin)
	if info == nil {
		c.String(http.StatusNotFound, "0")
		return
	}
	info, err = modelIns.Withdraw(info, dto.Amount)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	response := withdrawResponse{
		Origin: *info,
	}
	c.String(http.StatusCreated, response.ToStr())
}

// transfer
type transferDto struct {
	Type        string `json:"type"`
	Origin      string `json:"origin"`
	Amount      int    `json:"amount"`
	Destination string `json:"destination"`
}
type transferResponse struct {
	Origin      dao.Account `json:"origin"`
	Destination dao.Account `json:"destination"`
}

func (d transferResponse) ToStr() string {
	str, _ := json.Marshal(d)
	return string(str)
}

func (a Account) doTransfer(c *gin.Context, reqBody []byte) {
	var err error
	dto := transferDto{}
	if err = json.Unmarshal(reqBody, &dto); err != nil {
		c.String(http.StatusInternalServerError, "invalid request body")
		return
	}

	modelIns := model.NewAccountModel()
	// origin info
	originInfo := modelIns.Get(dto.Origin)
	if originInfo == nil {
		c.String(http.StatusNotFound, "0")
		return
	}

	// dest info
	destInfo := modelIns.Get(dto.Destination)
	if destInfo == nil {
		destInfo, err = modelIns.Create(dto.Destination, 0)
	}

	originInfo, destInfo, err = modelIns.Transfer(originInfo, destInfo, dto.Amount)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	response := transferResponse{
		Origin:      *originInfo,
		Destination: *destInfo,
	}
	c.String(http.StatusCreated, response.ToStr())
}
