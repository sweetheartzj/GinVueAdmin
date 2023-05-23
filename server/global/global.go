package global

import (
	"GinVueAdmin/server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBLIST map[string]*gorm.DB
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger

	lock sync.RWMutex
)

func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBLIST[dbname]
}

func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBLIST[dbname]
	if !ok || db == nil {
		panic("db is not init")
	}
	return db
}
