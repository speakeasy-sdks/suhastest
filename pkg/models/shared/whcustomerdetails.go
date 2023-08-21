// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WHcustomerDetails struct {
	CustomerEmail *string `json:"customer_email,omitempty"`
	CustomerID    *string `json:"customer_id,omitempty"`
	CustomerName  *string `json:"customer_name,omitempty"`
	CustomerPhone *string `json:"customer_phone,omitempty"`
}

func (o *WHcustomerDetails) GetCustomerEmail() *string {
	if o == nil {
		return nil
	}
	return o.CustomerEmail
}

func (o *WHcustomerDetails) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *WHcustomerDetails) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *WHcustomerDetails) GetCustomerPhone() *string {
	if o == nil {
		return nil
	}
	return o.CustomerPhone
}
