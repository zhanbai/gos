package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(w http.ResponseWriter, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	rs, err := json.Marshal(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rs)
}
