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
func (r *Req) SaveFormFile(formName, outputFile string) (err error) {
	file, _, err := r.FormFile(formName)
	defer file.Close()
	if err != nil {
		return err
	}

	dst, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, file)

	return err
}

func (r *Req) MustSaveFormFile(formName, outputFile string) {
	err := r.SaveFormFile(formName, outputFile)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to save file: %v", err)))
	}
}
