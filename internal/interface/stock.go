package interfaces

import stock "github.com/Sahil2k07/gRPC-GO/internal/generated/stock/proto"

type (
	StockService interface {
		stock.StockServiceServer
	}
)
