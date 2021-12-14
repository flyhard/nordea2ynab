// Code generated by xgen. DO NOT EDIT.

package ofx

// OFX is The OFX element "OFX" is of type "OFX"
type OFX struct {
	SIGNONMSGSRQV1     *SignonRequestMessageSetV1               `xml:"SIGNONMSGSRQV1"`
	SIGNUPMSGSRQV1     *SignupRequestMessageSetV1               `xml:"SIGNUPMSGSRQV1"`
	BANKMSGSRQV1       *BankRequestMessageSetV1                 `xml:"BANKMSGSRQV1"`
	CREDITCARDMSGSRQV1 *CreditcardRequestMessageSetV1           `xml:"CREDITCARDMSGSRQV1"`
	INVSTMTMSGSRQV1    *InvestmentStatementRequestMessageSetV1  `xml:"INVSTMTMSGSRQV1"`
	INTERXFERMSGSRQV1  *InterTransferRequestMessageSetV1        `xml:"INTERXFERMSGSRQV1"`
	WIREXFERMSGSRQV1   *WireTransferRequestMessageSetV1         `xml:"WIREXFERMSGSRQV1"`
	BILLPAYMSGSRQV1    *BillPayRequestMessageSetV1              `xml:"BILLPAYMSGSRQV1"`
	EMAILMSGSRQV1      *EmailRequestMessageSetV1                `xml:"EMAILMSGSRQV1"`
	IMAGEMSGSRQV1      *ImageRequestMessageSetV1                `xml:"IMAGEMSGSRQV1"`
	LOANMSGSRQV1       *LoanRequestMessageSetV1                 `xml:"LOANMSGSRQV1"`
	SECLISTMSGSRQV1    *SecurityListRequestMessageSetV1         `xml:"SECLISTMSGSRQV1"`
	PRESDIRMSGSRQV1    *PresentmentDirRequestMessageSetV1       `xml:"PRESDIRMSGSRQV1"`
	PRESDLVMSGSRQV1    *PresentmentDeliveryRequestMessageSetV1  `xml:"PRESDLVMSGSRQV1"`
	PROFMSGSRQV1       *ProfileRequestMessageSetV1              `xml:"PROFMSGSRQV1"`
	SIGNONMSGSRSV1     *SignonResponseMessageSetV1              `xml:"SIGNONMSGSRSV1"`
	SIGNUPMSGSRSV1     *SignupResponseMessageSetV1              `xml:"SIGNUPMSGSRSV1"`
	BANKMSGSRSV1       *BankResponseMessageSetV1                `xml:"BANKMSGSRSV1"`
	CREDITCARDMSGSRSV1 *CreditcardResponseMessageSetV1          `xml:"CREDITCARDMSGSRSV1"`
	INVSTMTMSGSRSV1    *InvestmentStatementResponseMessageSetV1 `xml:"INVSTMTMSGSRSV1"`
	INTERXFERMSGSRSV1  *InterTransferResponseMessageSetV1       `xml:"INTERXFERMSGSRSV1"`
	WIREXFERMSGSRSV1   *WireTransferResponseMessageSetV1        `xml:"WIREXFERMSGSRSV1"`
	BILLPAYMSGSRSV1    *BillPayResponseMessageSetV1             `xml:"BILLPAYMSGSRSV1"`
	EMAILMSGSRSV1      *EmailResponseMessageSetV1               `xml:"EMAILMSGSRSV1"`
	LOANMSGSRSV1       *LoanResponseMessageSetV1                `xml:"LOANMSGSRSV1"`
	SECLISTMSGSRSV1    *SecurityListResponseMessageSetV1        `xml:"SECLISTMSGSRSV1"`
	PRESDIRMSGSRSV1    *PresentmentDirResponseMessageSetV1      `xml:"PRESDIRMSGSRSV1"`
	PRESDLVMSGSRSV1    *PresentmentDeliveryResponseMessageSetV1 `xml:"PRESDLVMSGSRSV1"`
	PROFMSGSRSV1       *ProfileResponseMessageSetV1             `xml:"PROFMSGSRSV1"`
}

