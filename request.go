package zip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Request struct
type Req struct {
	*http.Request
}

func (r *Req) JSON(s interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&s)
}

// Save form file to disk
func (r *Req) SaveFormFile(formName, dstName string) (err error) {
	// Get form file
	file, _, err := r.FormFile(formName)
	defer file.Close()
	if err != nil {
		return err
	}

	// Create new dst file
	dst, err := os.Create(dstName)
	if err != nil {
		return err
	}

	// Copy form file to destination
	_, err = io.Copy(dst, file)

	return err
}

// Save form file to disk and panic if err
func (r *Req) MustSaveFormFile(formName, dstName string) {
	err := r.SaveFormFile(formName, dstName)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to save file: %v", err)))
	}
}
