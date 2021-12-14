// Code generated by xgen. DO NOT EDIT.

package v3

import (
	"github.com/flyhard/nordea2ynab/ofx"
)

// OFXRequestType is The OFX request root element for the SOAP Body.
type OFXRequestType struct {
	SIGNONMSGSRQV1     []*SignonRequestMessageSetV1              `xml:"SIGNONMSGSRQV1"`
	SIGNUPMSGSRQV1     []*SignupRequestMessageSetV1              `xml:"SIGNUPMSGSRQV1"`
	BANKMSGSRQV1       []*BankRequestMessageSetV1                `xml:"BANKMSGSRQV1"`
	CREDITCARDMSGSRQV1 []*CreditcardRequestMessageSetV1          `xml:"CREDITCARDMSGSRQV1"`
	INVSTMTMSGSRQV1    []*InvestmentStatementRequestMessageSetV1 `xml:"INVSTMTMSGSRQV1"`
	INTERXFERMSGSRQV1  []*InterTransferRequestMessageSetV1       `xml:"INTERXFERMSGSRQV1"`
	WIREXFERMSGSRQV1   []*WireTransferRequestMessageSetV1        `xml:"WIREXFERMSGSRQV1"`
	BILLPAYMSGSRQV1    []*BillPayRequestMessageSetV1             `xml:"BILLPAYMSGSRQV1"`
	EMAILMSGSRQV1      []*EmailRequestMessageSetV1               `xml:"EMAILMSGSRQV1"`
	IMAGEMSGSRQV1      []*ImageRequestMessageSetV1               `xml:"IMAGEMSGSRQV1"`
	LOANMSGSRQV1       []*LoanRequestMessageSetV1                `xml:"LOANMSGSRQV1"`
	SECLISTMSGSRQV1    []*SecurityListRequestMessageSetV1        `xml:"SECLISTMSGSRQV1"`
	PRESDIRMSGSRQV1    []*PresentmentDirRequestMessageSetV1      `xml:"PRESDIRMSGSRQV1"`
	PRESDLVMSGSRQV1    []*PresentmentDeliveryRequestMessageSetV1 `xml:"PRESDLVMSGSRQV1"`
	PROFMSGSRQV1       []*ProfileRequestMessageSetV1             `xml:"PROFMSGSRQV1"`
}

// OFXResponseType is The OFX response root element for the SOAP Body.
type OFXResponseType struct {
	SIGNONMSGSRSV1     []*SignonResponseMessageSetV1              `xml:"SIGNONMSGSRSV1"`
	SIGNUPMSGSRSV1     []*SignupResponseMessageSetV1              `xml:"SIGNUPMSGSRSV1"`
	BANKMSGSRSV1       []*BankResponseMessageSetV1                `xml:"BANKMSGSRSV1"`
	CREDITCARDMSGSRSV1 []*CreditcardResponseMessageSetV1          `xml:"CREDITCARDMSGSRSV1"`
	INVSTMTMSGSRSV1    []*InvestmentStatementResponseMessageSetV1 `xml:"INVSTMTMSGSRSV1"`
	INTERXFERMSGSRSV1  []*InterTransferResponseMessageSetV1       `xml:"INTERXFERMSGSRSV1"`
	WIREXFERMSGSRSV1   []*WireTransferResponseMessageSetV1        `xml:"WIREXFERMSGSRSV1"`
	BILLPAYMSGSRSV1    []*BillPayResponseMessageSetV1             `xml:"BILLPAYMSGSRSV1"`
	EMAILMSGSRSV1      []*EmailResponseMessageSetV1               `xml:"EMAILMSGSRSV1"`
	LOANMSGSRSV1       []*LoanResponseMessageSetV1                `xml:"LOANMSGSRSV1"`
	SECLISTMSGSRSV1    []*SecurityListResponseMessageSetV1        `xml:"SECLISTMSGSRSV1"`
	PRESDIRMSGSRSV1    []*PresentmentDirResponseMessageSetV1      `xml:"PRESDIRMSGSRSV1"`
	PRESDLVMSGSRSV1    []*PresentmentDeliveryResponseMessageSetV1 `xml:"PRESDLVMSGSRSV1"`
	PROFMSGSRSV1       []*ProfileResponseMessageSetV1             `xml:"PROFMSGSRSV1"`
}

