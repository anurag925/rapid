package mailers

import "rapid/pkg/models"

type OtpMailer struct {
	From string
}

func NewOtpMailer() *OtpMailer {
	return &OtpMailer{From: "support@testmail.com"}
}

func (m *OtpMailer) SendTransactionOtpMail(a models.Account, otp int) (err error) {
	type Body struct {
		Otp int
	}
	err = call([]string{a.Email}, m.From, "OTP for transaction", "transaction_otp", Body{Otp: otp})
	return
}
