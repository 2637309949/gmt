# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/helloworld/handler.proto](#proto/helloworld/handler.proto)
    - [Helloworld](#helloworld.Helloworld)
  
- [proto/helloworld/helloworld.proto](#proto/helloworld/helloworld.proto)
    - [Article](#helloworld.Article)
    - [ArticleAddReq](#helloworld.ArticleAddReq)
    - [ArticleAddRes](#helloworld.ArticleAddRes)
    - [ArticleDelReq](#helloworld.ArticleDelReq)
    - [ArticleDelRes](#helloworld.ArticleDelRes)
    - [ArticleOneReq](#helloworld.ArticleOneReq)
    - [ArticleOneRes](#helloworld.ArticleOneRes)
    - [ArticlePageReq](#helloworld.ArticlePageReq)
    - [ArticlePageRes](#helloworld.ArticlePageRes)
    - [ArticleUpdateReq](#helloworld.ArticleUpdateReq)
    - [ArticleUpdateRes](#helloworld.ArticleUpdateRes)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/helloworld/handler.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/helloworld/handler.proto


 

 

 


<a name="helloworld.Helloworld"></a>

### Helloworld


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ArticleAdd | [ArticleAddReq](#helloworld.ArticleAddReq) | [ArticleAddRes](#helloworld.ArticleAddRes) |  |
| ArticleDel | [ArticleDelReq](#helloworld.ArticleDelReq) | [ArticleDelRes](#helloworld.ArticleDelRes) |  |
| ArticleUpdate | [ArticleUpdateReq](#helloworld.ArticleUpdateReq) | [ArticleUpdateRes](#helloworld.ArticleUpdateRes) |  |
| ArticleOne | [ArticleOneReq](#helloworld.ArticleOneReq) | [ArticleOneRes](#helloworld.ArticleOneRes) |  |
| ArticlePage | [ArticlePageReq](#helloworld.ArticlePageReq) | [ArticlePageRes](#helloworld.ArticlePageRes) |  |

 



<a name="proto/helloworld/helloworld.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/helloworld/helloworld.proto



<a name="helloworld.Article"></a>

### Article



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticleAddReq"></a>

### ArticleAddReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticleAddRes"></a>

### ArticleAddRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticleDelReq"></a>

### ArticleDelReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticleDelRes"></a>

### ArticleDelRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticleOneReq"></a>

### ArticleOneReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticleOneRes"></a>

### ArticleOneRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticlePageReq"></a>

### ArticlePageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| order | [string](#string) |  |  |






<a name="helloworld.ArticlePageRes"></a>

### ArticlePageRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [uint64](#uint64) |  |  |
| size | [int64](#int64) |  |  |
| total_record | [int64](#int64) |  |  |
| total_page | [int64](#int64) |  |  |
| data | [Article](#helloworld.Article) | repeated |  |






<a name="helloworld.ArticleUpdateReq"></a>

### ArticleUpdateReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |






<a name="helloworld.ArticleUpdateRes"></a>

### ArticleUpdateRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| creater | [int64](#int64) |  |  |
| create_time | [string](#string) |  |  |
| updater | [int64](#int64) |  |  |
| update_time | [string](#string) |  |  |
| is_delete | [uint32](#uint32) |  |  |
| remark | [string](#string) |  |  |
| name | [string](#string) |  |  |
| x_user | [string](#string) |  |  |





 

 

 

 



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

