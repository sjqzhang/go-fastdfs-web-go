package commons

const (
	ApiStatusSuccess  = "ok"
	ApiPrefix         = "/api"
	ApiReload         = ApiPrefix + "/reload"
	ApiStat           = ApiPrefix + "/stat"
	ApiUpload         = ApiPrefix + "/upload"
	ApiDelete         = ApiPrefix + "/delete"
	ApiGetFileInfo    = ApiPrefix + "/get_file_info"
	ApiListDir        = ApiPrefix + "/list_dir"
	ApiRepairStat     = ApiPrefix + "/repair_stat"
	ApiRemoveEmptyDir = ApiPrefix + "/remove_empty_dir"
	ApiBackup         = ApiPrefix + "/backup"
	ApiRepair         = ApiPrefix + "/repair"
	ApiRepairFileInfo = ApiPrefix + "/repair_fileinfo"
	ApiStatus         = ApiPrefix + "/status"
	ApiBigUpload      = ApiPrefix + "/big/upload/"

	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeText = "text/plain"
	ContentTypeHtml = "text/html"
)
