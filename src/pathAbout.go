package src

import (
	"path"
)

func IsTruePath(filePath *string) bool {
	var flag bool = true

	_, file := path.Split(*filePath)
	if file == "" {
		flag = false
	}

	if ext := path.Ext(*filePath); ext == "" {
		*filePath += ".docx"
	}
	return flag
}