// BankRequestMessageSetV1 is The OFX element "BANKMSGSRQV1" is of type "BankRequestMessageSetV1"
type BankRequestMessageSetV1 struct {
	STMTTRNRQ      []*StatementTransactionRequest      `xml:"STMTTRNRQ"`
	STMTENDTRNRQ   []*StatementEndTransactionRequest   `xml:"STMTENDTRNRQ"`
	INTRATRNRQ     []*IntraTransactionRequest          `xml:"INTRATRNRQ"`
	RECINTRATRNRQ  []*RecurringIntraTransactionRequest `xml:"RECINTRATRNRQ"`
	STPCHKTRNRQ    []*StopCheckTransactionRequest      `xml:"STPCHKTRNRQ"`
	BANKMAILTRNRQ  []*BankMailTransactionRequest       `xml:"BANKMAILTRNRQ"`
	BANKMAILSYNCRQ []*BankMailSyncRequest              `xml:"BANKMAILSYNCRQ"`
	STPCHKSYNCRQ   []*StopCheckSyncRequest             `xml:"STPCHKSYNCRQ"`
	INTRASYNCRQ    []*IntraSyncRequest                 `xml:"INTRASYNCRQ"`
	RECINTRASYNCRQ []*RecurringIntraSyncRequest        `xml:"RECINTRASYNCRQ"`
	*AbstractRequestMessageSet
}

// BankResponseMessageSetV1 is The OFX element "BANKMSGSRSV1" is of type "BankResponseMessageSetV1"
type BankResponseMessageSetV1 struct {
	STMTTRNRS      []*StatementTransactionResponse      `xml:"STMTTRNRS"`
	STMTENDTRNRS   []*StatementEndTransactionResponse   `xml:"STMTENDTRNRS"`
	INTRATRNRS     []*IntraTransactionResponse          `xml:"INTRATRNRS"`
	RECINTRATRNRS  []*RecurringIntraTransactionResponse `xml:"RECINTRATRNRS"`
	STPCHKTRNRS    []*StopCheckTransactionResponse      `xml:"STPCHKTRNRS"`
	BANKMAILTRNRS  []*BankMailTransactionResponse       `xml:"BANKMAILTRNRS"`
	BANKMAILSYNCRS []*BankMailSyncResponse              `xml:"BANKMAILSYNCRS"`
	STPCHKSYNCRS   []*StopCheckSyncResponse             `xml:"STPCHKSYNCRS"`
	INTRASYNCRS    []*IntraSyncResponse                 `xml:"INTRASYNCRS"`
	RECINTRASYNCRS []*RecurringIntraSyncResponse        `xml:"RECINTRASYNCRS"`
	*AbstractResponseMessageSet
}

// ImageRequestMessageSetV1 is The OFX element "IMAGEMSGSRQV1" is of type "ImageRequestMessageSetV1"
type ImageRequestMessageSetV1 struct {
	IMAGETRNRQ []*ImageTransactionRequest `xml:"IMAGETRNRQ"`
	*AbstractRequestMessageSet
}

// LoanRequestMessageSetV1 is The OFX element "LOANMSGSRQV1" is of type "LoanRequestMessageSetV1"
type LoanRequestMessageSetV1 struct {
	LOANSTMTTRNRQ    []*LoanStatementTransactionRequest    `xml:"LOANSTMTTRNRQ"`
	AMRTSTMTTRNRQ    []*AmortizationTransactionRequest     `xml:"AMRTSTMTTRNRQ"`
	LOANSTMTENDTRNRQ []*LoanStatementEndTransactionRequest `xml:"LOANSTMTENDTRNRQ"`
	LOANMAILTRNRQ    []*LoanMailTransactionRequest         `xml:"LOANMAILTRNRQ"`
	LOANMAILSYNCRQ   []*LoanMailSyncRequest                `xml:"LOANMAILSYNCRQ"`
	*AbstractRequestMessageSet
}

// LoanResponseMessageSetV1 is The OFX element "LOANMSGSRSV1" is of type "LoanResponseMessageSetV1"
type LoanResponseMessageSetV1 struct {
	LOANSTMTTRNRS    []*LoanStatementTransactionResponse    `xml:"LOANSTMTTRNRS"`
	AMRTSTMTTRNRS    []*AmortizationTransactionResponse     `xml:"AMRTSTMTTRNRS"`
	LOANSTMTENDTRNRS []*LoanStatementEndTransactionResponse `xml:"LOANSTMTENDTRNRS"`
	LOANMAILTRNRS    []*LoanMailTransactionResponse         `xml:"LOANMAILTRNRS"`
	LOANMAILSYNCRS   []*LoanMailSyncResponse                `xml:"LOANMAILSYNCRS"`
	*AbstractResponseMessageSet
}

