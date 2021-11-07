package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"aph-go-service-master/database"
	"aph-go-service-master/datastruct"
	"aph-go-service-master/logging"
	"aph-go-service-master/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	MLikeService(context.Context, string) string
}

type aphService struct{}

var ErrEmpty = errors.New("empty string")

func (aphService) MLikeService(_ context.Context, name string) string {

	return call_ServiceMLike(name)
}

func call_ServiceMLike(name string) string {

	jmlLike := database.SelectLike()
	Like := service.Mlike(name, jmlLike)

	return Like

}

func makeMLikeEndpoint(aph AphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.MLikeRequest)
		logging.Log(fmt.Sprintf("Like Request %s", req.NAME))
		v := aph.MLikeService(ctx, req.NAME)
		logging.Log(fmt.Sprintf("Response Final Message %s", v))
		return datastruct.MLikeResponse{v}, nil
	}
}

func decodeMLikeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.MLikeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	aph := aphService{}

	MLikeHandler := httptransport.NewServer(
		makeMLikeEndpoint(aph),
		decodeMLikeRequest,
		encodeResponse,
	)

	http.Handle("/MLike", MLikeHandler)
}

func GetInitDb() {
	database.InitDb()
}