// OFXRequest ...
type OFXRequest *OFXRequestType

// OFXResponse ...
type OFXResponse *OFXResponseType

// SONRQ ...
type SONRQ *ofx.SignonRequest

// SONRS ...
type SONRS *ofx.SignonResponse

// BankRequestMessageSetV1 is The OFX element "BANKMSGSRQV1" is of type "BankRequestMessageSetV1"
type BankRequestMessageSetV1 struct {
	STMTTRNRQ      []*ofx.StatementTransactionRequest      `xml:"STMTTRNRQ"`
	STMTENDTRNRQ   []*ofx.StatementEndTransactionRequest   `xml:"STMTENDTRNRQ"`
	INTRATRNRQ     []*ofx.IntraTransactionRequest          `xml:"INTRATRNRQ"`
	RECINTRATRNRQ  []*ofx.RecurringIntraTransactionRequest `xml:"RECINTRATRNRQ"`
	STPCHKTRNRQ    []*ofx.StopCheckTransactionRequest      `xml:"STPCHKTRNRQ"`
	BANKMAILTRNRQ  []*ofx.BankMailTransactionRequest       `xml:"BANKMAILTRNRQ"`
	BANKMAILSYNCRQ []*ofx.BankMailSyncRequest              `xml:"BANKMAILSYNCRQ"`
	STPCHKSYNCRQ   []*ofx.StopCheckSyncRequest             `xml:"STPCHKSYNCRQ"`
	INTRASYNCRQ    []*ofx.IntraSyncRequest                 `xml:"INTRASYNCRQ"`
	RECINTRASYNCRQ []*ofx.RecurringIntraSyncRequest        `xml:"RECINTRASYNCRQ"`
	*AbstractRequestMessageSet
}

// BankResponseMessageSetV1 is The OFX element "BANKMSGSRSV1" is of type "BankResponseMessageSetV1"
type BankResponseMessageSetV1 struct {
	STMTTRNRS      []*ofx.StatementTransactionResponse      `xml:"STMTTRNRS"`
	STMTENDTRNRS   []*ofx.StatementEndTransactionResponse   `xml:"STMTENDTRNRS"`
	INTRATRNRS     []*ofx.IntraTransactionResponse          `xml:"INTRATRNRS"`
	RECINTRATRNRS  []*ofx.RecurringIntraTransactionResponse `xml:"RECINTRATRNRS"`
	STPCHKTRNRS    []*ofx.StopCheckTransactionResponse      `xml:"STPCHKTRNRS"`
	BANKMAILTRNRS  []*ofx.BankMailTransactionResponse       `xml:"BANKMAILTRNRS"`
	BANKMAILSYNCRS []*ofx.BankMailSyncResponse              `xml:"BANKMAILSYNCRS"`
	STPCHKSYNCRS   []*ofx.StopCheckSyncResponse             `xml:"STPCHKSYNCRS"`
	INTRASYNCRS    []*ofx.IntraSyncResponse                 `xml:"INTRASYNCRS"`
	RECINTRASYNCRS []*ofx.RecurringIntraSyncResponse        `xml:"RECINTRASYNCRS"`
	*AbstractResponseMessageSet
}

// ImageRequestMessageSetV1 is The OFX element "IMAGEMSGSRQV1" is of type "ImageRequestMessageSetV1"
type ImageRequestMessageSetV1 struct {
	IMAGETRNRQ []*ofx.ImageTransactionRequest `xml:"IMAGETRNRQ"`
	*AbstractRequestMessageSet
}

