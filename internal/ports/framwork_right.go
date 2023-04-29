package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(answer int, oreration string) error
}
