package db

type Model struct {
	*DatabaseConnection
	tableName string
}
