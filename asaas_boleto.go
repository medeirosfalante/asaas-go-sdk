package asaas

import (
	"encoding/json"
	"fmt"
)

type Payment struct {
	Object            string          `json:"object"`
	ID                string          `json:"id"`
	DateCreated       string          `json:"dateCreated"`
	Customer          string          `json:"customer"`
	DueDate           string          `json:"dueDate"`
	Value             float32         `json:"value"`
	NetValue          float32         `json:"netValue"`
	BillingType       string          `json:"billingType"`
	Status            string          `json:"status"`
	Description       string          `json:"description"`
	ExternalReference string          `json:"externalReference"`
	OriginalDueDate   string          `json:"originalDueDate"`
	PaymentDate       string          `json:"paymentDate"`
	ClientPaymentDate string          `json:"clientPaymentDate"`
	InvoiceURL        string          `json:"invoiceUrl"`
	BankSlipURL       string          `json:"bankSlipUrl"`
	InvoiceNumber     string          `json:"invoiceNumber"`
	Discount          PaymentDiscount `json:"discount"`
	Fine              PaymentFine     `json:"fine"`
	Interest          PaymentInterest `json:"interest"`
	Deleted           bool            `json:"deleted"`
	PostalService     bool            `json:"postalService"`
	Anticipated       bool            `json:"anticipated"`
}

type PaymentDiscount struct {
	Value            float32 `json:"value"`
	DueDateLimitDays int32   `json:"dueDateLimitDays"`
}
type PaymentFine struct {
	Value float32 `json:"value"`
}

type PaymentInterest struct {
	Value float32 `json:"value"`
}

type PymentBoleto struct {
	Customer          string  `json:"customer"`
	DueDate           string  `json:"dueDate"`
	Value             float32 `json:"value"`
	ExternalReference string  `json:"externalReference"`
	Description       string  `json:"description"`
}

func (asaas *AsaasClient) PaymentBoleto(req PymentBoleto) (*Payment, *Error, error) {
	payment := Payment{
		Customer:          req.Customer,
		BillingType:       "boleto",
		DueDate:           req.DueDate,
		Value:             req.Value,
		Description:       req.Description,
		ExternalReference: req.ExternalReference,
		PostalService:     false,
	}
	data, _ := json.Marshal(payment)
	var response *Payment
	err, errAPI := asaas.Request("POST", fmt.Sprintf("payments"), data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
