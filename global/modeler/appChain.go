package modeler

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/demdxx/gocast"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type ChainSymbol string

func (c ChainSymbol) String() string {
	return string(c)
}

func (c ChainSymbol) Check() ChainSymbol {
	return lo.If(c != TRXt, TRX).Else(TRXt)
}

const (
	ETH  ChainSymbol = "ETH"
	BNB  ChainSymbol = "BNB"
	HT   ChainSymbol = "HT"
	BTC  ChainSymbol = "BTC"
	TRX  ChainSymbol = "TRX"
	TRXt ChainSymbol = "TRXT"
)

type ChainList []Chain

func (list ChainList) MustDefault(db *gorm.DB) ChainList {
	l, _ := list.Default(db)
	return l
}

func (list ChainList) Default(db *gorm.DB) (ChainList, error) {
	memList := make([]Chain, 0)
	if err := db.Find(&memList).Error; err != nil {
		return nil, err
	}
	for i, chain := range list {
		find := false
		for _, c := range memList {
			if c.Symbol == chain.Symbol {
				find = true
				list[i] = c
				break
			}
		}
		if find {
			continue
		}
		err := db.Model(&chain).Where("`symbol` = ?", chain.Symbol).Find(chain).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		if err == gorm.ErrRecordNotFound {
			chain.CreatedAt = time.Now()
			chain.UpdatedAt = time.Now()
			if err := db.Create(&chain).Error; err != nil {
				return nil, err
			}
		}

		list[i] = chain
	}

	return list, nil
}

func (list ChainList) Chain(coin ChainSymbol) (Chain, bool) {
	return lo.Find(list, func(item Chain) bool {
		if item.Symbol == coin {
			return true
		} else {
			return false
		}
	})
}

func (list ChainList) ById(id uint) (Chain, bool) {
	return lo.Find(list, func(item Chain) bool {
		if item.Id == id {
			return true
		} else {
			return false
		}
	})
}

func (list ChainList) BySymbol(symbol ChainSymbol) (Chain, bool) {
	return lo.Find(list, func(item Chain) bool {
		if item.Symbol == symbol {
			return true
		} else {
			return false
		}
	})
}

