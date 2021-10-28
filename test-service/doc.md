# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/test/handler.proto](#proto/test/handler.proto)
    - [Test](#test.Test)
  
- [proto/test/test.proto](#proto/test/test.proto)
    - [Entity](#test.Entity)
    - [EntityAddReq](#test.EntityAddReq)
    - [EntityAddRes](#test.EntityAddRes)
    - [EntityDelReq](#test.EntityDelReq)
    - [EntityDelRes](#test.EntityDelRes)
    - [EntityOneReq](#test.EntityOneReq)
    - [EntityOneRes](#test.EntityOneRes)
    - [EntityPageReq](#test.EntityPageReq)
    - [EntityPageRes](#test.EntityPageRes)
    - [EntityUpdateReq](#test.EntityUpdateReq)
    - [EntityUpdateRes](#test.EntityUpdateRes)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/test/handler.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/test/handler.proto


 

 

 


<a name="test.Test"></a>

### Test


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| EntityAdd | [EntityAddReq](#test.EntityAddReq) | [EntityAddRes](#test.EntityAddRes) |  |
| EntityDel | [EntityDelReq](#test.EntityDelReq) | [EntityDelRes](#test.EntityDelRes) |  |
| EntityUpdate | [EntityUpdateReq](#test.EntityUpdateReq) | [EntityUpdateRes](#test.EntityUpdateRes) |  |
| EntityOne | [EntityOneReq](#test.EntityOneReq) | [EntityOneRes](#test.EntityOneRes) |  |
| EntityPage | [EntityPageReq](#test.EntityPageReq) | [EntityPageRes](#test.EntityPageRes) |  |

 



<a name="proto/test/test.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/test/test.proto



<a name="test.Entity"></a>

### Entity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityAddReq"></a>

### EntityAddReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityAddRes"></a>

### EntityAddRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityDelReq"></a>

### EntityDelReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityDelRes"></a>

### EntityDelRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityOneReq"></a>

### EntityOneReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityOneRes"></a>

### EntityOneRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityPageReq"></a>

### EntityPageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |






<a name="test.EntityPageRes"></a>

### EntityPageRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| total_record | [int64](#int64) |  |  |
| total_page | [int64](#int64) |  |  |
| data | [Entity](#test.Entity) | repeated |  |






<a name="test.EntityUpdateReq"></a>

### EntityUpdateReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="test.EntityUpdateRes"></a>

### EntityUpdateRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

