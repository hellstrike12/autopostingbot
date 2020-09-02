package api

import (
	"github.com/zelenin/go-tdlib/client"
	"gitlab.com/shitposting/autoposting-bot/repository"
)

// DownloadFile downloads synchronously a file with maximum priority.
func DownloadFile(fileID int32) (*client.File, error) {

	file, err := repository.Tdlib.DownloadFile(&client.DownloadFileRequest{
		FileId:      fileID,
		Priority:    32,
		Synchronous: true,
	})

	return file, err

}
