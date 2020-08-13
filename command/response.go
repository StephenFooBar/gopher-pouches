package command

type Response struct {
	Message string
	Success bool
}

const (
	DataStoreNotSet string = "Data store is not set. Please run config to set data store information."
	InvalidCommand  string = "Invalid Command."
	Successful      string = "Successfully executed."
)
