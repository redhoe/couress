package lang

import "github.com/nicksnyder/go-i18n/v2/i18n"

func MustMessage(localizer *i18n.Localizer, msg *i18n.Message) string {
	str, _ := localizer.Localize(&i18n.LocalizeConfig{DefaultMessage: msg})
	return str
}

var MessageAddressVerifyError = &i18n.Message{
	ID:    "AddressVerifyError",
	Other: "address verify error",
}

var MessageAddressNotFound = &i18n.Message{
	ID:    "AddressNotFound",
	Other: "address not found",
}

var MessageAddressBalanceIsLow = &i18n.Message{
	ID:    "AddressBalanceIsLow",
	Other: "address balance is low",
}

var MessageAddressCoinBalanceIsLow = &i18n.Message{
	ID:    "AddressCoinBalanceIsLow",
	Other: "address coin balance is low",
}

var MessageSymbolNotFound = &i18n.Message{
	ID:    "SymbolNotFound",
	Other: "symbol not found",
}

var MessageDocumentTagNotFound = &i18n.Message{
	ID:    "MessageDocumentTagNotFound",
	Other: "tag not found",
}

var MessageDocumentNotFound = &i18n.Message{
	ID:    "MessageDocumentNotFound",
	Other: "document not found",
}

var MessagePleaseUploadPictures = &i18n.Message{
	ID:    "MessagePleaseUploadPictures",
	Other: "Please upload pictures",
}

var MessagePleaseSelectCommentType = &i18n.Message{
	ID:    "MessagePleaseSelectCommentType",
	Other: "Please select a comment type",
}

var MessageTransactionSignError = &i18n.Message{
	ID:    "MessageTransactionSignError",
	Other: "transaction sign error",
}

var MessageTransactionNotFund = &i18n.Message{
	ID:    "MessageTransactionNotFund",
	Other: "Transaction not found",
}

var MessageTransactionIsSuccess = &i18n.Message{
	ID:    "MessageTransactionIsSuccess",
	Other: "Transaction is success",
}

var MessageTransferAmountNotIsZero = &i18n.Message{
	ID:    "MessageTransferAmountNotIsZero",
	Other: "Transfer amount not is zero",
}

var MessageAccountFeesIsLow = &i18n.Message{
	ID:    "MessageAccountFeesIsLow",
	Other: "Account fees is low",
}
