package controllers

import (
	"fmt"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"

	"github.com/go-git/go-git/v5"
)

func Trigger(c *gin.Context) {
	// https://github.com/aushafy/yaml-file-test.git
	// wa628222211111
	isRun := c.Query("run") // /api/v1/trigger?run=true

	if isRun == "true" {
		_, err := git.PlainClone("/tmp", false, &git.CloneOptions{
			URL:      "https://github.com/aushafy/yaml-file-test.git",
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	}

	c.String(http.StatusNotImplemented, "Not Implement")
}
