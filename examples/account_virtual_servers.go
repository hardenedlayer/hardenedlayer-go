/** filename: account_virtual_servers.go
 *
 *	example for getting virtual servers of given account.
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
	//customers[len(customers)-1].SetAPIDebug()

	fmt.Printf("\n# Account, Get virtual server instances...\n")
	for i, customer := range customers {
		fmt.Printf("CUSTOMER %d: %s\n", i, customer.UserName)
		vms := customer.VMs()
		for _, vm := range vms {
			fmt.Printf(" [%d] %s.%s\n", *vm.Id, *vm.Hostname, *vm.Domain)
		}
		vms = customer.VMs()
		for _, vm := range vms {
			fmt.Printf(" [%d] %s.%s\n", *vm.Id, *vm.Hostname, *vm.Domain)
		}
	}
	fmt.Printf("\nEND --\n")
}
