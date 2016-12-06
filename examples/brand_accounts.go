/** filename: brand_accounts.go
 *
 *	example for getting child accounts of given brand master.
 *
 * vim: set ts=4 sw=4:
 */

package main

import (
	"fmt"
	"github.com/hardenedlayer/hardenedlayer-go/util"
	"github.com/hardenedlayer/hardenedlayer-go/customer"
)

func main() {
	var customers []*customer.Customer

	customers = append(customers, customer.New())
	customers[len(customers)-1].SetLogger(util.GetLogger(util.DEBUG))

	fmt.Printf("\n# Brand, Get child accounts...\n")
	for i, customer := range customers {
		fmt.Printf("CUSTOMER %d: %s\n", i, customer.UserName)
		accounts := customer.ChildAccounts()
		for _, account := range accounts {
			fmt.Printf(" [%d] %s\n", *account.Id, *account.CompanyName)
		}
	}

	fmt.Printf("\n# Brand, Open Tickets...\n")
	for i, customer := range customers {
		fmt.Printf("CUSTOMER %d: %s\n", i, customer.UserName)
		tickets := customer.GetOpenTickets()
		for _, ticket := range tickets {
			fmt.Printf("Ticket: %v %v %v %v %v\n",
					*ticket.Id,
					*ticket.AccountId,
					*ticket.GroupId,
					*ticket.Title,
					*ticket.TotalUpdateCount,
				)
		}
	}
	fmt.Printf("\nEND --\n")
}
