package controlFile

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

/*
operation-attributes-tag:

    attributes-charset (charset): utf-8
    attributes-natural-language (naturalLanguage): en-us

job-attributes-tag:

    printer-uri (uri): http://10.250.14.153/printers/DEGIG02-LPR003
    job-originating-user-name (nameWithoutLanguage): anonymous
    job-name (nameWithoutLanguage): interchange_(D)-MA-CO-2108_2020-01-09T14-31-55-032Z.pdf
    copies (integer): 1
    document-format (mimeMediaType): application/pdf
    job-priority (integer): 50
    job-uuid (uri): urn:uuid:833da9c5-dbfd-3c71-416f-72f5209adf3d
    job-originating-host-name (nameWithoutLanguage): 10.250.14.129
    time-at-creation (integer): 1578576743
    time-at-processing (integer): 1578576743
    time-at-completed (integer): 1578576743
    job-id (integer): 325
    job-state (enum): completed
    job-state-reasons (keyword): processing-to-stop-point
    job-media-sheets-completed (integer): 0
    job-printer-uri (uri): ipp://degig02-lsv001.contargo.net:631/printers/DEGIG02-LPR003
    job-k-octets (integer): 20
    job-hold-until (keyword): no-hold
    job-sheets (1setOf nameWithoutLanguage): none,none
    job-printer-state-message (textWithoutLanguage): 
    job-printer-state-reasons (keyword): none


 */

type OperationsAttributesTag struct {
	AttributesCharset         *string
	AttributesNaturalLanguage *string
}

type JobAttributesTag struct {
	PrinterUri              *string
	JobOriginatingUserName  *string
	JobName                 *string
	Copies                  *uint32
	DocumentFormat          *string
	JobPriority             *uint32
	JobUuid                 *string
	JobOriginatingHostName  *string
	TimeAtCreation          *uint32
	TimeAtProcessing        *uint32
	TimeAtCompleted         *uint32
	JobId                   *uint32
	JobState                *string
	JobStateReasons         *string
	JobMediaSheetsCompleted *uint32
	JobPrinterUri           *string
	JobKOctets              *uint32
	JobHoldUntil            *string
	JobSheets               *string
	JobPrinterStateMessage  *string
	JobPrinterStateReasons  *string
}

type Job struct {
	OperationsAttributesTag OperationsAttributesTag
	JobAttributesTag        JobAttributesTag
}

func (j *Job) PrintContent() {
	bytes, _ := json.Marshal(j)
	fmt.Println(string(bytes))
}

func compareToString(stringToCompare string, byteStream []byte, position int) bool {
	lengthOfString := len(stringToCompare)
	if position >= 3 && byteStream[position-2] != 0 {
		return false
	}
	if position+lengthOfString > len(byteStream) {
		return false
	}
	byteStreamPart := string(byteStream[position : position+lengthOfString])
	return byteStreamPart == stringToCompare
}

func findEnd(start int, jobFileBytes []byte) int {
	for i, singleByte := range jobFileBytes[start:] {
		if singleByte == 0 {
			return start + i - 1
		}
	}
	return -1
}

func extractString(position int, jobFileBytes []byte) *string {
	position = position + 2
	end := findEnd(position, jobFileBytes)
	ret := string(jobFileBytes[position:end])
	return &ret
}

func extractInt(position int, jobFileBytes []byte) *uint32 {
	position = position + 2
	uint := binary.BigEndian.Uint32(jobFileBytes[position : position+4])
	return &uint
}