// LoanRequestMessageSetV1 is The OFX element "LOANMSGSRQV1" is of type "LoanRequestMessageSetV1"
type LoanRequestMessageSetV1 struct {
	LOANSTMTTRNRQ    []*ofx.LoanStatementTransactionRequest    `xml:"LOANSTMTTRNRQ"`
	AMRTSTMTTRNRQ    []*ofx.AmortizationTransactionRequest     `xml:"AMRTSTMTTRNRQ"`
	LOANSTMTENDTRNRQ []*ofx.LoanStatementEndTransactionRequest `xml:"LOANSTMTENDTRNRQ"`
	LOANMAILTRNRQ    []*ofx.LoanMailTransactionRequest         `xml:"LOANMAILTRNRQ"`
	LOANMAILSYNCRQ   []*ofx.LoanMailSyncRequest                `xml:"LOANMAILSYNCRQ"`
	*AbstractRequestMessageSet
}

// LoanResponseMessageSetV1 is The OFX element "LOANMSGSRSV1" is of type "LoanResponseMessageSetV1"
type LoanResponseMessageSetV1 struct {
	LOANSTMTTRNRS    []*ofx.LoanStatementTransactionResponse    `xml:"LOANSTMTTRNRS"`
	AMRTSTMTTRNRS    []*ofx.AmortizationTransactionResponse     `xml:"AMRTSTMTTRNRS"`
	LOANSTMTENDTRNRS []*ofx.LoanStatementEndTransactionResponse `xml:"LOANSTMTENDTRNRS"`
	LOANMAILTRNRS    []*ofx.LoanMailTransactionResponse         `xml:"LOANMAILTRNRS"`
	LOANMAILSYNCRS   []*ofx.LoanMailSyncResponse                `xml:"LOANMAILSYNCRS"`
	*AbstractResponseMessageSet
}

// BillPayRequestMessageSetV1 is The OFX element "BILLPAYMSGSRQV1" is of type "BillPayRequestMessageSetV1"
type BillPayRequestMessageSetV1 struct {
	PAYEETRNRQ    []*ofx.PayeeTransactionRequest            `xml:"PAYEETRNRQ"`
	PAYEESYNCRQ   []*ofx.PayeeSyncRequest                   `xml:"PAYEESYNCRQ"`
	PMTTRNRQ      []*ofx.PaymentTransactionRequest          `xml:"PMTTRNRQ"`
	RECPMTTRNRQ   []*ofx.RecurringPaymentTransactionRequest `xml:"RECPMTTRNRQ"`
	PMTINQTRNRQ   []*ofx.PaymentInquiryTransactionRequest   `xml:"PMTINQTRNRQ"`
	PMTMAILTRNRQ  []*ofx.PaymentMailTransactionRequest      `xml:"PMTMAILTRNRQ"`
	PMTSYNCRQ     []*ofx.PaymentSyncRequest                 `xml:"PMTSYNCRQ"`
	RECPMTSYNCRQ  []*ofx.RecurringPaymentSyncRequest        `xml:"RECPMTSYNCRQ"`
	PMTMAILSYNCRQ []*ofx.PaymentMailSyncRequest             `xml:"PMTMAILSYNCRQ"`
	*AbstractRequestMessageSet
}

// BillPayResponseMessageSetV1 is The OFX element "BILLPAYMSGSRSV1" is of type "BillPayResponseMessageSetV1"
type BillPayResponseMessageSetV1 struct {
	PAYEETRNRS    []*ofx.PayeeTransactionResponse            `xml:"PAYEETRNRS"`
	PAYEESYNCRS   []*ofx.PayeeSyncResponse                   `xml:"PAYEESYNCRS"`
	PMTTRNRS      []*ofx.PaymentTransactionResponse          `xml:"PMTTRNRS"`
	RECPMTTRNRS   []*ofx.RecurringPaymentTransactionResponse `xml:"RECPMTTRNRS"`
	PMTINQTRNRS   []*ofx.PaymentInquiryTransactionResponse   `xml:"PMTINQTRNRS"`
	PMTMAILTRNRS  []*ofx.PaymentMailTransactionResponse      `xml:"PMTMAILTRNRS"`
	PMTSYNCRS     []*ofx.PaymentSyncResponse                 `xml:"PMTSYNCRS"`
	RECPMTSYNCRS  []*ofx.RecurringPaymentSyncResponse        `xml:"RECPMTSYNCRS"`
	PMTMAILSYNCRS []*ofx.PaymentMailSyncResponse             `xml:"PMTMAILSYNCRS"`
	*AbstractResponseMessageSet
}

