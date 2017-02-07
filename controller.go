package querylisttool

import (
	"net/http"
)

func QueryListRest(resp http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	var params = make(map[string]interface{})

	key := req.Form.Get("key")

	for k, _ := range req.Form {

		if k != "key" {
			params[k] = req.Form.Get(k)
		}
	}

	rets, err := QueryDataList(key, params)
	if err != nil {
		http.Error(resp, err.Error(), 500)
		return
	}

	resp.Write([]byte(rets))

}
