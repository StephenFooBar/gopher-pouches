package command

type Response struct {
	Message string
	Success bool
}

const (
	DataStoreNotSet      string = "Data store is not set. Please run config to set data store information."
	InvalidCommand       string = "Invalid Command."
	Successful           string = "Successfully executed."
	ConfigFileMissing    string = "Config file is missing. Please check that file exists or run config to set up."
	InvalidConfig        string = "Config file is not valid. Please check that the file is a proper yaml file with valid elements. Run config to reset the file."
	DatastoreTypeMissing string = "Database type is unspecified or missing in the config file."
)
