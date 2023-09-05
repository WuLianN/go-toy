package errcode

var (
	ErrorGetPictureListFail = NewError(20010001, "获取图片列表失败")
	ErrorCountPictureFail   = NewError(20010005, "统计图片失败")
	ERROR_UPLOAD_FILE_FAIL = NewError(20030001, "上传文件失败")
)