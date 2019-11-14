package asaas_test

import (
	"os"
	"testing"

	"github.com/rafaeltokyo/asaas-go-sdk"
)

func TestBoleto(t *testing.T) {
	client := asaas.NewAsaasClient(os.Getenv("ASAAS_TOKEN"))

	paymentBoleto := asaas.PymentBoleto{
		Customer:          "cus_13bFHumeyglN",
		DueDate:           "12/12/2019",
		Value:             10,
		ExternalReference: "1111",
		Description:       "teste",
	}

	boletoResponse, errAPI, err := client.PaymentBoleto(paymentBoleto)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}

	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}

	if boletoResponse == nil {
		t.Error("boletoResponse is null")
		return
	}

}
