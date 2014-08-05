package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/arkors/log/model"
	"github.com/martini-contrib/render"
)

func CreateLog(r render.Render, res *http.Request) {
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Fail to read request body!"})
		return
	}
	if len(data) == 0 && res.Method != "Get" {
		r.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Data can't be nil!"})
		return
	}
	model.InsertLog(data)
	r.JSON(http.StatusOK, nil)
	return
}