// BillPayRequestMessageSetV1 is The OFX element "BILLPAYMSGSRQV1" is of type "BillPayRequestMessageSetV1"
type BillPayRequestMessageSetV1 struct {
	PAYEETRNRQ    []*PayeeTransactionRequest            `xml:"PAYEETRNRQ"`
	PAYEESYNCRQ   []*PayeeSyncRequest                   `xml:"PAYEESYNCRQ"`
	PMTTRNRQ      []*PaymentTransactionRequest          `xml:"PMTTRNRQ"`
	RECPMTTRNRQ   []*RecurringPaymentTransactionRequest `xml:"RECPMTTRNRQ"`
	PMTINQTRNRQ   []*PaymentInquiryTransactionRequest   `xml:"PMTINQTRNRQ"`
	PMTMAILTRNRQ  []*PaymentMailTransactionRequest      `xml:"PMTMAILTRNRQ"`
	PMTSYNCRQ     []*PaymentSyncRequest                 `xml:"PMTSYNCRQ"`
	RECPMTSYNCRQ  []*RecurringPaymentSyncRequest        `xml:"RECPMTSYNCRQ"`
	PMTMAILSYNCRQ []*PaymentMailSyncRequest             `xml:"PMTMAILSYNCRQ"`
	*AbstractRequestMessageSet
}

// BillPayResponseMessageSetV1 is The OFX element "BILLPAYMSGSRSV1" is of type "BillPayResponseMessageSetV1"
type BillPayResponseMessageSetV1 struct {
	PAYEETRNRS    []*PayeeTransactionResponse            `xml:"PAYEETRNRS"`
	PAYEESYNCRS   []*PayeeSyncResponse                   `xml:"PAYEESYNCRS"`
	PMTTRNRS      []*PaymentTransactionResponse          `xml:"PMTTRNRS"`
	RECPMTTRNRS   []*RecurringPaymentTransactionResponse `xml:"RECPMTTRNRS"`
	PMTINQTRNRS   []*PaymentInquiryTransactionResponse   `xml:"PMTINQTRNRS"`
	PMTMAILTRNRS  []*PaymentMailTransactionResponse      `xml:"PMTMAILTRNRS"`
	PMTSYNCRS     []*PaymentSyncResponse                 `xml:"PMTSYNCRS"`
	RECPMTSYNCRS  []*RecurringPaymentSyncResponse        `xml:"RECPMTSYNCRS"`
	PMTMAILSYNCRS []*PaymentMailSyncResponse             `xml:"PMTMAILSYNCRS"`
	*AbstractResponseMessageSet
}

// PresentmentDeliveryRequestMessageSetV1 is The OFX element "PRESDLVMSGSRQV1" is of type "PresentmentDeliveryRequestMessageSetV1"
type PresentmentDeliveryRequestMessageSetV1 struct {
	PRESLISTTRNRQ        []*PresentmentListTransactionRequest             `xml:"PRESLISTTRNRQ"`
	PRESGRPACCTINFOTRNRQ []*PresentmentGroupAccountInfoTransactionRequest `xml:"PRESGRPACCTINFOTRNRQ"`
	PRESDETAILTRNRQ      []*PresentmentDetailTransactionRequest           `xml:"PRESDETAILTRNRQ"`
	BILLTBLSTRUCTTRNRQ   []*BillTableStructureTransactionRequest          `xml:"BILLTBLSTRUCTTRNRQ"`
	PRESNOTIFYTRNRQ      []*PresentmentNotifyTransactionRequest           `xml:"PRESNOTIFYTRNRQ"`
	BILLSTATUSMODTRNRQ   []*BillStatusModTransactionRequest               `xml:"BILLSTATUSMODTRNRQ"`
	PRESMAILSYNCRQ       []*PresentmentMailSyncRequest                    `xml:"PRESMAILSYNCRQ"`
	PRESMAILTRNRQ        []*PresentmentMailTransactionRequest             `xml:"PRESMAILTRNRQ"`
	*AbstractRequestMessageSet
}