// PresentmentDeliveryRequestMessageSetV1 is The OFX element "PRESDLVMSGSRQV1" is of type "PresentmentDeliveryRequestMessageSetV1"
type PresentmentDeliveryRequestMessageSetV1 struct {
	PRESLISTTRNRQ        []*ofx.PresentmentListTransactionRequest             `xml:"PRESLISTTRNRQ"`
	PRESGRPACCTINFOTRNRQ []*ofx.PresentmentGroupAccountInfoTransactionRequest `xml:"PRESGRPACCTINFOTRNRQ"`
	PRESDETAILTRNRQ      []*ofx.PresentmentDetailTransactionRequest           `xml:"PRESDETAILTRNRQ"`
	BILLTBLSTRUCTTRNRQ   []*ofx.BillTableStructureTransactionRequest          `xml:"BILLTBLSTRUCTTRNRQ"`
	PRESNOTIFYTRNRQ      []*ofx.PresentmentNotifyTransactionRequest           `xml:"PRESNOTIFYTRNRQ"`
	BILLSTATUSMODTRNRQ   []*ofx.BillStatusModTransactionRequest               `xml:"BILLSTATUSMODTRNRQ"`
	PRESMAILSYNCRQ       []*ofx.PresentmentMailSyncRequest                    `xml:"PRESMAILSYNCRQ"`
	PRESMAILTRNRQ        []*ofx.PresentmentMailTransactionRequest             `xml:"PRESMAILTRNRQ"`
	*AbstractRequestMessageSet
}

// PresentmentDeliveryResponseMessageSetV1 is The OFX element "PRESDLVMSGSRSV1" is of type "PresentmentDeliveryResponseMessageSetV1"
type PresentmentDeliveryResponseMessageSetV1 struct {
	PRESLISTTRNRS        []*ofx.PresentmentListTransactionResponse             `xml:"PRESLISTTRNRS"`
	PRESGRPACCTINFOTRNRS []*ofx.PresentmentGroupAccountInfoTransactionResponse `xml:"PRESGRPACCTINFOTRNRS"`
	PRESDETAILTRNRS      []*ofx.PresentmentDetailTransactionResponse           `xml:"PRESDETAILTRNRS"`
	BILLTBLSTRUCTTRNRS   []*ofx.BillTableStructureTransactionResponse          `xml:"BILLTBLSTRUCTTRNRS"`
	PRESNOTIFYTRNRS      []*ofx.PresentmentNotifyTransactionResponse           `xml:"PRESNOTIFYTRNRS"`
	BILLSTATUSMODTRNRS   []*ofx.BillStatusModTransactionResponse               `xml:"BILLSTATUSMODTRNRS"`
	PRESMAILSYNCRS       []*ofx.PresentmentMailSyncResponse                    `xml:"PRESMAILSYNCRS"`
	PRESMAILTRNRS        []*ofx.PresentmentMailTransactionResponse             `xml:"PRESMAILTRNRS"`
	*AbstractResponseMessageSet
}

