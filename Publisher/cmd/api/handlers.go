package main

//go-json is a proxy of native json with better performance
import (
	"log"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/nelthaarion/snappfood-test-project-publisher/cmd/repositories"
	usecases "github.com/nelthaarion/snappfood-test-project-publisher/cmd/useCases"
)

func HandleOrder(res http.ResponseWriter, req *http.Request) {
	resp := GenerateResponse()
	res.Header().Set("content-type", "application/json")

	if req.Method == "POST" {
		uc := usecases.NewUseCase()
		order := repositories.Order{}
		err := json.NewDecoder(req.Body).Decode(&order)
		if err != nil {
			log.Println(err.Error())
			resp.SetError(true).SetMessage("bad request").Send(res, 400)
			return
		}
		err = uc.SetOrder(order).SendToRedis()
		if err != nil {
			log.Println(err)
			resp.SetError(true).SetMessage("something went wrong").Send(res, 500)
			return
		}
		resp.SetError(false).SetData(order).SetMessage("order created successfully").Send(res, 201)
	} else {
		resp.SetError(true).SetMessage("method not allowed").Send(res, 405)
	}
}
