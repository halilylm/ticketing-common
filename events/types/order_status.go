package types

type OrderStatus int

const (
	// Created defines when the order has
	// been created but the ticket it is trying to
	// order has not been reserved
	Created OrderStatus = iota

	// Cancelled defines when the ticket the order
	// is trying to reserve has already been
	// reserved or when the user has cancelled the order.
	// the order expires before payment
	Cancelled

	// AwaitingPayment defines the order has successfully
	// reserved the ticket
	AwaitingPayment

	// Complete defines the situation when the order
	// has reserved the ticket and the user has provided
	// payment successfully
	Complete
)
