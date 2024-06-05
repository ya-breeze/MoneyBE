// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Geek Budget - OpenAPI 3.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Contact: ilya.korolev@outlook.com
 */

package goserver
import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	AccountsAPIService := NewAccountsAPIService()
	AccountsAPIController := NewAccountsAPIController(AccountsAPIService)

	AggregationsAPIService := NewAggregationsAPIService()
	AggregationsAPIController := NewAggregationsAPIController(AggregationsAPIService)

	AuthAPIService := NewAuthAPIService()
	AuthAPIController := NewAuthAPIController(AuthAPIService)

	BankImportersAPIService := NewBankImportersAPIService()
	BankImportersAPIController := NewBankImportersAPIController(BankImportersAPIService)

	CurrenciesAPIService := NewCurrenciesAPIService()
	CurrenciesAPIController := NewCurrenciesAPIController(CurrenciesAPIService)

	MatchersAPIService := NewMatchersAPIService()
	MatchersAPIController := NewMatchersAPIController(MatchersAPIService)

	NotificationsAPIService := NewNotificationsAPIService()
	NotificationsAPIController := NewNotificationsAPIController(NotificationsAPIService)

	TransactionsAPIService := NewTransactionsAPIService()
	TransactionsAPIController := NewTransactionsAPIController(TransactionsAPIService)

	UnprocessedTransactionsAPIService := NewUnprocessedTransactionsAPIService()
	UnprocessedTransactionsAPIController := NewUnprocessedTransactionsAPIController(UnprocessedTransactionsAPIService)

	UserAPIService := NewUserAPIService()
	UserAPIController := NewUserAPIController(UserAPIService)

	router := NewRouter(AccountsAPIController, AggregationsAPIController, AuthAPIController, BankImportersAPIController, CurrenciesAPIController, MatchersAPIController, NotificationsAPIController, TransactionsAPIController, UnprocessedTransactionsAPIController, UserAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}