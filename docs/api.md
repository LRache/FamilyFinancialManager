# API

## 统一规范

通用字段：

- 时间传递统一传递时间戳

- 金额统一以最小单位分传递

- 响应结构:

```json
{
  "message": "string",
  "response": 相应包裹
}
```

## 用户鉴权

### 注册

`POST /api/v1/auth/register`

请求体

```json
{
  "username": "string",
  "password": "string",
}
```

响应包裹

```json
{
  "userid": 0,
  "username": "string",
}
```

### 登录

`POST /api/v1/auth/login`

请求体

```json
{
  "username": "string",
  "password": "string"
}
```

响应包裹

```json
{
  "token": "string", // 用于自动登录的token
}
```

## 家庭成员管理

### 家庭成员邀请

`POST /api/v1/family/invite`

请求体

```json
{
  "username": "string"
}
```

响应包裹

```json
{
  "username": "string",
  "status": "string"
}
```

### 家庭成员列表

`GET /api/v1/family/members`

响应包裹

```json
{
  "members": [
    {
      "username": "string",
      "email": "string",
    }
  ]
}
```

## 账单管理

### 上传账单

`POST /api/v1/bills`

请求体

```json
{
  "type": 0, // 账单类型，0表示支出，1表示收入
  "amount": 0, // 账单数额
  "category": "", // 账单类别
  "occurred_at": "", // 发生时间
  "note": "string" // 备注
}
```

### 查询账单

`GET /api/v1/bills`

请求参数

```json
{
  "type": 0, // 账单类型，0表示支出，1表示收入
  "start_date": "", // 开始日期
  "end_date": "", // 结束日期
  "category": "", // 账单类别
  "member": "" // 家庭成员
}
```

响应包裹

```json
{
  "bills": [
    {
      "id": 0,
      "type": 0,
      "amount": 0,
      "category": "string",
      "occurred_at": 0,
      "note": "string"
    }
  ]
}
```

### 添加定期收支

`POST /api/v1/bills/recurring`

请求体

```json
{
  "type": 0, // 账单类型，0表示支出，1表示收入
  "amount": 0, // 账单数额
  "category": "", // 账单类别
  "occurred_at": "", // 发生时间
  "note": "string", // 备注
  "interval": "monthly" // 账单周期，支持 daily, weekly, monthly
}
```

响应包裹

```json
{
  "id": 0,
  "type": 0,
  "amount": 0,
  "category": "string",
  "occurred_at": 0,
  "note": "string",
  "interval": "monthly"
}
```

### 查询定期收支

`GET /api/v1/bills/recurring`

响应包裹

```json
{
  "bills": [
    {
      "id": 0,
      "type": 0,
      "amount": 0,
      "category": "string",
      "occurred_at": 0,
      "note": "string",
      "interval": "monthly"
    }
  ]
}
```

## 预算和支出统计

### 设置预算

`POST /api/v1/budget`

请求体

```json
{
  "amount": 0,
  "start_date": "",
  "category": "string",
  "note": "string"
}
```

### 查询预算

`GET /api/v1/budget`

```json
{
  "budget": {
    "id": 0,
    "start_date": "",
    "amount": 0,
    "category": "string",
    "note": "string"
  }
}
```

### 查询支出统计

```GET /api/v1/outcome```

请求参数

```json
{
  "start_date": "",
  "end_date": "",
  "category": "string"
}
```

响应包裹

```json
{
  "amount": 0,
  "category": "string"
}
```

### 查询收入统计

```GET /api/v1/bills/income```

请求参数

```json
{
  "start_date": "",
  "end_date": "",
  "category": "string"
}
```

响应包裹

```json
{
  "amount": 0,
  "category": "string"
}
```
