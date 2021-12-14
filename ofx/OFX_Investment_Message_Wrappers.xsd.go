// Code generated by xgen. DO NOT EDIT.

package ofx

// InvestmentMailSyncRequest is The OFX element "INVMAILSYNCRQ" is of type "InvestmentMailSyncRequest"
type InvestmentMailSyncRequest struct {
	INCIMAGES    string                              `xml:"INCIMAGES"`
	USEHTML      string                              `xml:"USEHTML"`
	INVACCTFROM  *InvestmentAccount                  `xml:"INVACCTFROM"`
	OFXEXTENSION *OFXExtensionType                   `xml:"OFXEXTENSION"`
	INVMAILTRNRQ []*InvestmentMailTransactionRequest `xml:"INVMAILTRNRQ"`
	*AbstractSyncRequest
}

// InvestmentMailSyncResponse is The OFX element "INVMAILSYNCRS" is of type "InvestmentMailSyncResponse"
type InvestmentMailSyncResponse struct {
	INVACCTFROM  *InvestmentAccount                   `xml:"INVACCTFROM"`
	OFXEXTENSION *OFXExtensionType                    `xml:"OFXEXTENSION"`
	INVMAILTRNRS []*InvestmentMailTransactionResponse `xml:"INVMAILTRNRS"`
	*AbstractSyncResponse
}

// InvestmentMailTransactionRequest is The OFX element "INVMAILTRNRQ" is of type "InvestmentMailTransactionRequest"
type InvestmentMailTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType      `xml:"OFXEXTENSION"`
	INVMAILRQ    *InvestmentMailRequest `xml:"INVMAILRQ"`
	*AbstractTransactionRequest
}

// InvestmentMailTransactionResponse is The OFX element "INVMAILTRNRS" is of type "InvestmentMailTransactionResponse"
type InvestmentMailTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType       `xml:"OFXEXTENSION"`
	INVMAILRS    *InvestmentMailResponse `xml:"INVMAILRS"`
	*AbstractTransactionResponse
}

// InvestmentStatementTransactionRequest is The OFX element "INVSTMTTRNRQ" is of type "InvestmentStatementTransactionRequest"
type InvestmentStatementTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType           `xml:"OFXEXTENSION"`
	INVSTMTRQ    *InvestmentStatementRequest `xml:"INVSTMTRQ"`
	*AbstractTransactionRequest
}

// InvestmentStatementTransactionResponse is The OFX element "INVSTMTTRNRS" is of type "InvestmentStatementTransactionResponse"
type InvestmentStatementTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType            `xml:"OFXEXTENSION"`
	INVSTMTRS    *InvestmentStatementResponse `xml:"INVSTMTRS"`
	*AbstractTransactionResponse
}

// InvestmentStatementEndTransactionRequest is The OFX element "INVSTMTTRNRQ" is of type "InvestmentStatementTransactionRequest"
type InvestmentStatementEndTransactionRequest struct {
	OFXEXTENSION *OFXExtensionType              `xml:"OFXEXTENSION"`
	INVSTMTENDRQ *InvestmentStatementEndRequest `xml:"INVSTMTENDRQ"`
	*AbstractTransactionRequest
}

// InvestmentStatementEndTransactionResponse is The OFX element "INVSTMTTRNRS" is of type "InvestmentStatementTransactionResponse"
type InvestmentStatementEndTransactionResponse struct {
	OFXEXTENSION *OFXExtensionType               `xml:"OFXEXTENSION"`
	INVSTMTENDRS *InvestmentStatementEndResponse `xml:"INVSTMTENDRS"`
	*AbstractTransactionResponse
}
