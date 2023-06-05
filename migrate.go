package couress

import (
	"fmt"
	"github.com/redhoe/couress/global"
	"github.com/redhoe/couress/global/core/loger"
	"github.com/redhoe/couress/global/modeler"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(args ...modeler.MigrateTable) {
	db := global.GbDB
	slog := loger.NewLogger("migrate", "migrate").Sugar()
	CurrentDatabase := db.Migrator().CurrentDatabase()
	slog.Info(fmt.Sprintf("当前数据库[%s]", CurrentDatabase))
	slog.Info("开始迁移InnoDB引擎表")
	mTables := make([]modeler.MigrateTable, 0)
	adminTables := []modeler.MigrateTable{
		modeler.NewAdministrator(),
		modeler.NewAdministratorLog(),
		modeler.NewRole(),
		modeler.NewPermission(),
		modeler.NewConfig(),
		modeler.NewOrders(), // 测试用
	}
	baseTables := []modeler.MigrateTable{
		modeler.NewChain(),
		modeler.NewChainNode(),
		modeler.NewCoin(),
		modeler.NewCurrencyExchangeRate(),
		modeler.NewDocumentTag(),
		modeler.NewDocumentBanner(),
		modeler.NewDocument(),
		modeler.NewIssueTag(),
		modeler.NewIssue(),
		modeler.NewIssueMessage(),
		modeler.NewMarketCoin(),
		modeler.NewMarket(),
		modeler.NewPoster(),
		modeler.NewVersionDocument(),
		modeler.NewVersion(),
		modeler.NewVersionGkeyModel(),
		modeler.NewWalletCiphertext(),
		modeler.NewWalletIdentity(),
		modeler.NewWalletChain(),
		modeler.NewWalletCoin(),
	}
	mTables = append(mTables, adminTables...)
	mTables = append(mTables, baseTables...)
	mTables = append(mTables, args...)
	migrationTable(db, slog, mTables) //httpCommon
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
