package mexchttpmarket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kattana-io/mexc-golang-sdk/consts"
	"net/http"
)

// GetOpenOrders https://www.mexc.com/api-docs/spot-v3/spot-account-trade#current-open-orders
func (s *Service) GetOpenOrders(ctx context.Context, req *GetOpenOrdersRequest) ([]*GetOrderResponse, error) {
	params := make(map[string]string)

	params["timestamp"] = s.getTimestamp()

	if req.Symbol != nil {
		params["symbol"] = *req.Symbol
	}
	if req.RecvWindow != nil {
		params["recvWindow"] = fmt.Sprintf("%d", *req.RecvWindow)
	}

	res, err := s.client.SendRequest(ctx, http.MethodGet, consts.EndpointOpenOrders, params)
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

type GetOpenOrdersRequest struct {
	Symbol     *string `json:"symbol,omitempty"`
	RecvWindow *int64  `json:"recvWindow,omitempty"`
}
