package storagex

import "mime/multipart"

type FileUploadObject struct {
	File     multipart.File `json:"file"`
	FileName string         `json:"file_name"`
}
