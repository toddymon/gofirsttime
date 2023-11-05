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

func main() {
	r := gin.Default()
	r.GET("/accounts", getAccountsHandler)
	r.Run()
}
