package global

import (
	"context"
	"fmt"
	"sort"

	"admin/app/admin/gen/query"
	"admin/pkg/config/store/mongodb"
	"admin/pkg/config/store/mysql"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)

// DBClient 数据源客户端集合。
// 按数据库类型分别保存（key 为配置中的数据源名），未配置的类型不会初始化。
type DBClient struct {
	// Mysql MySQL 数据源
	Mysql map[string]*gorm.DB
	// Mongodb MongoDB 数据源
	Mongodb map[string]*mongo.Client
	// DefaultMysqlName 默认 MySQL 数据源名
	DefaultMysqlName string
	// DefaultMongoName 默认 MongoDB 数据源名
	DefaultMongoName string
}

// CloseDatabase 关闭所有数据源连接（进程退出时调用）
func CloseDatabase() {
	if AppDB != nil {
		AppDB.Close()
	}
}

// initDatabase 按类型初始化数据源：MySQL 必配，MongoDB 未配置则跳过
func initDatabase() error {
	AppDB = &DBClient{
		Mysql:            make(map[string]*gorm.DB),
		Mongodb:          make(map[string]*mongo.Client),
		DefaultMysqlName: AppConfig.Store.Mysql.Default,
		DefaultMongoName: AppConfig.Store.Mongodb.Default,
	}
	if err := initMysqlSources(); err != nil {
		return err
	}
	return initMongoSources()
}

// initMysqlSources 初始化全部 MySQL 数据源，并将默认库设置为 gen 的全局库
func initMysqlSources() error {
	sources := AppConfig.Store.Mysql.Sources
	if len(sources) == 0 {
		return fmt.Errorf("store: 未配置任何 mysql 数据源")
	}
	for name, cfg := range sources {
		cfg := cfg
		db, err := mysql.New(&cfg)
		if err != nil {
			return fmt.Errorf("store: mysql 数据源 %q 初始化失败: %w", name, err)
		}
		AppDB.Mysql[name] = db
	}
	// 未指定默认名时，若只有一个数据源则取它
	if AppDB.DefaultMysqlName == "" && len(AppDB.Mysql) == 1 {
		for name := range AppDB.Mysql {
			AppDB.DefaultMysqlName = name
		}
	}
	def := AppDB.DefaultMysql()
	if def == nil {
		return fmt.Errorf("store: mysql 默认数据源 %q 不存在", AppDB.DefaultMysqlName)
	}
	// gen 生成的查询使用全局默认库
	query.SetDefault(def)
	return nil
}

// initMongoSources 初始化全部 MongoDB 数据源（未配置则跳过）
func initMongoSources() error {
	sources := AppConfig.Store.Mongodb.Sources
	if len(sources) == 0 {
		return nil
	}
	for name, cfg := range sources {
		cfg := cfg
		client, err := mongodb.New(context.Background(), &cfg)
		if err != nil {
			return fmt.Errorf("store: mongodb 数据源 %q 初始化失败: %w", name, err)
		}
		AppDB.Mongodb[name] = client
	}
	if AppDB.DefaultMongoName == "" && len(AppDB.Mongodb) == 1 {
		for name := range AppDB.Mongodb {
			AppDB.DefaultMongoName = name
		}
	}
	if AppDB.DefaultMongoName != "" {
		if _, ok := AppDB.Mongodb[AppDB.DefaultMongoName]; !ok {
			return fmt.Errorf("store: mongodb 默认数据源 %q 不存在", AppDB.DefaultMongoName)
		}
	}
	return nil
}

// DefaultMysql 返回默认 MySQL 连接
func (c *DBClient) DefaultMysql() *gorm.DB {
	if c == nil {
		return nil
	}
	if db, ok := c.Mysql[c.DefaultMysqlName]; ok {
		return db
	}
	// 兜底：按名称排序后取第一个，保证结果稳定
	for _, name := range c.MysqlNames() {
		return c.Mysql[name]
	}
	return nil
}

// MysqlDB 返回指定名称的 MySQL 连接，不存在时返回 nil
func (c *DBClient) MysqlDB(name string) *gorm.DB {
	if c == nil {
		return nil
	}
	return c.Mysql[name]
}

// MysqlNames 返回已初始化的 MySQL 数据源名（有序）
func (c *DBClient) MysqlNames() []string {
	if c == nil {
		return nil
	}
	names := make([]string, 0, len(c.Mysql))
	for name := range c.Mysql {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// DefaultMongo 返回默认 MongoDB 客户端
func (c *DBClient) DefaultMongo() *mongo.Client {
	if c == nil {
		return nil
	}
	if client, ok := c.Mongodb[c.DefaultMongoName]; ok {
		return client
	}
	for _, name := range c.MongoNames() {
		return c.Mongodb[name]
	}
	return nil
}

// MongoClient 返回指定名称的 MongoDB 客户端，不存在时返回 nil
func (c *DBClient) MongoClient(name string) *mongo.Client {
	if c == nil {
		return nil
	}
	return c.Mongodb[name]
}

// MongoDatabase 返回指定数据源的默认库；name 为空时使用默认数据源
func (c *DBClient) MongoDatabase(name string) *mongo.Database {
	client := c.DefaultMongo()
	if name != "" {
		client = c.MongoClient(name)
	}
	if client == nil {
		return nil
	}
	dbName := ""
	if conf, ok := AppConfig.Store.Mongodb.Sources[effectiveMongoName(c, name)]; ok {
		dbName = conf.Database
	}
	return client.Database(dbName)
}

// MongoNames 返回已初始化的 MongoDB 数据源名（有序）
func (c *DBClient) MongoNames() []string {
	if c == nil {
		return nil
	}
	names := make([]string, 0, len(c.Mongodb))
	for name := range c.Mongodb {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Ping 检查所有已初始化数据源的连通性，返回不可用的数据源（key 为 类型/名称）
func (c *DBClient) Ping(ctx context.Context) map[string]error {
	result := make(map[string]error)
	if c == nil {
		return result
	}
	for _, name := range c.MysqlNames() {
		sqlDB, err := c.Mysql[name].DB()
		if err != nil {
			result["mysql/"+name] = err
			continue
		}
		if err = sqlDB.PingContext(ctx); err != nil {
			result["mysql/"+name] = err
		}
	}
	for _, name := range c.MongoNames() {
		if err := c.Mongodb[name].Ping(ctx, nil); err != nil {
			result["mongodb/"+name] = err
		}
	}
	return result
}

// Close 关闭所有数据源连接
func (c *DBClient) Close() {
	if c == nil {
		return
	}
	for _, name := range c.MongoNames() {
		_ = c.Mongodb[name].Disconnect(context.Background())
	}
	for _, name := range c.MysqlNames() {
		if sqlDB, err := c.Mysql[name].DB(); err == nil {
			_ = sqlDB.Close()
		}
	}
}

// effectiveMongoName 取实际使用的数据源名（name 为空时用默认名）
func effectiveMongoName(c *DBClient, name string) string {
	if name != "" {
		return name
	}
	return c.DefaultMongoName
}
