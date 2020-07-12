package pagseg

import (
	"GoPagSeguro/pkg/core/net"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	PaymentId   = "CHAR_344a0907-8aa6-4b7a-943c-897383adf45f"
	ReferenceId = "76c35c0b-34d5-4ecc-af9d-3b2cecab033c"

	ChargeOkResponse          = fmt.Sprintf("{\n\t\"id\": \"%v\",\n\t\"reference_id\": \"%v\",\n\t\"status\": \"PAID\",\n\t\"created_at\": \"2019-09-06T13:43:45.588-03:00\",\n\t\"paid_at\": \"2019-09-06T13:43:45.588-03:00\",\n\t\"description\": \"Motivo da cobrança\",\n\t\"amount\": {\n\t\t\"value\": 1000,\n\t\t\"currency\": \"BRL\",\n\t\t\"summary\": {\n\t\t\t\"total\": 1000,\n\t\t\t\"paid\": 1000,\n\t\t\t\"refunded\": 0\n\t\t}\n\t},\n\t\"payment_response\": {\n\t\t\"code\": \"20000\",\n\t\t\"message\": \"SUCESSO\",\n\t\t\"reference\": \"1567788227697\"\n\t},\n\t\"payment_method\": {\n\t\t\"type\": \"CREDIT_CARD\",\n\t\t\"installments\": 1,\n\t\t\"capture\": true,\n\t\t\"card\": {\n\t\t\t\"brand\": \"VISA\",\n\t\t\t\"first_digits\": \"411111\",\n\t\t\t\"last_digits\": \"1111\",\n\t\t\t\"exp_month\": \"03\",\n\t\t\t\"exp_year\": \"2026\",\n\t\t\t\"holder\": {\n\t\t\t\t\"name\": \"Jose da Silva\"\n\t\t\t}\n\t\t}\n\t},\n\t\"links\": [{\n\t\t\t\"rel\": \"SELF\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/344a0907-8aa6-4b7a-943c-897383adf45f\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"GET\"\n\t\t},\n\t\t{\n\t\t\t\"rel\": \"CHARGE.CANCEL\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/344a0907-8aa6-4b7a-943c-897383adf45f/cancel\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"POST\"\n\t\t}\n\t],\n\t\"notification_urls\": [],\n\t\"metadata\": {}\n}", PaymentId, ReferenceId)
	ChargeFailResponse        = fmt.Sprintf("{\n\t\"id\": \"%v\",\n\t\"reference_id\": \"%v\",\n\t\"status\": \"DECLINED\",\n\t\"created_at\": \"2019-04-17T20:07:07.002-02\",\n\t\"description\": \"Motivo da cobrança\",\n\t\"amount\": {\n\t\t\"value\": 1000,\n\t\t\"currency\": \"BRL\",\n\t\t\"summary\": {\n\t\t\t\"total\": 1000,\n\t\t\t\"paid\": 0,\n\t\t\t\"refunded\": 0\n\t\t}\n\t},\n\t\"payment_response\": {\n\t\t\"code\": \"20001\",\n\t\t\"message\": \"CONTATE A CENTRAL DO SEU CARTÃO\",\n\t\t\"reference\": \"090400069049\"\n\t},\n\t\"payment_method\": {\n\t\t\"type\": \"CREDIT_CARD\",\n\t\t\"installments\": 1,\n\t\t\"capture\": true,\n\t\t\"card\": {\n\t\t\t\"brand\": \"VISA\",\n\t\t\t\"first_digits\": \"411111\",\n\t\t\t\"last_digits\": \"1111\",\n\t\t\t\"exp_month\": \"03\",\n\t\t\t\"exp_year\": \"2026\",\n\t\t\t\"holder\": {\n\t\t\t\t\"name\": \"Jose da Silva\"\n\t\t\t}\n\t\t}\n\t},\n\t\"notification_urls\": [],\n\t\"links\": [{\n\t\t\"rel\": \"SELF\",\n\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/CHAR_D32A01A9-92A6-4755-B21D-7B6A1291F7AD\",\n\t\t\"media\": \"application/json\",\n\t\t\"type\": \"GET\"\n\t}]\n}", PaymentId, ReferenceId)
	GetChargeOkResponse       = fmt.Sprintf("{\n\t\"id\": \"%v\",\n\t\"reference_id\": \"%v\",\n\t\"status\": \"AUTHORIZED\",\n\t\"created_at\": \"2019-04-17T20:07:07.002-02\",\n\t\"paid_at\": \"2019-04-17T20:07:07.002-02\",\n\t\"description\": \"Motivo da cobrança\",\n\t\"amount\": {\n\t\t\"value\": 1000,\n\t\t\"currency\": \"BRL\",\n\t\t\"summary\": {\n\t\t\t\"total\": 1000,\n\t\t\t\"paid\": 0,\n\t\t\t\"refunded\": 0\n\t\t}\n\t},\n\t\"payment_response\": {\n\t\t\"code\": 20000,\n\t\t\"message\": \"SUCESSO\",\n\t\t\"reference\": \"071200027526\"\n\t},\n\t\"payment_method\": {\n\t\t\"type\": \"CREDIT_CARD\",\n\t\t\"installments\": 1,\n\t\t\"capture\": false,\n\t\t\"card\": {\n\t\t\t\"brand\": \"VISA\",\n\t\t\t\"first_digits\": \"411111\",\n\t\t\t\"last_digits\": \"1111\",\n\t\t\t\"exp_month\": \"03\",\n\t\t\t\"exp_year\": \"2026\",\n\t\t\t\"holder\": {\n\t\t\t\t\"name\": \"Jose da Silva\"\n\t\t\t}\n\t\t}\n\t},\n\t\"notification_urls\": [\n\t\t\"https://yourserver.com/nas_erp/277be731-3b7c-4dac-8c4e-4c3f4a1fdc46/\"\n\t],\n\t\"links\": [{\n\t\t\t\"rel\": \"SELF\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/CHAR_A024DA52-C821-4A94-816F-803AD5307823\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"GET\"\n\t\t},\n\t\t{\n\t\t\t\"rel\": \"CHARGE.CAPTURE\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/CHAR_A024DA52-C821-4A94-816F-803AD5307823/capture\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"POST\"\n\t\t},\n\t\t{\n\t\t\t\"rel\": \"CHARGE.CANCEL\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/CHAR_A024DA52-C821-4A94-816F-803AD5307823/cancel\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"POST\"\n\t\t}\n\t]\n}", PaymentId, ReferenceId)
	CaptureOkResponse         = fmt.Sprintf("{\n\t\"id\": \"%v\",\n\t\"reference_id\": \"%v\",\n\t\"status\": \"PAID\",\n\t\"created_at\": \"2019-04-17T20:07:07.002-02\",\n\t\"paid_at\": \"2019-04-17T20:07:07.002-02\",\n\t\"description\": \"Motivo da cobrança\",\n\t\"amount\": {\n\t\t\"value\": 1000,\n\t\t\"currency\": \"BRL\",\n\t\t\"summary\": {\n\t\t\t\"total\": 1000,\n\t\t\t\"paid\": 1000,\n\t\t\t\"refunded\": 0\n\t\t}\n\t},\n\t\"payment_response\": {\n\t\t\"code\": 20000,\n\t\t\"message\": \"SUCESSO\",\n\t\t\"reference\": \"071200027526\"\n\t},\n\t\"payment_method\": {\n\t\t\"type\": \"CREDIT_CARD\",\n\t\t\"installments\": 1,\n\t\t\"capture\": false,\n\t\t\"card\": {\n\t\t\t\"brand\": \"VISA\",\n\t\t\t\"first_digits\": \"411111\",\n\t\t\t\"last_digits\": \"1111\",\n\t\t\t\"expiry_month\": \"03\",\n\t\t\t\"expiry_year\": \"2026\",\n\t\t\t\"holder\": {\n\t\t\t\t\"name\": \"Jose da Silva\"\n\t\t\t}\n\t\t}\n\t},\n\t\"notification_urls\": [\n\t\t\"https://yourserver.com/nas_ecommerce/277be731-3b7c-4dac-8c4e-4c3f4a1fdc46/\"\n\t],\n\t\"links\": [{\n\t\t\t\"rel\": \"SELF\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/CHAR_D32A01A9-92A6-4755-B21D-7B6A1291F7AD\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"GET\"\n\t\t},\n\t\t{\n\t\t\t\"rel\": \"CHARGE.CAPTURE\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/CHAR_D32A01A9-92A6-4755-B21D-7B6A1291F7AD/capture\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"POST\"\n\t\t},\n\t\t{\n\t\t\t\"rel\": \"CHARGE.CANCEL\",\n\t\t\t\"href\": \"https://sandbox.api.pagseguro.com/charges/CHAR_D32A01A9-92A6-4755-B21D-7B6A1291F7AD/cancel\",\n\t\t\t\"media\": \"application/json\",\n\t\t\t\"type\": \"POST\"\n\t\t}\n\t]\n}", PaymentId, ReferenceId)
	CancelAndRefundOkResponse = fmt.Sprintf("{\n\t\"id\": \"%v\",\n\t\"reference_id\": \"%v\",\n\t\"status\": \"CANCELED\",\n\t\"created_at\": \"2019-08-21T15:14:58.121-03:00\",\n\t\"paid_at\": \"2019-08-21T15:14:58.121-03:00\",\n\t\"description\": \"Motivo da cobranca\",\n\t\"amount\": {\n\t\t\"value\": 1000,\n\t\t\"currency\": \"BRL\",\n\t\t\"summary\": {\n\t\t\t\"total\": 1000,\n\t\t\t\"paid\": 1000,\n\t\t\t\"refunded\": 1000\n\t\t}\n\t},\n\t\"payment_response\": {\n\t\t\"code\": \"20000\",\n\t\t\"message\": \"SUCESSO\",\n\t\t\"reference\": \"1566411299393\"\n\t},\n\t\"payment_method\": {\n\t\t\"type\": \"CREDIT_CARD\",\n\t\t\"installments\": 1,\n\t\t\"capture\": true,\n\t\t\"card\": {\n\t\t\t\"brand\": \"VISA\",\n\t\t\t\"first_digits\": \"411111\",\n\t\t\t\"last_digits\": \"1111\",\n\t\t\t\"exp_month\": \"12\",\n\t\t\t\"exp_year\": \"2026\",\n\t\t\t\"holder\": {\n\t\t\t\t\"name\": \"Jose da Silva\"\n\t\t\t}\n\t\t}\n\t},\n\t\"links\": [],\n\t\"notification_urls\": [\n\t\t\"https://api.runscope.com/radar/inbound/f9e7bcbd-50dc-4821-8959-9854796d01c3\"\n\t],\n\t\"metadata\": {}\n}", PaymentId, ReferenceId)
)

type HttpRequesterMock struct {
	Headers map[string]string
}

func (r *HttpRequesterMock) SetHeaders(headers map[string]string) net.Requester {
	return r
}

func (_ HttpRequesterMock) DoRequest(r *http.Request) (*http.Response, error) {
	var body string
	statusCode := http.StatusCreated

	paths := strings.Split(r.URL.Path, "/")

	switch paths[len(paths)-1] {
	case "charges":
		body = ChargeOkResponse
	case "capture":
		body = CaptureOkResponse
	case "cancel":
		body = CancelAndRefundOkResponse
	default:
		body = GetChargeOkResponse
		statusCode = http.StatusOK
	}

	response := &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}
	return response, nil
}

func (_ HttpRequesterMock) ReadBody(body io.Reader) ([]byte, error) {
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
