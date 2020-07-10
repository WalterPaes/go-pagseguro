package services

import (
	"GoPagSeguro/pkg/core/net"
	"GoPagSeguro/pkg/integrations/pagseg"
)

type Payment interface {
	Pay()
	Capture()
	Get()
	Cancel()
}

type payment struct {
	PagSeg pagseg.PagSeguro
}

func NewPayment() Payment {
	requester := net.NewHttpRequest() //http.NewHttpConnector("https://sandbox.api.pagseguro.com", map[string]string{})
	pagseguro := pagseg.NewPagSeguro(requester)
	return &payment{PagSeg: pagseguro}
}

func (p payment) Pay()     {}
func (p payment) Capture() {}
func (p payment) Get()     {}
func (p payment) Cancel()  {}
