// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package stock

import (
	"context"
	"example_shop/kitex_gen/example/shop/base"
	"fmt"
	thrift "github.com/cloudwego/kitex/pkg/protocol/bthrift/apache"
)

type GetItemStockRequest struct {
	ItemId int64 `thrift:"item_id,1,required" frugal:"1,required,i64" json:"item_id"`
}

func NewGetItemStockRequest() *GetItemStockRequest {
	return &GetItemStockRequest{}
}

func (p *GetItemStockRequest) InitDefault() {
}

func (p *GetItemStockRequest) GetItemId() (v int64) {
	return p.ItemId
}
func (p *GetItemStockRequest) SetItemId(val int64) {
	p.ItemId = val
}

var fieldIDToName_GetItemStockRequest = map[int16]string{
	1: "item_id",
}

func (p *GetItemStockRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetItemId bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetItemId = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetItemId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemStockRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GetItemStockRequest[fieldId]))
}

func (p *GetItemStockRequest) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ItemId = _field
	return nil
}

func (p *GetItemStockRequest) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStockRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemStockRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("item_id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.ItemId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemStockRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemStockRequest(%+v)", *p)

}

func (p *GetItemStockRequest) DeepEqual(ano *GetItemStockRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ItemId) {
		return false
	}
	return true
}

func (p *GetItemStockRequest) Field1DeepEqual(src int64) bool {

	if p.ItemId != src {
		return false
	}
	return true
}

type GetItemStockResponse struct {
	Stock        int64              `thrift:"stock,1" frugal:"1,default,i64" json:"stock"`
	BaseResponse *base.BaseResponse `thrift:"BaseResponse,255" frugal:"255,default,base.BaseResponse" json:"BaseResponse"`
}

func NewGetItemStockResponse() *GetItemStockResponse {
	return &GetItemStockResponse{}
}

func (p *GetItemStockResponse) InitDefault() {
}

func (p *GetItemStockResponse) GetStock() (v int64) {
	return p.Stock
}

var GetItemStockResponse_BaseResponse_DEFAULT *base.BaseResponse

func (p *GetItemStockResponse) GetBaseResponse() (v *base.BaseResponse) {
	if !p.IsSetBaseResponse() {
		return GetItemStockResponse_BaseResponse_DEFAULT
	}
	return p.BaseResponse
}
func (p *GetItemStockResponse) SetStock(val int64) {
	p.Stock = val
}
func (p *GetItemStockResponse) SetBaseResponse(val *base.BaseResponse) {
	p.BaseResponse = val
}

var fieldIDToName_GetItemStockResponse = map[int16]string{
	1:   "stock",
	255: "BaseResponse",
}

func (p *GetItemStockResponse) IsSetBaseResponse() bool {
	return p.BaseResponse != nil
}

func (p *GetItemStockResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetItemStockResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GetItemStockResponse) ReadField1(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Stock = _field
	return nil
}
func (p *GetItemStockResponse) ReadField255(iprot thrift.TProtocol) error {
	_field := base.NewBaseResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.BaseResponse = _field
	return nil
}

func (p *GetItemStockResponse) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStockResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetItemStockResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("stock", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Stock); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetItemStockResponse) writeField255(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("BaseResponse", thrift.STRUCT, 255); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.BaseResponse.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetItemStockResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetItemStockResponse(%+v)", *p)

}

func (p *GetItemStockResponse) DeepEqual(ano *GetItemStockResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Stock) {
		return false
	}
	if !p.Field255DeepEqual(ano.BaseResponse) {
		return false
	}
	return true
}

func (p *GetItemStockResponse) Field1DeepEqual(src int64) bool {

	if p.Stock != src {
		return false
	}
	return true
}
func (p *GetItemStockResponse) Field255DeepEqual(src *base.BaseResponse) bool {

	if !p.BaseResponse.DeepEqual(src) {
		return false
	}
	return true
}

type StockService interface {
	GetItemStock(ctx context.Context, req *GetItemStockRequest) (r *GetItemStockResponse, err error)
}

type StockServiceGetItemStockArgs struct {
	Req *GetItemStockRequest `thrift:"req,1" frugal:"1,default,GetItemStockRequest" json:"req"`
}

func NewStockServiceGetItemStockArgs() *StockServiceGetItemStockArgs {
	return &StockServiceGetItemStockArgs{}
}

func (p *StockServiceGetItemStockArgs) InitDefault() {
}

var StockServiceGetItemStockArgs_Req_DEFAULT *GetItemStockRequest

func (p *StockServiceGetItemStockArgs) GetReq() (v *GetItemStockRequest) {
	if !p.IsSetReq() {
		return StockServiceGetItemStockArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *StockServiceGetItemStockArgs) SetReq(val *GetItemStockRequest) {
	p.Req = val
}

var fieldIDToName_StockServiceGetItemStockArgs = map[int16]string{
	1: "req",
}

func (p *StockServiceGetItemStockArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *StockServiceGetItemStockArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StockServiceGetItemStockArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StockServiceGetItemStockArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewGetItemStockRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *StockServiceGetItemStockArgs) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStock_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *StockServiceGetItemStockArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *StockServiceGetItemStockArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StockServiceGetItemStockArgs(%+v)", *p)

}

func (p *StockServiceGetItemStockArgs) DeepEqual(ano *StockServiceGetItemStockArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *StockServiceGetItemStockArgs) Field1DeepEqual(src *GetItemStockRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type StockServiceGetItemStockResult struct {
	Success *GetItemStockResponse `thrift:"success,0,optional" frugal:"0,optional,GetItemStockResponse" json:"success,omitempty"`
}

func NewStockServiceGetItemStockResult() *StockServiceGetItemStockResult {
	return &StockServiceGetItemStockResult{}
}

func (p *StockServiceGetItemStockResult) InitDefault() {
}

var StockServiceGetItemStockResult_Success_DEFAULT *GetItemStockResponse

func (p *StockServiceGetItemStockResult) GetSuccess() (v *GetItemStockResponse) {
	if !p.IsSetSuccess() {
		return StockServiceGetItemStockResult_Success_DEFAULT
	}
	return p.Success
}
func (p *StockServiceGetItemStockResult) SetSuccess(x interface{}) {
	p.Success = x.(*GetItemStockResponse)
}

var fieldIDToName_StockServiceGetItemStockResult = map[int16]string{
	0: "success",
}

func (p *StockServiceGetItemStockResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *StockServiceGetItemStockResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_StockServiceGetItemStockResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *StockServiceGetItemStockResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewGetItemStockResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *StockServiceGetItemStockResult) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("GetItemStock_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *StockServiceGetItemStockResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *StockServiceGetItemStockResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("StockServiceGetItemStockResult(%+v)", *p)

}

func (p *StockServiceGetItemStockResult) DeepEqual(ano *StockServiceGetItemStockResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *StockServiceGetItemStockResult) Field0DeepEqual(src *GetItemStockResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
