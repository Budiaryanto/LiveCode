package main

import "path"

type AppConfig struct {
	dataPath string
	db       []*Warehouse
}

func newConfig() *AppConfig {
	const DATA_PATH = "dataWarehouse"
	const DATA_FILE_NAME = "db_warehouse.json"

	var warehouseCollection = make([]*Warehouse, 0)
	return &AppConfig{
		dataPath: path.Join(DATA_PATH, DATA_FILE_NAME),
		db:       warehouseCollection,
	}
}

func (c *AppConfig) getDB() []*Warehouse {
	return c.db
}

func (c *AppConfig) getDbPath() string {
	return c.dataPath
}