// InvestmentStatementRequestMessageSetV1 is The OFX element "INVSTMTMSGSRQV1" is of type "InvestmentStatementRequestMessageSetV1"
type InvestmentStatementRequestMessageSetV1 struct {
	INVSTMTTRNRQ  []*ofx.InvestmentStatementTransactionRequest `xml:"INVSTMTTRNRQ"`
	INVMAILTRNRQ  []*ofx.InvestmentMailTransactionRequest      `xml:"INVMAILTRNRQ"`
	INVMAILSYNCRQ []*ofx.InvestmentMailSyncRequest             `xml:"INVMAILSYNCRQ"`
	INVSTMTENDRQ  []*ofx.InvestmentStatementEndRequest         `xml:"INVSTMTENDRQ"`
	*AbstractRequestMessageSet
}

// InvestmentStatementResponseMessageSetV1 is The OFX element "INVSTMTMSGSRSV1" is of type "InvestmentStatementResponseMessageSetV1"
type InvestmentStatementResponseMessageSetV1 struct {
	INVSTMTTRNRS  []*ofx.InvestmentStatementTransactionResponse `xml:"INVSTMTTRNRS"`
	INVMAILTRNRS  []*ofx.InvestmentMailTransactionResponse      `xml:"INVMAILTRNRS"`
	INVMAILSYNCRS []*ofx.InvestmentMailSyncResponse             `xml:"INVMAILSYNCRS"`
	INVSTMTENDRS  []*ofx.InvestmentStatementEndResponse         `xml:"INVSTMTENDRS"`
	*AbstractResponseMessageSet
}

// CreditcardRequestMessageSetV1 is The OFX element "CREDITCARDMSGSRQV1" is of type "CreditcardRequestMessageSetV1"
type CreditcardRequestMessageSetV1 struct {
	CCSTMTTRNRQ    []*ofx.CreditCardStatementTransactionRequest    `xml:"CCSTMTTRNRQ"`
	CCSTMTENDTRNRQ []*ofx.CreditCardStatementEndTransactionRequest `xml:"CCSTMTENDTRNRQ"`
	*AbstractRequestMessageSet
}

// CreditcardResponseMessageSetV1 is The OFX element "CREDITCARDMSGSRSV1" is of type "CreditcardResponseMessageSetV1"
type CreditcardResponseMessageSetV1 struct {
	CCSTMTTRNRS    []*ofx.CreditCardStatementTransactionResponse    `xml:"CCSTMTTRNRS"`
	CCSTMTENDTRNRS []*ofx.CreditCardStatementEndTransactionResponse `xml:"CCSTMTENDTRNRS"`
	*AbstractResponseMessageSet
}

// EmailRequestMessageSetV1 is The OFX element "EMAILMSGSRQV1" is of type "EmailRequestMessageSetV1"
type EmailRequestMessageSetV1 struct {
	MAILTRNRQ    []*ofx.MailTransactionRequest    `xml:"MAILTRNRQ"`
	MAILSYNCRQ   []*ofx.MailSyncRequest           `xml:"MAILSYNCRQ"`
	GETMIMETRNRQ []*ofx.GetMimeTransactionRequest `xml:"GETMIMETRNRQ"`
	*AbstractRequestMessageSet
}

// EmailResponseMessageSetV1 is The OFX element "EMAILMSGSRSV1" is of type "EmailResponseMessageSetV1"
type EmailResponseMessageSetV1 struct {
	MAILTRNRS    []*ofx.MailTransactionResponse    `xml:"MAILTRNRS"`
	MAILSYNCRS   []*ofx.MailSyncResponse           `xml:"MAILSYNCRS"`
	GETMIMETRNRS []*ofx.GetMimeTransactionResponse `xml:"GETMIMETRNRS"`
	*AbstractResponseMessageSet
}

// InterTransferRequestMessageSetV1 is The OFX element "INTERXFERMSGSRQV1" is of type "InterTransferRequestMessageSetV1"
type InterTransferRequestMessageSetV1 struct {
	INTERTRNRQ     []*ofx.InterTransactionRequest          `xml:"INTERTRNRQ"`
	RECINTERTRNRQ  []*ofx.RecurringInterTransactionRequest `xml:"RECINTERTRNRQ"`
	INTERSYNCRQ    []*ofx.InterSyncRequest                 `xml:"INTERSYNCRQ"`
	RECINTERSYNCRQ []*ofx.RecurringInterSyncRequest        `xml:"RECINTERSYNCRQ"`
	*AbstractRequestMessageSet
}

