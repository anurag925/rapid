package scripts

import (
	"fmt"
	"github.com/anurag925/rapid/pkg/mailers"
	"github.com/anurag925/rapid/pkg/models"
)

func SendTransactionOtpMail() {
	err := mailers.NewOtpMailer().SendTransactionOtpMail(models.Account{
		Email: "dev@rampnow.io",
	}, 123456)
	fmt.Printf("error %+v", err)
}