var DefaultChainList = ChainList{
	{Name: "Bitcoin", Alias: "BTC", Type: ChainTypeBITCOIN, Symbol: BTC, Enable: true, Testnet: false, Config: &ChainConfig{Decimal: 8}},
	{Name: "Ethereum", Alias: CoinTypeERC20, Type: ChainTypeETHEREUM, Symbol: ETH, Enable: true, Testnet: true,
		Config: &ChainConfig{Decimal: 18, ChainId: 1, WebUrl: "https://etherscan.io/", RpcUrl: "https://mainnet.infura.io/v3/14e5c24b98634138a9127fc8db299970"},
		SwapConfig: &ChainSwapConfig{
			RouterV2ContractAddress: "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
			UsdtTokenAddress:        "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		},
	},
	{Name: "Ethereum goerli testnet", Alias: CoinTypeERC20, Type: ChainTypeETHEREUM, Symbol: ChainSymbol("ETHg"), Enable: true, Testnet: true,
		Config: &ChainConfig{Decimal: 18, ChainId: 5, WebUrl: "https://goerli.etherscan.io/", RpcUrl: "https://goerli.infura.io/v3/14e5c24b98634138a9127fc8db299970"},
		SwapConfig: &ChainSwapConfig{
			RouterV2ContractAddress: "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
			UsdtTokenAddress:        "0xC2C527C0CACF457746Bd31B2a698Fe89de2b6d49",
		},
	},
	{Name: "BNB Smart Chain", Alias: CoinTypeBEP20, Type: ChainTypeETHEREUM, Symbol: BNB, Enable: true, Testnet: false,
		Config: &ChainConfig{Decimal: 18, ChainId: 56, WebUrl: "https://bscscan.com/", RpcUrl: "https://bsc-dataseed1.binance.org/"},
		SwapConfig: &ChainSwapConfig{
			RouterV2ContractAddress: "0x10ED43C718714eb63d5aA57B78B54704E256024E",
			UsdtTokenAddress:        "0x55d398326f99059fF775485246999027B3197955",
		},
	},
	{Name: "BNB Smart Chain testnet", Alias: CoinTypeBEP20, Type: ChainTypeETHEREUM, Symbol: ChainSymbol("BNBt"), Testnet: true, Enable: true,
		Config: &ChainConfig{Decimal: 18, ChainId: 97, WebUrl: "https://testnet.bscscan.com/", RpcUrl: "https://data-seed-prebsc-1-s1.binance.org:8545/"}},
	{Name: "Huobi ECO Chain", Alias: CoinTypeHTC20, Type: ChainTypeETHEREUM, Symbol: HT, Enable: true, Testnet: false,
		Config: &ChainConfig{Decimal: 18, ChainId: 128, WebUrl: "https://www.hecoinfo.com/", RpcUrl: "https://http-mainnet.hecochain.com"}},
	{Name: "Huobi ECO Chain testnet", Alias: CoinTypeHTC20, Type: ChainTypeETHEREUM, Symbol: ChainSymbol("HTt"), Testnet: true, Enable: true,
		Config: &ChainConfig{Decimal: 18, ChainId: 256, WebUrl: "https://testnet.hecoinfo.com/", RpcUrl: "https://http-testnet.hecochain.com"}},
	{Name: "Polygon", Alias: CoinTypeERC20, Type: ChainTypeETHEREUM, Symbol: ChainSymbol("MATIC"), Testnet: false, Enable: true,
		Config: &ChainConfig{Decimal: 18, ChainId: 137, WebUrl: "https://polygonscan.com/", RpcUrl: "https://rpc.ankr.com/polygon/1985783691c9320af01c055f8a9a315ce53308b409e1e833084f1d81b216e612", WsUrl: "wss://rpc.ankr.com/polygon/ws/1985783691c9320af01c055f8a9a315ce53308b409e1e833084f1d81b216e612"}},
	{Name: "Fantom", Alias: CoinTypeERC20, Type: ChainTypeETHEREUM, Symbol: ChainSymbol("FTM"), Testnet: false, Enable: true,
		Config: &ChainConfig{Decimal: 18, ChainId: 250, WebUrl: "https://ftmscan.com/", RpcUrl: "https://rpc.ankr.com/polygon/1985783691c9320af01c055f8a9a315ce53308b409e1e833084f1d81b216e612", WsUrl: "wss://rpc.ankr.com/fantom/ws/1985783691c9320af01c055f8a9a315ce53308b409e1e833084f1d81b216e612"}},
	{Name: "Bitcoin test3", Alias: "BTC", Type: ChainTypeBITCOIN, Symbol: ChainSymbol("BTCt"), Enable: true, Testnet: true,
		Config: &ChainConfig{Decimal: 8, RpcUrl: ""},
	},
}

type ChainType string

const (
	ChainTypeETHEREUM ChainType = "ETHEREUM" // 类ETH的链条 使用web3
	ChainTypeBITCOIN  ChainType = "BITCOIN"  // 类BTC的链条
	ChainTypeEOS      ChainType = "EOS"
	ChainTypeTron     ChainType = "TRON"
)

var ChainTypeList = []ChainType{
	ChainTypeBITCOIN,
	ChainTypeETHEREUM,
	ChainTypeEOS,
	ChainTypeTron,
}

type Chain struct {
	MysqlModel
	Name        string           `json:"name"  gorm:"type:varchar(100);comment:全名"`
	Alias       CoinType         `json:"alias" gorm:"type:varchar(20);comment:别名"`
	Type        ChainType        `json:"type" gorm:"type:varchar(10);comment:链类型"`
	Symbol      ChainSymbol      `json:"symbol" gorm:"type:varchar(10);comment:主币名"`
	Enable      bool             `json:"enable" gorm:"comment:是否开放"`
	Icon        string           `json:"icon" gorm:"type:varchar(200);comment:图标"`
	Desc        string           `json:"desc" gorm:"type:varchar(255);comment:描述"` // 链种描述
	Testnet     bool             `json:"testnet" gorm:"comment:是否测试网络"`
	Config      *ChainConfig     `json:"config" gorm:"type:json;comment:配置"`
	SwapConfig  *ChainSwapConfig `json:"swap_config" gorm:"type:json"`
	MarketApiId *string          `json:"marketApiId" gorm:"type:varchar(80);comment:配置匹配的apiId"`
	UsdRate     decimal.Decimal  `json:"usdRate" gorm:"type:decimal(38,18);comment:USD汇率"`
}

func NewChain() *Chain {
	return &Chain{}
}

func (Chain) TableName() string {
	return "chain"
}

func (Chain) Comment() string {
	return "链信息"
}

func (c *Chain) GetRecord(db *gorm.DB, column, value interface{}) bool {
	if err := db.Where(fmt.Sprintf("%s = ?", gocast.ToString(column)), value).Debug().
		Find(&c).Error; err != nil {
		return false
	}
	if c.Id != 0 {
		return true
	}
	return false
}

func (*Chain) DataInit(db *gorm.DB) error {
	for _, req := range DefaultChainList {
		find := NewChain()
		if stat := db.Model(&Chain{}).Where("symbol = ?", req.Symbol).
			Find(&find).Statement; stat.RowsAffected == 0 {
			if err := db.Create(&req).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

type ChainConfig struct {
	ApiUrl     string  `json:"api_url"`
	ApiKey     string  `json:"api_key"`
	Decimal    int32   `json:"decimal"`
	RpcUrl     string  `json:"rpc_url"`
	WsUrl      string  `json:"ws_url"`
	ChainId    int     `json:"chain_id"`
	WebUrl     string  `json:"web_url"`
	DefaultGas float64 `json:"default_gas"`
	ApiToken   string  `json:"api_token"`
}

func (c *ChainConfig) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ChainConfig) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type ChainSwapConfig struct {
	RouterV2ContractAddress string `json:"router_v2_contract_address"`
	UsdtTokenAddress        string `json:"usdt_token"`
}

func (c *ChainSwapConfig) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ChainSwapConfig) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type ChainNode struct {
	MysqlModel
	ChainId                int64      `json:"chain_id" gorm:"unique;comment:链Id"`
	LastHandlerBlockNumber *int64     `json:"last_handler_block_number" gorm:"comment:最后处理块"`
	LastHandlerTime        *time.Time `json:"last_handler_time" gorm:"comment:最后处理时间"`
	LastNetBlockNumber     *int64     `json:"last_net_block_number" gorm:"comment:网络块高"`
}

func (ChainNode) TableName() string {
	return "chain_node"
}

func (ChainNode) Comment() string {
	return "各链当前区块高度记录"
}

func NewChainNode() *ChainNode {
	return &ChainNode{}
}

func (this *ChainNode) CheckByChainId(db *gorm.DB, chainId uint) bool {
	if err := db.Where("chain_id = ?", chainId).
		Find(&this).Error; err != nil {
		return false
	}
	if this.Id != 0 {
		return true
	} else {
		this.ChainId = int64(chainId)
		this.LastHandlerBlockNumber = lo.ToPtr(int64(0))
		this.LastNetBlockNumber = lo.ToPtr(int64(0))
		this.LastHandlerTime = lo.ToPtr(time.Now())
		if err := db.Create(&this).Error; err != nil {
			return false
		}
		return true
	}
}

func (this *ChainNode) GetRecord(db *gorm.DB, column, value interface{}) bool {
	if err := db.Where(fmt.Sprintf("%s = ?", gocast.ToString(column)), value).
		Find(&this).Error; err != nil {
		return false
	}
	if this.Id != 0 {
		return true
	}
	return false
}