// InterTransferResponseMessageSetV1 is The OFX element "INTERXFERMSGSRSV1" is of type "InterTransferResponseMessageSetV1"
type InterTransferResponseMessageSetV1 struct {
	INTERTRNRS     []*ofx.InterTransactionResponse          `xml:"INTERTRNRS"`
	RECINTERTRNRS  []*ofx.RecurringInterTransactionResponse `xml:"RECINTERTRNRS"`
	INTERSYNCRS    []*ofx.InterSyncResponse                 `xml:"INTERSYNCRS"`
	RECINTERSYNCRS []*ofx.RecurringInterSyncResponse        `xml:"RECINTERSYNCRS"`
	*AbstractResponseMessageSet
}

// ProfileRequestMessageSetV1 is The OFX element "PROFMSGSRQV1" is of type "ProfileRequestMessageSetV1"
type ProfileRequestMessageSetV1 struct {
	PROFTRNRQ *ofx.ProfileTransactionRequest `xml:"PROFTRNRQ"`
	*AbstractRequestMessageSet
}

// ProfileResponseMessageSetV1 is The OFX element "PROFMSGSRSV1" is of type "ProfileResponseMessageSetV1"
type ProfileResponseMessageSetV1 struct {
	PROFTRNRS *ofx.ProfileTransactionResponse `xml:"PROFTRNRS"`
	*AbstractResponseMessageSet
}

// SecurityListRequestMessageSetV1 is The OFX element "SECLISTMSGSRQV1" is of type "SecurityListRequestMessageSetV1"
type SecurityListRequestMessageSetV1 struct {
	SECLISTTRNRQ []*ofx.SecurityListTransactionRequest `xml:"SECLISTTRNRQ"`
	*AbstractRequestMessageSet
}

// SecurityListResponseMessageSetV1 is The OFX element "SECLISTMSGSRSV1" is of type "SecurityListResponseMessageSetV1"
type SecurityListResponseMessageSetV1 struct {
	SECLISTTRNRS []*ofx.SecurityListTransactionResponse `xml:"SECLISTTRNRS"`
	SECLIST      *ofx.SecurityList                      `xml:"SECLIST"`
	*AbstractResponseMessageSet
}

// SignonRequestMessageSetV1 is The OFX element "SIGNONMSGSRQV1" is of type "SignonRequestMessageSetV1"
type SignonRequestMessageSetV1 struct {
	SONRQ          *ofx.SignonRequest               `xml:"SONRQ"`
	PINCHTRNRQ     *ofx.PinChangeTransactionRequest `xml:"PINCHTRNRQ"`
	CHALLENGETRNRQ *ofx.ChallengeTransactionRequest `xml:"CHALLENGETRNRQ"`
	*AbstractRequestMessageSet
}

// SignonResponseMessageSetV1 is The OFX element "SIGNONMSGSRSV1" is of type "SignonResponseMessageSetV1"
type SignonResponseMessageSetV1 struct {
	SONRS          *ofx.SignonResponse               `xml:"SONRS"`
	PINCHTRNRS     *ofx.PinChangeTransactionResponse `xml:"PINCHTRNRS"`
	CHALLENGETRNRS *ofx.ChallengeTransactionResponse `xml:"CHALLENGETRNRS"`
	*AbstractResponseMessageSet
}

