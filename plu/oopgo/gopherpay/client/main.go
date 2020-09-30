package main

import "../payment"

func main() {

	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	option.ProcessPayment(500)

	option = payment.CreateCashAccount()
	option.ProcessPayment(500)

	// 	credit := payment.CreateCreditAccount(
	// 		"Debra Ingram",
	// 		"1111-2222-3333-4444",
	// 		5,
	// 		2021,
	// 		123)
	// 	fmt.Println(credit.CardNumber())

	// 	fmt.Printf("Owner name: %v\n", credit.OwnerName())
	// 	fmt.Printf("Owner name: %v\n", credit.CardNumber())
	// 	fmt.Println("Attempting to set card number")
	// 	err := credit.SetCardNumber(")(olvler")
	// 	if err != nil {
	// 		fmt.Printf("NO WOOK\n", err)
	// 	}
	// 	fmt.Printf("\nDOLLAH CREDIT: %v\n", credit.AvailableCredit())
}
