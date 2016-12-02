
package customer

import (
	"fmt"

	"github.com/hardenedlayer/hardenedlayer-go/util"

	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/datatypes"
)


type Customer struct {
	UserName string
	APIKey string
	UderId int
	Session *session.Session
}

func New(args ...interface{}) *Customer {
	sess := session.New(args...)
	cust := &Customer {
		Session: sess,
	}
	return cust
}

func (c *Customer) VMs() (vms []datatypes.Virtual_Guest) {
	vms = c.getVMs()
	return vms
}

func (c *Customer) getVMs() (vms []datatypes.Virtual_Guest) {
	acc := services.GetAccountService(c.Session)
	vms, err := acc.Mask("id;hostname;domain").Limit(10).GetVirtualGuests()
	if err != nil {
		util.PrintError(err)
		return
	}
	fmt.Printf("")
	return vms
}