// SignupRequestMessageSetV1 is The OFX element "SIGNUPMSGSRQV1" is of type "SignupRequestMessageSetV1"
type SignupRequestMessageSetV1 struct {
	ENROLLTRNRQ       []*ofx.EnrollTransactionRequest         `xml:"ENROLLTRNRQ"`
	ACCTINFOTRNRQ     []*ofx.AccountInfoTransactionRequest    `xml:"ACCTINFOTRNRQ"`
	USERINFOTRNRQ     []*ofx.UserInfoTransactionRequest       `xml:"USERINFOTRNRQ"`
	CHGUSERINFOTRNRQ  []*ofx.ChangeUserInfoTransactionRequest `xml:"CHGUSERINFOTRNRQ"`
	CHGUSERINFOSYNCRQ []*ofx.ChangeUserInfoSyncRequest        `xml:"CHGUSERINFOSYNCRQ"`
	ACCTTRNRQ         []*ofx.AccountTransactionRequest        `xml:"ACCTTRNRQ"`
	ACCTSYNCRQ        []*ofx.AccountSyncRequest               `xml:"ACCTSYNCRQ"`
	*AbstractRequestMessageSet
}

// SignupResponseMessageSetV1 is The OFX element "SIGNUPMSGSRSV1" is of type "SignupResponseMessageSetV1"
type SignupResponseMessageSetV1 struct {
	ENROLLTRNRS       []*ofx.EnrollTransactionResponse         `xml:"ENROLLTRNRS"`
	ACCTINFOTRNRS     []*ofx.AccountInfoTransactionResponse    `xml:"ACCTINFOTRNRS"`
	USERINFOTRNRS     []*ofx.UserInfoTransactionResponse       `xml:"USERINFOTRNRS"`
	CHGUSERINFOTRNRS  []*ofx.ChangeUserInfoTransactionResponse `xml:"CHGUSERINFOTRNRS"`
	CHGUSERINFOSYNCRS []*ofx.ChangeUserInfoSyncResponse        `xml:"CHGUSERINFOSYNCRS"`
	ACCTTRNRS         []*ofx.AccountTransactionResponse        `xml:"ACCTTRNRS"`
	ACCTSYNCRS        []*ofx.AccountSyncResponse               `xml:"ACCTSYNCRS"`
	*AbstractResponseMessageSet
}

// WireTransferRequestMessageSetV1 is The OFX element "WIREXFERMSGSRQV1" is of type "WireTransferRequestMessageSetV1"
type WireTransferRequestMessageSetV1 struct {
	WIRETRNRQ  []*ofx.WireTransactionRequest `xml:"WIRETRNRQ"`
	WIRESYNCRQ []*ofx.WireSyncRequest        `xml:"WIRESYNCRQ"`
	*AbstractRequestMessageSet
}

// WireTransferResponseMessageSetV1 is The OFX element "WIREXFERMSGSRSV1" is of type "WireTransferResponseMessageSetV1"
type WireTransferResponseMessageSetV1 struct {
	WIRETRNRS  []*ofx.WireTransactionResponse `xml:"WIRETRNRS"`
	WIRESYNCRS []*ofx.WireSyncResponse        `xml:"WIRESYNCRS"`
	*AbstractResponseMessageSet
}

// PresentmentDirRequestMessageSetV1 is The OFX element "PRESDIRMSGSRQV1" is of type "PresentmentDirRequestMessageSetV1"
type PresentmentDirRequestMessageSetV1 struct {
	FINDBILLERTRNRQ *ofx.FindBillerTransactionRequest `xml:"FINDBILLERTRNRQ"`
	*AbstractRequestMessageSet
}

// PresentmentDirResponseMessageSetV1 is The OFX element "PRESDIRMSGSRSV1" is of type "PresentmentDirResponseMessageSetV1"
type PresentmentDirResponseMessageSetV1 struct {
	FINDBILLERTRNRS *ofx.FindBillerTransactionResponse `xml:"FINDBILLERTRNRS"`
	*AbstractResponseMessageSet
}

// AbstractRequestMessageSet ...
type AbstractRequestMessageSet struct {
	*AbstractTopLevelMessageSet
}

// AbstractResponseMessageSet ...
type AbstractResponseMessageSet struct {
	*AbstractTopLevelMessageSet
}

// AbstractTopLevelMessageSet ...
type AbstractTopLevelMessageSet struct {
}