// PresentmentDeliveryResponseMessageSetV1 is The OFX element "PRESDLVMSGSRSV1" is of type "PresentmentDeliveryResponseMessageSetV1"
type PresentmentDeliveryResponseMessageSetV1 struct {
	PRESLISTTRNRS        []*PresentmentListTransactionResponse             `xml:"PRESLISTTRNRS"`
	PRESGRPACCTINFOTRNRS []*PresentmentGroupAccountInfoTransactionResponse `xml:"PRESGRPACCTINFOTRNRS"`
	PRESDETAILTRNRS      []*PresentmentDetailTransactionResponse           `xml:"PRESDETAILTRNRS"`
	BILLTBLSTRUCTTRNRS   []*BillTableStructureTransactionResponse          `xml:"BILLTBLSTRUCTTRNRS"`
	PRESNOTIFYTRNRS      []*PresentmentNotifyTransactionResponse           `xml:"PRESNOTIFYTRNRS"`
	BILLSTATUSMODTRNRS   []*BillStatusModTransactionResponse               `xml:"BILLSTATUSMODTRNRS"`
	PRESMAILSYNCRS       []*PresentmentMailSyncResponse                    `xml:"PRESMAILSYNCRS"`
	PRESMAILTRNRS        []*PresentmentMailTransactionResponse             `xml:"PRESMAILTRNRS"`
	*AbstractResponseMessageSet
}

// InvestmentStatementRequestMessageSetV1 is The OFX element "INVSTMTMSGSRQV1" is of type "InvestmentStatementRequestMessageSetV1"
type InvestmentStatementRequestMessageSetV1 struct {
	INVSTMTTRNRQ    []*InvestmentStatementTransactionRequest    `xml:"INVSTMTTRNRQ"`
	INVSTMTENDTRNRQ []*InvestmentStatementEndTransactionRequest `xml:"INVSTMTENDTRNRQ"`
	INVMAILTRNRQ    []*InvestmentMailTransactionRequest         `xml:"INVMAILTRNRQ"`
	INVMAILSYNCRQ   []*InvestmentMailSyncRequest                `xml:"INVMAILSYNCRQ"`
	*AbstractRequestMessageSet
}

// InvestmentStatementResponseMessageSetV1 is The OFX element "INVSTMTMSGSRSV1" is of type "InvestmentStatementResponseMessageSetV1"
type InvestmentStatementResponseMessageSetV1 struct {
	INVSTMTTRNRS    []*InvestmentStatementTransactionResponse    `xml:"INVSTMTTRNRS"`
	INVSTMTENDTRNRS []*InvestmentStatementEndTransactionResponse `xml:"INVSTMTENDTRNRS"`
	INVMAILTRNRS    []*InvestmentMailTransactionResponse         `xml:"INVMAILTRNRS"`
	INVMAILSYNCRS   []*InvestmentMailSyncResponse                `xml:"INVMAILSYNCRS"`
	*AbstractResponseMessageSet
}

// CreditcardRequestMessageSetV1 is The OFX element "CREDITCARDMSGSRQV1" is of type "CreditcardRequestMessageSetV1"
type CreditcardRequestMessageSetV1 struct {
	CCSTMTTRNRQ    []*CreditCardStatementTransactionRequest    `xml:"CCSTMTTRNRQ"`
	CCSTMTENDTRNRQ []*CreditCardStatementEndTransactionRequest `xml:"CCSTMTENDTRNRQ"`
	*AbstractRequestMessageSet
}

// CreditcardResponseMessageSetV1 is The OFX element "CREDITCARDMSGSRSV1" is of type "CreditcardResponseMessageSetV1"
type CreditcardResponseMessageSetV1 struct {
	CCSTMTTRNRS    []*CreditCardStatementTransactionResponse    `xml:"CCSTMTTRNRS"`
	CCSTMTENDTRNRS []*CreditCardStatementEndTransactionResponse `xml:"CCSTMTENDTRNRS"`
	*AbstractResponseMessageSet
}

