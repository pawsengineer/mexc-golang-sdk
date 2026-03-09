package mexchttpmarket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kattana-io/mexc-golang-sdk/consts"
	"net/http"
)

// CancelAllOpenOrders https://www.mexc.com/api-docs/spot-v3/spot-account-trade#cancel-all-open-orders-on-a-symbol
func (s *Service) CancelAllOpenOrders(ctx context.Context, req *CancelAllOpenOrdersRequest) ([]*GetOrderResponse, error) {
	params := make(map[string]string)

	params["symbol"] = req.Symbol
	params["timestamp"] = s.getTimestamp()

	if req.RecvWindow != nil {
		params["recvWindow"] = fmt.Sprintf("%d", *req.RecvWindow)
	}

	res, err := s.client.SendRequest(ctx, http.MethodDelete, consts.EndpointOpenOrders, params)
	if err != nil {
		return nil, err
	}

	var orders []*GetOrderResponse
	err = json.Unmarshal(res, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

type CancelAllOpenOrdersRequest struct {
	Symbol     string `json:"symbol"`
	RecvWindow *int64 `json:"recvWindow,omitempty"`
}
