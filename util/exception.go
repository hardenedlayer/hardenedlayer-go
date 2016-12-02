/**
 *
 */

package util

import (
	"fmt"
	"github.com/softlayer/softlayer-go/sl"
)

func PrintError(err error) {
	apiErr := err.(sl.Error)
	fmt.Printf(
		"--> Exception: %s (%d %s)\n",
		apiErr.Exception,
		apiErr.StatusCode,
		apiErr.Message)
}
