// Code generated by Kitex v0.12.1. DO NOT EDIT.

package stock

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"

	"example_shop/kitex_gen/example/shop/base"
)

var (
	_ = base.KitexUnusedProtection
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *GetItemStockRequest) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetItemId bool = false
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetItemId = true
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	if !issetItemId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemStockRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
RequiredFieldNotSetError:
	return offset, thrift.NewProtocolException(thrift.INVALID_DATA, fmt.Sprintf("required field %s is not set", fieldIDToName_GetItemStockRequest[fieldId]))
}

func (p *GetItemStockRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.ItemId = _field
	return offset, nil
}

func (p *GetItemStockRequest) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *GetItemStockRequest) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *GetItemStockRequest) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *GetItemStockRequest) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 1)
	offset += thrift.Binary.WriteI64(buf[offset:], p.ItemId)
	return offset
}

func (p *GetItemStockRequest) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *GetItemStockResponse) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField255(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemStockResponse[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *GetItemStockResponse) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := thrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Stock = _field
	return offset, nil
}

func (p *GetItemStockResponse) FastReadField255(buf []byte) (int, error) {
	offset := 0
	_field := base.NewBaseResponse()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.BaseResponse = _field
	return offset, nil
}

func (p *GetItemStockResponse) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *GetItemStockResponse) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField255(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *GetItemStockResponse) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field255Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *GetItemStockResponse) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I64, 1)
	offset += thrift.Binary.WriteI64(buf[offset:], p.Stock)
	return offset
}

func (p *GetItemStockResponse) fastWriteField255(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 255)
	offset += p.BaseResponse.FastWriteNocopy(buf[offset:], w)
	return offset
}

func (p *GetItemStockResponse) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I64Length()
	return l
}

func (p *GetItemStockResponse) field255Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += p.BaseResponse.BLength()
	return l
}

func (p *StockServiceGetItemStockArgs) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StockServiceGetItemStockArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *StockServiceGetItemStockArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0
	_field := NewGetItemStockRequest()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Req = _field
	return offset, nil
}

func (p *StockServiceGetItemStockArgs) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *StockServiceGetItemStockArgs) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *StockServiceGetItemStockArgs) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *StockServiceGetItemStockArgs) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 1)
	offset += p.Req.FastWriteNocopy(buf[offset:], w)
	return offset
}

func (p *StockServiceGetItemStockArgs) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += p.Req.BLength()
	return l
}

func (p *StockServiceGetItemStockResult) FastRead(buf []byte) (int, error) {

	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StockServiceGetItemStockResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *StockServiceGetItemStockResult) FastReadField0(buf []byte) (int, error) {
	offset := 0
	_field := NewGetItemStockResponse()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = _field
	return offset, nil
}

func (p *StockServiceGetItemStockResult) FastWrite(buf []byte) int {
	return p.FastWriteNocopy(buf, nil)
}

func (p *StockServiceGetItemStockResult) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *StockServiceGetItemStockResult) BLength() int {
	l := 0
	if p != nil {
		l += p.field0Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *StockServiceGetItemStockResult) fastWriteField0(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *StockServiceGetItemStockResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += thrift.Binary.FieldBeginLength()
		l += p.Success.BLength()
	}
	return l
}

func (p *StockServiceGetItemStockArgs) GetFirstArgument() interface{} {
	return p.Req
}

func (p *StockServiceGetItemStockResult) GetResult() interface{} {
	return p.Success
}
