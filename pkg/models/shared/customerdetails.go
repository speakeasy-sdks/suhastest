// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CustomerDetails - The customer details that are necessary. Note that you can pass dummy details if your use case does not require the customer details.
type CustomerDetails struct {
	// Customer bank account. Required if you want to do a bank account check (TPV)
	CustomerBankAccountNumber *string `json:"customer_bank_account_number,omitempty"`
	// Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV)
	CustomerBankCode *string `json:"customer_bank_code,omitempty"`
	// Customer bank IFSC. Required if you want to do a bank account check (TPV)
	CustomerBankIfsc *string `json:"customer_bank_ifsc,omitempty"`
	// Customer email address.
	CustomerEmail *string `json:"customer_email,omitempty"`
	// A unique identifier for the customer. Use alphanumeric values only.
	CustomerID string `json:"customer_id"`
	// Customer phone number.
	CustomerPhone string `json:"customer_phone"`
}

func (o *CustomerDetails) GetCustomerBankAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.CustomerBankAccountNumber
}

func (o *CustomerDetails) GetCustomerBankCode() *string {
	if o == nil {
		return nil
	}
	return o.CustomerBankCode
}

func (o *CustomerDetails) GetCustomerBankIfsc() *string {
	if o == nil {
		return nil
	}
	return o.CustomerBankIfsc
}

func (o *CustomerDetails) GetCustomerEmail() *string {
	if o == nil {
		return nil
	}
	return o.CustomerEmail
}

func (o *CustomerDetails) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *CustomerDetails) GetCustomerPhone() string {
	if o == nil {
		return ""
	}
	return o.CustomerPhone
}
