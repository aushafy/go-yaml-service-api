package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aushafy/go-yaml-service-api/models"
	"github.com/gin-gonic/gin"

	"github.com/go-git/go-git/v5"
)

var (
	errInvalidBody = errors.New("Invalid request body")
)

func Trigger(c *gin.Context) {
	/*

		this service would be clone/pull https://github.com/aushafy/yaml-file-test.git for the example and add phone number to values.yaml file inside that repo

	*/

	// /api/v1/trigger?run=true
	isRun := c.Query("run")

	if isRun == "true" {
		_, err := git.PlainClone("/tmp/repo", false, &git.CloneOptions{
			URL:      "https://github.com/aushafy/yaml-file-test.git",
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Printf("%s", err.Error())
		}

		var number models.Data

		if err := c.ShouldBindJSON(&number); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
			return
		}

		// dataNum := c.Query("number")
		// site	:= c.Query("site")

		// marshal yaml file from json file

		// write data number from struct to file

		c.String(http.StatusOK, "Ok")
	} else {
		c.String(http.StatusBadRequest, "Bad Request")
	}
}
