package utils

import (
	ctx "context"
	"io/ioutil"

	"go.mau.fi/whatsmeow/proto/waE2E"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

func GetImageBytes(url string) ([]byte, error) {
	r, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return b, err
}

func UploadImage(url string, client *whatsmeow.Client, receiver types.JID) error {
	imageBytes, err := GetImageBytes(url)
	if err != nil {
		return err
	}

	resp, err := client.Upload(ctx.Background(), imageBytes, whatsmeow.MediaImage)
	if err != nil {
		return err
	}

	imageMsg := &waE2E.ImageMessage{
		DirectPath: &resp.DirectPath,
		MediaKey:   resp.MediaKey,
		FileLength: &resp.FileLength,
	}
	// Url:           &resp.URL,
	// DirectPath:    &resp.DirectPath,
	// MediaKey:      resp.MediaKey,
	// FileEncSha256: resp.FileEncSHA256,
	// FileSha256:    resp.FileSHA256,
	// FileLength:    &resp.FileLength,
	_, err = client.SendMessage(ctx.Background(), receiver, &waE2E.Message{
		ImageMessage: imageMsg,
	})

	return err
}
