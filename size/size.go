package size

import (
	"fmt"
	"github.com/qwenode/gogo/convert"
	"math"
)

// size key
const (
	SIZE_UNIT_BYTE = iota
	SIZE_UNIT_KILO_BYTE
	SIZE_UNIT_MEGA_BYTE
	SIZE_UNIT_GIGA_BYTE
	SIZE_UNIT_TERA_BYTE
	SIZE_UNIT_PETA_BYTE
	SIZE_UNIT_EXA_BYTE
	SIZE_UNIT_ZETTA_BYTE
	SIZE_UNIT_YOTTA_BYTE
)

// size label
const (
	SIZE_UNIT_BYTE_LABEL       = "B"
	SIZE_UNIT_KILO_BYTE_LABEL  = "KB"
	SIZE_UNIT_MEGA_BYTE_LABEL  = "MB"
	SIZE_UNIT_GIGA_BYTE_LABEL  = "GB"
	SIZE_UNIT_TERA_BYTE_LABEL  = "TB"
	SIZE_UNIT_PETA_BYTE_LABEL  = "PB"
	SIZE_UNIT_EXA_BYTE_LABEL   = "EB"
	SIZE_UNIT_ZETTA_BYTE_LABEL = "ZB"
	SIZE_UNIT_YOTTA_BYTE_LABEL = "YB"
)

// ToHumanReadableSize get readable size
func ToHumanReadableSize(size float64, precision int, unit int) string {
	unitLabel := getReadableLabel(unit)
	if unit == SIZE_UNIT_BYTE {
		return fmt.Sprintf(`%.`+convert.ToString(precision)+"f%s", size, unitLabel)
	}
	humanSize := size / math.Pow(1024, float64(unit))
	return fmt.Sprintf(`%.`+convert.ToString(precision)+"f%s", humanSize, unitLabel)
}

func getReadableLabelList() map[int]string {
	return map[int]string{
		SIZE_UNIT_BYTE:       SIZE_UNIT_BYTE_LABEL,
		SIZE_UNIT_KILO_BYTE:  SIZE_UNIT_KILO_BYTE_LABEL,
		SIZE_UNIT_MEGA_BYTE:  SIZE_UNIT_MEGA_BYTE_LABEL,
		SIZE_UNIT_GIGA_BYTE:  SIZE_UNIT_GIGA_BYTE_LABEL,
		SIZE_UNIT_TERA_BYTE:  SIZE_UNIT_TERA_BYTE_LABEL,
		SIZE_UNIT_PETA_BYTE:  SIZE_UNIT_PETA_BYTE_LABEL,
		SIZE_UNIT_EXA_BYTE:   SIZE_UNIT_EXA_BYTE_LABEL,
		SIZE_UNIT_ZETTA_BYTE: SIZE_UNIT_ZETTA_BYTE_LABEL,
		SIZE_UNIT_YOTTA_BYTE: SIZE_UNIT_YOTTA_BYTE_LABEL,
	}
}

func getReadableLabel(unit int) string {
	list := getReadableLabelList()
	if label, ok := list[unit]; ok {
		return label
	}
	return ""
}
