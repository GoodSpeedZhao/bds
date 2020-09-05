package comparable

import (
	"strings"
	"time"

	"github.com/GoodSpeedZhao/bds/utils"
)

// Should return a number:
//    -1 , if a < b
//     0 , if a == b
//     1 , if a > b
type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
	aInt := a.(int)
	bInt := b.(int)
	switch {
	case aInt > bInt:
		return 1
	case aInt < bInt:
		return -1
	default:
		return 0
	}
}

func Int8Comparator(a, b interface{}) int {
	aInt8 := a.(int8)
	bInt8 := b.(int8)
	switch {
	case aInt8 > bInt8:
		return 1
	case aInt8 < bInt8:
		return -1
	default:
		return 0
	}
}

func Int16Comparator(a, b interface{}) int {
	aInt16 := a.(int16)
	bInt16 := b.(int16)
	switch {
	case aInt16 > bInt16:
		return 1
	case aInt16 < bInt16:
		return -1
	default:
		return 0
	}
}

func Int32Comparator(a, b interface{}) int {
	aInt32 := a.(int32)
	bInt32 := b.(int32)
	switch {
	case aInt32 > bInt32:
		return 1
	case aInt32 < bInt32:
		return -1
	default:
		return 0
	}
}

func Int64Comparator(a, b interface{}) int {
	aInt64 := a.(int64)
	bInt64 := b.(int64)
	switch {
	case aInt64 > bInt64:
		return 1
	case aInt64 < bInt64:
		return -1
	default:
		return 0
	}
}

func UintComparator(a, b interface{}) int {
	aUint := a.(uint)
	bUint := b.(uint)
	switch {
	case aUint > bUint:
		return 1
	case aUint < bUint:
		return -1
	default:
		return 0
	}
}

func Uint8Comparator(a, b interface{}) int {
	aUint8 := a.(uint8)
	bUint8 := b.(uint8)
	switch {
	case aUint8 > bUint8:
		return 1
	case aUint8 < bUint8:
		return -1
	default:
		return 0
	}
}

func Uint16Comparator(a, b interface{}) int {
	aUint16 := a.(uint16)
	bUint16 := b.(uint16)
	switch {
	case aUint16 > bUint16:
		return 1
	case aUint16 < bUint16:
		return -1
	default:
		return 0
	}
}

func Uint32Comparator(a, b interface{}) int {
	aUint32 := a.(uint32)
	bUint32 := b.(uint32)
	switch {
	case aUint32 > bUint32:
		return 1
	case aUint32 < bUint32:
		return -1
	default:
		return 0
	}
}

func Uint64Comparator(a, b interface{}) int {
	aUint64 := a.(uint64)
	bUint64 := b.(uint64)
	switch {
	case aUint64 > bUint64:
		return 1
	case aUint64 < bUint64:
		return -1
	default:
		return 0
	}
}

func Float32Comparator(a, b interface{}) int {
	aFloat32 := a.(float32)
	bFloat32 := b.(float32)
	if utils.IsFloat32Equal(aFloat32, bFloat32) {
		return 0
	} else if aFloat32 > bFloat32 {
		return 1
	} else {
		return -1
	}
}

func Float64Comparator(a, b interface{}) int {
	aFloat64 := a.(float64)
	bFloat64 := b.(float64)
	if utils.IsFloat64Equal(aFloat64, bFloat64) {
		return 0
	} else if aFloat64 > bFloat64 {
		return 1
	} else {
		return -1
	}
}

func ByteComparator(a, b interface{}) int {
	aByte := a.(byte)
	bByte := b.(byte)
	switch {
	case aByte > bByte:
		return 1
	case aByte < bByte:
		return -1
	default:
		return 0
	}
}

func RuneComparator(a, b interface{}) int {
	aRune := a.(rune)
	bRune := b.(rune)
	switch {
	case aRune > bRune:
		return 1
	case aRune < bRune:
		return -1
	default:
		return 0
	}
}

func TimeComparator(a, b interface{}) int {
	aTime := a.(time.Time)
	bTime := b.(time.Time)
	switch {
	case aTime.After(bTime):
		return 1
	case aTime.Before(bTime):
		return -1
	default:
		return 0
	}
}

func StringComparator(a, b interface{}) int {
	aString := a.(string)
	bString := b.(string)
	return strings.Compare(aString, bString)
}
