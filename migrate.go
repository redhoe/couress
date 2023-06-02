package couress

import (
	"fmt"
	"github.com/redhoe/couress/global"
	"github.com/redhoe/couress/global/modeler"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(args ...modeler.MigrateTable) {
	CurrentDatabase := global.GbDB.Migrator().CurrentDatabase()
	global.GbSLOG.Info(fmt.Sprintf("当前数据库[%s]", CurrentDatabase))
	global.GbSLOG.Info("开始迁移InnoDB引擎表")
	db := global.GbDB
	slog := global.GbSLOG
	mTables := make([]modeler.MigrateTable, 0)
	adminTables := []modeler.MigrateTable{
		modeler.Administrator{},
		modeler.AdministratorLog{},
		modeler.Role{},
		modeler.Permission{},
		modeler.Config{},
		modeler.Orders{}, // 测试用
	}
	baseTables := []modeler.MigrateTable{
		modeler.Chain{},
		modeler.ChainNode{},
		modeler.Coin{},
		modeler.CurrencyExchangeRate{},
		modeler.DocumentTag{},
		modeler.DocumentBanner{},
		modeler.Document{},
		modeler.IssueTag{},
		modeler.Issue{},
		modeler.IssueMessage{},
		modeler.MarketCoin{},
		modeler.Market{},
		modeler.Poster{},
		modeler.VersionDocument{},
		modeler.Version{},
		modeler.VersionGkeyModel{},
		modeler.WalletCiphertext{},
		modeler.WalletIdentity{},
		modeler.WalletChain{},
		modeler.WalletCoin{},
	}
	mTables = append(mTables, adminTables...)
	mTables = append(mTables, baseTables...)
	mTables = append(mTables, args...)
	migrationTable(db, slog, mTables) //common
	dataInit(db, slog)
	slog.Info(fmt.Sprintf("数据库迁移完成"))
}

// migrationTable 迁移数据表
func migrationTable(db *gorm.DB, slog *zap.SugaredLogger, tables []modeler.MigrateTable) {
	for _, table := range tables {
		slog.Info(fmt.Sprintf("开始迁移[%s]表", table.TableName()))
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
		comment := fmt.Sprintf("COMMENT='%s'", table.Comment())
		db = db.Set("gorm:table_options", comment)
		err := db.Migrator().AutoMigrate(table)
		if err != nil {
			slog.Error(fmt.Sprintf("[%s]表迁移失败：%s", table.TableName(), err.Error()))
		}
	}
}

// dataInit 数据初始化
func dataInit(db *gorm.DB, slog *zap.SugaredLogger) {
	dataAdiInit(db, slog)
	dataOtherInit(db, slog)
}

func dataAdiInit(db *gorm.DB, slog *zap.SugaredLogger) {
	if err := modeler.NewAdministrator().DataInit(db); err != nil {
		slog.Error("error:", err.Error())
	}
	if err := modeler.NewRole().DataInit(db); err != nil {
		slog.Error("error:", err.Error())
	}
	if err := modeler.NewPermission().DataInit(db); err != nil {
		slog.Error("error:", err.Error())
	}
}

func dataOtherInit(db *gorm.DB, slog *zap.SugaredLogger) {
	if err := modeler.NewChain().DataInit(db); err != nil {
		slog.Error("error:", err.Error())
	}
}
