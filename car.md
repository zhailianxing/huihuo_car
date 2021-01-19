# /common.send_code
接口功能：发送验证码
### 请求参数
| name | type |note |
| --- | --- | --- |
| phone | string | 手机号|

### 返回值结构

| name | type | note |
| --- | --- | --- |
|ret | int | 状态码。 0:正常。 非0:发生错误 |
|msg | string | 状态描述信息 |
|data | Object | 数据 |

#### data
| name | type | note |
| --- | --- | --- |
|verify_code | string |验证码 |


# /user.register
接口功能：用户注册
### 请求参数
| name | type |note |
| --- | --- | --- |
| phone | string | 手机号|
| pwd | string | 密码|
| verify_code | string | 验证码|

### 返回值结构
| name | type | note |
| --- | --- | --- |
|ret | int | 状态码。 0:正常。 非0:发生错误 |
|msg | string | 状态描述信息 |
|data | Object | 数据 |

#### data
| name | type | note |
| --- | --- | --- |
|ret | int | 状态码。 0:正常。 非0:发生错误 |
|msg | string | 状态描述信息 |


# /user.login
接口功能：用户登陆
### 请求参数
| name | type |note |
| --- | --- | --- |
| phone | string | 手机号|
| pwd | string | 密码|
| verify_code | string | 验证码|
| login_type | int | 登陆类型。 1:短信验证码登陆；2:密码登陆 |

注:当login_type=1时，pwd可以不设值。 同理,当login_type=2时，verify_code可以不设值。 

### 返回值结构
| name | type | note |
| --- | --- | --- |
|ret | int | 状态码。 0:正常。 非0:发生错误 |
|msg | string | 状态描述信息 |
|data | Object | 数据 |

### 返回值结构
| name | type | note |
| --- | --- | --- |
|user_id | int | 用户id |
| token | string |用户token |

# /user.report
接口功能：用户信息上报
### 请求参数
| name | type |note |
| --- | --- | --- |
| company_id | string | 企业id|
| brand_id | string | 品牌id|
| brand_type | int | 品牌类型|
| device_id | string | 设备id |
| token | string | 登陆时服务器下发的token |

### 返回值结构

| name | type | note |
| --- | --- | --- |
|ret | int | 状态码。 0:正常。 非0:发生错误 |
|msg | string | 状态描述信息 |





