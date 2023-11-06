package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	CID    int    `json:"cid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var accounts = map[int]*Account{
	1: &Account{CID: 1, Name: "Toddymon", Status: "active"},
}

func getAccountsHandler(c *gin.Context) {
	items := []*Account{}
	for _, item := range accounts {
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items)
}

func createAccountHandler(c *gin.Context) {
	acc := Account{}

	err := c.ShouldBindJSON(&acc)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	index := len(accounts)
	index++
	cid := index
	acc.CID = cid

	accounts[cid] = &acc

	c.JSON(http.StatusCreated, "Account Created.")
}

func main() {
	r := gin.Default()
	r.GET("/accounts", getAccountsHandler)
	r.POST("/accounts", createAccountHandler)
	r.Run(":2023")
}
