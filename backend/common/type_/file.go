package type_

import "mime/multipart"

type VideoFile struct {
	Files  *multipart.FileHeader
	Userid int
}
