package modeler

import "strings"

func reflectModelList() []MigrateTable {
	return []MigrateTable{
		// app公共
		NewCoin(),
		NewChain(),
		NewChainNode(),
		NewPoster(),

		// 钱包公共
		NewWalletIdentity(),
		NewWalletCiphertext(),
		NewWalletChain(),
		NewWalletCoin(),

		// 公共行情信息
		NewMarket(),
		NewMarketCoin(),

		// 公共法币
		NewCurrencyExchangeRate(),

		// 文档信息-帮助中心
		NewDocument(),
		NewDocumentTag(),
		NewDocumentBanner(),
		// 版本信息
		NewVersion(),
		NewVersionGkeyModel(),
		NewVersionDocument(),
		// 反馈信息
		NewIssueMessage(),
		NewIssueTag(),
		NewIssue(),

		// 各链交易信息

	}
}

func GetModelNames() map[string]string {
	m := make(map[string]string)
	for _, t := range reflectModelList() {
		m[t.TableName()] = t.Comment()
	}
	return m
}

func GetModelMap(s string) *map[string]string {
	for _, t := range reflectModelList() {
		if strings.ToLower(t.TableName()) == strings.ToLower(s) {
			m := reflectModelToMap(t)
			return &m
		}
	}
	return nil
}
