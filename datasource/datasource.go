/**
 * 获取数据库连接
 */
 package datasource

 import (
	 "fmt"
	 "log"
	 "sync"
 
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/go-xorm/xorm"
 
	 "free/conf"
	 "time"
	 "github.com/spf13/viper"
 )
 
 var (
	 masterEngine *xorm.Engine
	 slaveEngine  *xorm.Engine
	 lock         sync.Mutex
 )
 
 // 主库，单例
 func InstanceMaster() *xorm.Engine {
	 if masterEngine != nil {
		 return masterEngine
	 }
	 lock.Lock()
	 defer lock.Unlock()
 
	 if masterEngine != nil {
		 return masterEngine
	 }
	 c := conf.GetDBConf()
	 driveSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		 c.User, c.Pwd, c.Host, c.Port, c.DbName)
	 log.Println(driveSource)
	 engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	 if err != nil {
		 log.Fatal("dbhelper.DbInstanceMaster,", err)
		 return nil
	 }
	 // Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	 engine.ShowSQL(viper.GetBool("app.debug"))
	 engine.SetTZLocation(conf.SysTimeLocation)
	 engine.SetConnMaxLifetime(5*time.Minute)
 
 
	 // 性能优化的时候才考虑，加上本机的SQL缓存
	 if (viper.GetBool("app.debug") == false) {
		 cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), viper.GetInt("mysql.MaxCacheSize"))
		 engine.SetDefaultCacher(cacher)
	 }
	 engine.SetMaxOpenConns(viper.GetInt("mysql.MaxOpenConns"))
	 engine.SetMaxIdleConns(viper.GetInt("mysql.MaxIdleConns"))
 
	 masterEngine = engine
	 return engine
 }
 
 
 