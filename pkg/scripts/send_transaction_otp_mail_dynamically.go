package scripts

import (
	"fmt"
	"github.com/anurag925/rapid/pkg/mailers"
	"github.com/anurag925/rapid/pkg/models"
	"reflect"
)

func SendTransactionOtpMailDynamically() {
	method := reflect.ValueOf(mailers.NewOtpMailer()).MethodByName("SendTransactionOtpMail")
	// Check if the method exists
	if method.IsValid() {
		// Call the method with an empty slice of arguments
		returnValues := method.Call([]reflect.Value{reflect.ValueOf(models.Account{Email: "dev@rampnow.io"}), reflect.ValueOf(123456)})
		if len(returnValues) > 0 {
			err := returnValues[0].Interface()
			if err != nil {
				// Handle the error
				fmt.Printf("Error: %v\n", err)
			}
		}
	} else {
		fmt.Printf("Method %s does not exist\n", "SendTransactionOtpMail")
	}
}