// EmailRequestMessageSetV1 is The OFX element "EMAILMSGSRQV1" is of type "EmailRequestMessageSetV1"
type EmailRequestMessageSetV1 struct {
	MAILTRNRQ    []*MailTransactionRequest    `xml:"MAILTRNRQ"`
	MAILSYNCRQ   []*MailSyncRequest           `xml:"MAILSYNCRQ"`
	GETMIMETRNRQ []*GetMimeTransactionRequest `xml:"GETMIMETRNRQ"`
	*AbstractRequestMessageSet
}

// EmailResponseMessageSetV1 is The OFX element "EMAILMSGSRSV1" is of type "EmailResponseMessageSetV1"
type EmailResponseMessageSetV1 struct {
	MAILTRNRS    []*MailTransactionResponse    `xml:"MAILTRNRS"`
	MAILSYNCRS   []*MailSyncResponse           `xml:"MAILSYNCRS"`
	GETMIMETRNRS []*GetMimeTransactionResponse `xml:"GETMIMETRNRS"`
	*AbstractResponseMessageSet
}

// InterTransferRequestMessageSetV1 is The OFX element "INTERXFERMSGSRQV1" is of type "InterTransferRequestMessageSetV1"
type InterTransferRequestMessageSetV1 struct {
	INTERTRNRQ     []*InterTransactionRequest          `xml:"INTERTRNRQ"`
	RECINTERTRNRQ  []*RecurringInterTransactionRequest `xml:"RECINTERTRNRQ"`
	INTERSYNCRQ    []*InterSyncRequest                 `xml:"INTERSYNCRQ"`
	RECINTERSYNCRQ []*RecurringInterSyncRequest        `xml:"RECINTERSYNCRQ"`
	*AbstractRequestMessageSet
}

// InterTransferResponseMessageSetV1 is The OFX element "INTERXFERMSGSRSV1" is of type "InterTransferResponseMessageSetV1"
type InterTransferResponseMessageSetV1 struct {
	INTERTRNRS     []*InterTransactionResponse          `xml:"INTERTRNRS"`
	RECINTERTRNRS  []*RecurringInterTransactionResponse `xml:"RECINTERTRNRS"`
	INTERSYNCRS    []*InterSyncResponse                 `xml:"INTERSYNCRS"`
	RECINTERSYNCRS []*RecurringInterSyncResponse        `xml:"RECINTERSYNCRS"`
	*AbstractResponseMessageSet
}

// ProfileRequestMessageSetV1 is The OFX element "PROFMSGSRQV1" is of type "ProfileRequestMessageSetV1"
type ProfileRequestMessageSetV1 struct {
	PROFTRNRQ *ProfileTransactionRequest `xml:"PROFTRNRQ"`
	*AbstractRequestMessageSet
}

// ProfileResponseMessageSetV1 is The OFX element "PROFMSGSRSV1" is of type "ProfileResponseMessageSetV1"
type ProfileResponseMessageSetV1 struct {
	PROFTRNRS *ProfileTransactionResponse `xml:"PROFTRNRS"`
	*AbstractResponseMessageSet
}

// SecurityListRequestMessageSetV1 is The OFX element "SECLISTMSGSRQV1" is of type "SecurityListRequestMessageSetV1"
type SecurityListRequestMessageSetV1 struct {
	SECLISTTRNRQ []*SecurityListTransactionRequest `xml:"SECLISTTRNRQ"`
	*AbstractRequestMessageSet
}

// SecurityListResponseMessageSetV1 is The OFX element "SECLISTMSGSRSV1" is of type "SecurityListResponseMessageSetV1"
type SecurityListResponseMessageSetV1 struct {
	SECLISTTRNRS []*SecurityListTransactionResponse `xml:"SECLISTTRNRS"`
	SECLIST      *SecurityList                      `xml:"SECLIST"`
	*AbstractResponseMessageSet
}

// SignonRequestMessageSetV1 is The OFX element "SIGNONMSGSRQV1" is of type "SignonRequestMessageSetV1"
type SignonRequestMessageSetV1 struct {
	SONRQ          *SignonRequest               `xml:"SONRQ"`
	PINCHTRNRQ     *PinChangeTransactionRequest `xml:"PINCHTRNRQ"`
	CHALLENGETRNRQ *ChallengeTransactionRequest `xml:"CHALLENGETRNRQ"`
	*AbstractRequestMessageSet
}

