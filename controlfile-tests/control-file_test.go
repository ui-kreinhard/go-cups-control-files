package controlfile_tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/ui-kreinhard/go-cups-control-files/controlFile"
	"io/ioutil"
	"testing"
)

func assertControlFileToJson(t *testing.T, controlFilename string) {
	dataRaw, err := ioutil.ReadFile(controlFilename + "_expectation.json")
	assert.NoError(t, err)
	jobExpectation := controlFile.Job{}
	err = json.Unmarshal(dataRaw, &jobExpectation)
	assert.NoError(t, err)
	controlFileBytes, err := ioutil.ReadFile(controlFilename)
	assert.NoError(t, err)
	jobActual := controlFile.ParseBytes(controlFileBytes)
	assert.Equal(t, jobExpectation, *jobActual)
}


func TestControlFileAgainstJson(t *testing.T) {
	t.Run("should parse successfully with attribute at end of the file", func(t *testing.T) {
		assertControlFileToJson(t, "files/panic_eof")
	})

	t.Run("should contain InputTray element", func(t *testing.T) {
		assertControlFileToJson(t, "files/inputTray")
	})
}