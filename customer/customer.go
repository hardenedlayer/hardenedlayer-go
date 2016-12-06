/*
 *
 *
 * vim: set ts=4 sw=4:
 */


package customer

import (
	"github.com/hardenedlayer/hardenedlayer-go/util"

	"github.com/softlayer/softlayer-go/session"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/datatypes"
)


type Customer struct {
	UserName string
	APIKey string

	AccountId int
	BrandId int
	IsReseller int
	CompanyName string
	Email string

	session *session.Session

	// cached data
	has_vms bool
	vms []datatypes.Virtual_Guest

	// for logging and error handling
	logger *util.Logger
	LastError string
}


/** brand related functions
 */

func (c *Customer) ChildAccounts() (accounts []datatypes.Account) {
	bra := c.brand()
	if bra.Session == nil {
		c.logger.Error("no brand session found for %v\n", c.AccountId)
		return
	}

	accounts, err := bra.Limit(10).GetAllOwnedAccounts()
	if err != nil {
		c.logger.APIError(err)
	}
	return
}

func (c *Customer) GetOpenTickets() (tickets []datatypes.Ticket) {
	bra := c.brand()
	if bra.Session == nil {
		c.logger.Error("no brand session found for %v\n", c.AccountId)
		return
	}

	tickets, err := bra.Limit(40).GetOpenTickets()
	if err != nil {
		c.logger.APIError(err)
	}
	return
}

func (c *Customer) brand() (brand services.Brand) {
	if c.IsReseller != 1 {	// f*ck
		c.logger.Warn("function for brand master was called by subaccount.\n")
	}

	brands, err := c.account().GetOwnedBrands()
	if err != nil {
		c.logger.APIError(err)
		return
	}
	if len(brands) != 1 {
		c.logger.Warn("cannot determined brand for reseller %v. (got %v)\n",
				c.AccountId, len(brands))
		return
	}
	b := brands[0]	// pick first one. need more?
	c.logger.Verb("Brand: %v %v\n", *b.Id, *b.Name)
	brand = services.GetBrandService(c.session).Id(*b.Id)
	return
}


/** account related functions
 */

/** cached wrapper for getting virtual server instances */
func (c *Customer) VMs() (vms []datatypes.Virtual_Guest) {
	if c.has_vms {
		c.logger.Verb("%v vms in cache. return from cache.\n", len(c.vms))
	} else {
		c.vms, c.has_vms = c.getVMs()
	}
	return c.vms
}

func (c *Customer) getVMs() (vms []datatypes.Virtual_Guest, status bool) {
	acc := c.account()
	c.logger.Debug("request vsi list for account %d...\n", c.AccountId)
	vms, err := acc.Mask("id;hostname;domain").Limit(10).GetVirtualGuests()
	if err != nil {
		c.logger.APIError(err)
		c.LastError = "cannot retrieve virtual guest list."
		return nil, false
	}
	return vms, true
}

func (c *Customer) account() (account services.Account) {
	account = services.GetAccountService(c.session)
	if c.AccountId < 1 {
		acc, err := account.GetObject()
		if err != nil {
			c.logger.APIError(err)
			return
		}
		c.AccountId = *acc.Id
		c.BrandId = *acc.BrandId
		c.IsReseller = *acc.IsReseller
		c.CompanyName = *acc.CompanyName
		c.Email = *acc.Email
		c.logger.Debug("Setup Account %d %s under %d\n",
				c.AccountId, c.CompanyName, c.BrandId)
	}
	return account
}


/** my own, customer structure
 *
 */

func (c *Customer) SetLogger(logger *util.Logger) {
	c.logger = logger
}

func (c *Customer) SetAPIDebug() {
	c.session.Debug = true
}

func (c *Customer) UnsetAPIDebug() {
	c.session.Debug = false
}

func New(args ...interface{}) *Customer {
	sess := session.New(args...)
	sess.Endpoint = "https://api.softlayer.com/rest/v3.1"
	logger := util.GetLogger(util.INFO)

	cust := &Customer {
		session: sess,
		UserName: sess.UserName,
		logger: logger,

		has_vms: false,
	}
	return cust
}
