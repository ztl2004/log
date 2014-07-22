package handler

import (
	"fmt"
	"github.com/arkors/log/model"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateLog(r render.Render, res *http.Request) {
	fmt.Println("Enter Create Log Modle.......")
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Fail to read request body!"})
		log.Fatalf("Fail to read request body", err)
	}
	if len(data) == 0 && res.Method != "Get" {
		r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Data can't be nil!"})
		log.Fatalf("Data can't be nil", data)
	}

	fmt.Println("Enter Create Log Modle")
	redata := model.InsertLog(data)
	r.JSON(http.StatusOK, redata)
}
