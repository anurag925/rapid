package scripts

import (
	"fmt"
	"rapid/pkg/mailers"
	"rapid/pkg/models"
)

func SendTransactionOtpMail() {
	err := mailers.NewOtpMailer().SendTransactionOtpMail(models.Account{
		Email: "dev@rampnow.io",
	}, 123456)
	fmt.Printf("error %+v", err)
}
