package main

import (
	"fmt"
)

type CartItem struct {
	Code     string
	Quantity int
}

type Cart struct {
	Items []CartItem
}

type InventoryItem struct {
	Code      string
	UnitPrice float64
}

var inventory = []InventoryItem{
	{Code: "P001", UnitPrice: 10.0},
	{Code: "P002", UnitPrice: 10.0},
	{Code: "P003", UnitPrice: 10.0},
	{Code: "P004", UnitPrice: 10.0},
	{Code: "P005", UnitPrice: 10.0},
	{Code: "P006", UnitPrice: 10.0},
	{Code: "P007", UnitPrice: 10.0},
	{Code: "P008", UnitPrice: 10.0},
	{Code: "P009", UnitPrice: 10.0},
	{Code: "P010", UnitPrice: 10.0},
}

type PaymentMethod interface {
	Pay(amount float64)
	GetName() string
}

type CreditCardPayment struct {
	CardNumber  string
	ExpiryMonth string
	ExpiryYear  string
	Cvv         string
}

func (cc CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card.\n", amount)
}
func (cc CreditCardPayment) GetName() string {
	return "Credit Card"
}

type PayPalPayment struct {
	Email string
}

func (pp PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal.\n", amount)
}
func (pp PayPalPayment) GetName() string {
	return "PayPal"
}

type CryptoPayment struct {
	WalletAddress string
}

func (cp CryptoPayment) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Crypto Wallet.\n", amount)
}
func (cp CryptoPayment) GetName() string {
	return "Crypto"
}

// ---------- Decorator Pattern for Discounts ----------

type DiscountDecorator struct {
	Wrapped PaymentMethod
	DiscountPercent float64
}

func (d DiscountDecorator) Pay(amount float64) {
	discount := amount * d.DiscountPercent / 100
	finalAmount := amount - discount
	fmt.Printf("Applying %.2f%% discount: -$%.2f\n", d.DiscountPercent, discount)
	d.Wrapped.Pay(finalAmount)
}

func (d DiscountDecorator) GetName() string {
	return d.Wrapped.GetName() + " with Discount"
}


type PaymentProcessor struct {
	Method PaymentMethod
}

func (p *PaymentProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment with method: %s\n", p.Method.GetName())
	p.Method.Pay(amount)
}


func getPrice(code string) float64 {
	for _, item := range inventory {
		if item.Code == code {
			return item.UnitPrice
		}
	}
	return 0
}

func calculateCartTotal(cart Cart) float64 {
	total := 0.0
	for _, item := range cart.Items {
		total += float64(item.Quantity) * getPrice(item.Code)
	}
	return total
}


func main() {
	cart := Cart{
		Items: []CartItem{
			{Code: "P001", Quantity: 2},
			{Code: "P010", Quantity: 1},
		},
	}

	totalAmountToPay := calculateCartTotal(cart)
	fmt.Printf("Cart total: $%.2f\n", totalAmountToPay)

	cc := CreditCardPayment{
		CardNumber:  "1234567812345678",
		ExpiryMonth: "12",
		ExpiryYear:  "2026",
		Cvv:         "123",
	}

	discountedMethod := DiscountDecorator{
		Wrapped: cc,
		DiscountPercent: 10,
	}

	processor := PaymentProcessor{
		Method: discountedMethod,
	}
	processor.ProcessPayment(totalAmountToPay)
}