func strategy(position int, jobFileBytes []byte, newJob *Job) {
	if compareToString("job-state-reasons", jobFileBytes, position) {
		newJob.JobAttributesTag.JobPrinterStateReasons = extractString(position + len("job-state-reasons"), jobFileBytes)
	} else if compareToString("job-printer-state-message", jobFileBytes, position) {
		newJob.JobAttributesTag.JobPrinterStateMessage = extractString(position + len("job-printer-state-message"), jobFileBytes)
	} else if compareToString("job-hold-until", jobFileBytes, position) {
		newJob.JobAttributesTag.JobHoldUntil = extractString(position + len("job-hold-until"), jobFileBytes)
	} else if compareToString("job-k-octets", jobFileBytes, position) {
		newJob.JobAttributesTag.JobKOctets = extractInt(position + len("job-k-octets"), jobFileBytes)
	} else if compareToString("job-media-sheets-completed", jobFileBytes, position) {
		newJob.JobAttributesTag.JobMediaSheetsCompleted = extractInt(position + len("job-media-sheets"), jobFileBytes)
	} else if compareToString("job-state-reasons", jobFileBytes, position) {
		newJob.JobAttributesTag.JobStateReasons = extractString(position+len("job-state-reasons"), jobFileBytes)	
	} else if compareToString("job-id", jobFileBytes, position) {
		newJob.JobAttributesTag.JobId = extractInt(position+len("job-id"), jobFileBytes)
	} else if compareToString("time-at-completed", jobFileBytes, position) {
		newJob.JobAttributesTag.TimeAtCompleted = extractInt(position+len("time-at-completed"), jobFileBytes)
	} else if compareToString("time-at-processing", jobFileBytes, position) {
		newJob.JobAttributesTag.TimeAtProcessing = extractInt(position+len("time-at-processing"), jobFileBytes)
	} else if compareToString("time-at-creation", jobFileBytes, position) {
		newJob.JobAttributesTag.TimeAtCreation = extractInt(position+len("time-at-creation"), jobFileBytes)
	} else if compareToString("job-originating-host-name", jobFileBytes, position) {
		newJob.JobAttributesTag.JobOriginatingHostName = extractString(position+len("job-originating-host-name"), jobFileBytes)
	} else if compareToString("job-uuid", jobFileBytes, position) {
		newJob.JobAttributesTag.JobUuid = extractString(position+len("job-uuid"), jobFileBytes)
	} else if compareToString("job-priority", jobFileBytes, position) {
		newJob.JobAttributesTag.JobPriority = extractInt(position+len("job-priority"), jobFileBytes)
	} else if compareToString("document-format", jobFileBytes, position) {
		newJob.JobAttributesTag.DocumentFormat = extractString(position+len("document-format"), jobFileBytes)
	} else if compareToString("copies", jobFileBytes, position) {
		newJob.JobAttributesTag.Copies = extractInt(position+len("copies"), jobFileBytes)
	} else if compareToString("job-name", jobFileBytes, position) {
		newJob.JobAttributesTag.JobName = extractString(position+len("job-name"), jobFileBytes)
	} else if compareToString("job-originating-user-name", jobFileBytes, position) {
		newJob.JobAttributesTag.JobOriginatingUserName = extractString(position+len("job-originating-user-name"), jobFileBytes)
	} else if compareToString("attributes-natural-language", jobFileBytes, position) {
		newJob.OperationsAttributesTag.AttributesNaturalLanguage = extractString(position+len("attributes-natural-language")-1, jobFileBytes)
	} else if compareToString("attributes-charset", jobFileBytes, position) {
		newJob.OperationsAttributesTag.AttributesCharset = extractString(position+len("attributes-charset"), jobFileBytes)
	} else if compareToString("printer-uri", jobFileBytes, position) {
		printerUri := extractString(position+len("printer-uri"), jobFileBytes)
		newJob.JobAttributesTag.PrinterUri = printerUri
	} else if compareToString("job-printer-uri", jobFileBytes, position) {
		jobPrinterUi := extractString(position+len("job-printer-uri"), jobFileBytes)
		newJob.JobAttributesTag.JobPrinterUri = jobPrinterUi
	}
}

func ParseBytes(jobFileBytes []byte) *Job {
	newJob := &Job{}
	for i, _ := range jobFileBytes {
		strategy(i, jobFileBytes, newJob)
	}
	return newJob
}
