package krb5types

// Reference: https://www.ietf.org/rfc/rfc4120.txt
// Section: 5.2.7
import (
	"time"
	"fmt"
	"encoding/asn1"
)

type PAData struct {
	PADataType  int    `asn1:"explicit,tag:1"`
	PADataValue []byte `asn1:"explicit,tag:2"`
}

// Do I need to define this one?
type PAEncTimestamp struct {
	PAEncTSEnc
}

type PAEncTSEnc struct {
	PATimestamp time.Time `asn1:"explicit,tag:0"`
	PAUSec      int `asn1:"explicit,optional,tag:1"`
}

type ETypeInfoEntry struct {
	EType int `asn1:"explicit,tag:0"`
	Salt  []byte `asn1:"explicit,optional,tag:1"`
}

type ETypeInfo []ETypeInfoEntry

type ETypeInfo2Entry struct {
	EType     int `asn1:"explicit,tag:0"`
	Salt      string `asn1:"explicit,optional,tag:1,ia5"`
	S2KParams []byte `asn1:"explicit,optional,tag:2"`
}

type ETypeInfo2 []ETypeInfo2Entry

func (pa *PAData) GetETypeInfo() (d ETypeInfo, err error) {
	dt := krbDictionary.PADataTypesByName["pa-etype-info"]
	if pa.PADataType != dt {
		err = fmt.Errorf("PAData does not contain PA EType Info data. TypeID Expected: %v; Actual: %v", dt, pa.PADataType)
		return
	}
	_, err = asn1.Unmarshal(pa.PADataValue, &d)
	return
}

func (pa *PAData) GetETypeInfo2() (d ETypeInfo2, err error) {
	dt := krbDictionary.PADataTypesByName["pa-etype-info2"]
	if pa.PADataType != dt {
		err = fmt.Errorf("PAData does not contain PA EType Info 2 data. TypeID Expected: %v; Actual: %v", dt, pa.PADataType)
		return
	}
	_, err = asn1.Unmarshal(pa.PADataValue, &d)
	return
}