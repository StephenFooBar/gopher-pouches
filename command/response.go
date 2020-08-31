package command

type Response struct {
	Message string
	Success bool
}

const (
	DataStoreNotSet       string = "Data store is not set. Please run config to set data store information."
	InvalidCommand        string = "Invalid Command."
	Successful            string = "Successfully executed."
	ConfigFileMissing     string = "Config file is missing. Please check that file exists or run config to set up."
	InvalidConfig         string = "Config file is not valid. Please check that the file is a proper yaml file with valid elements. Run config to reset the file."
	ConfigEntryMissing    string = "There is a missing entry in the config file. Run config to reset the file."
	DataStoreNotSupported string = "Data store specified in the config is not supported yet. Choose another or provide your own implementation."
)