// SignonResponseMessageSetV1 is The OFX element "SIGNONMSGSRSV1" is of type "SignonResponseMessageSetV1"
type SignonResponseMessageSetV1 struct {
	SONRS          *SignonResponse               `xml:"SONRS"`
	PINCHTRNRS     *PinChangeTransactionResponse `xml:"PINCHTRNRS"`
	CHALLENGETRNRS *ChallengeTransactionResponse `xml:"CHALLENGETRNRS"`
	*AbstractResponseMessageSet
}

// SignupRequestMessageSetV1 is The OFX element "SIGNUPMSGSRQV1" is of type "SignupRequestMessageSetV1"
type SignupRequestMessageSetV1 struct {
	ENROLLTRNRQ       []*EnrollTransactionRequest         `xml:"ENROLLTRNRQ"`
	ACCTINFOTRNRQ     []*AccountInfoTransactionRequest    `xml:"ACCTINFOTRNRQ"`
	USERINFOTRNRQ     []*UserInfoTransactionRequest       `xml:"USERINFOTRNRQ"`
	CHGUSERINFOTRNRQ  []*ChangeUserInfoTransactionRequest `xml:"CHGUSERINFOTRNRQ"`
	CHGUSERINFOSYNCRQ []*ChangeUserInfoSyncRequest        `xml:"CHGUSERINFOSYNCRQ"`
	ACCTTRNRQ         []*AccountTransactionRequest        `xml:"ACCTTRNRQ"`
	ACCTSYNCRQ        []*AccountSyncRequest               `xml:"ACCTSYNCRQ"`
	*AbstractRequestMessageSet
}

// SignupResponseMessageSetV1 is The OFX element "SIGNUPMSGSRSV1" is of type "SignupResponseMessageSetV1"
type SignupResponseMessageSetV1 struct {
	ENROLLTRNRS       []*EnrollTransactionResponse         `xml:"ENROLLTRNRS"`
	ACCTINFOTRNRS     []*AccountInfoTransactionResponse    `xml:"ACCTINFOTRNRS"`
	USERINFOTRNRS     []*UserInfoTransactionResponse       `xml:"USERINFOTRNRS"`
	CHGUSERINFOTRNRS  []*ChangeUserInfoTransactionResponse `xml:"CHGUSERINFOTRNRS"`
	CHGUSERINFOSYNCRS []*ChangeUserInfoSyncResponse        `xml:"CHGUSERINFOSYNCRS"`
	ACCTTRNRS         []*AccountTransactionResponse        `xml:"ACCTTRNRS"`
	ACCTSYNCRS        []*AccountSyncResponse               `xml:"ACCTSYNCRS"`
	*AbstractResponseMessageSet
}

// WireTransferRequestMessageSetV1 is The OFX element "WIREXFERMSGSRQV1" is of type "WireTransferRequestMessageSetV1"
type WireTransferRequestMessageSetV1 struct {
	WIRETRNRQ  []*WireTransactionRequest `xml:"WIRETRNRQ"`
	WIRESYNCRQ []*WireSyncRequest        `xml:"WIRESYNCRQ"`
	*AbstractRequestMessageSet
}

// WireTransferResponseMessageSetV1 is The OFX element "WIREXFERMSGSRSV1" is of type "WireTransferResponseMessageSetV1"
type WireTransferResponseMessageSetV1 struct {
	WIRETRNRS  []*WireTransactionResponse `xml:"WIRETRNRS"`
	WIRESYNCRS []*WireSyncResponse        `xml:"WIRESYNCRS"`
	*AbstractResponseMessageSet
}

// PresentmentDirRequestMessageSetV1 is The OFX element "PRESDIRMSGSRQV1" is of type "PresentmentDirRequestMessageSetV1"
type PresentmentDirRequestMessageSetV1 struct {
	FINDBILLERTRNRQ *FindBillerTransactionRequest `xml:"FINDBILLERTRNRQ"`
	*AbstractRequestMessageSet
}

// PresentmentDirResponseMessageSetV1 is The OFX element "PRESDIRMSGSRSV1" is of type "PresentmentDirResponseMessageSetV1"
type PresentmentDirResponseMessageSetV1 struct {
	FINDBILLERTRNRS *FindBillerTransactionResponse `xml:"FINDBILLERTRNRS"`
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
