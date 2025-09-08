# API

## 统一规范

通用字段：

- 时间传递统一传递时间戳

- 金额统一以最小单位分传递

- 每次用户请求需要携带用户`token`用于鉴权

- 响应结构:

| 字段      | 类型   | 说明     |
| --------- | ------ | -------- |
| message   | string | 信息     |
| response  | object | 相应包裹 |

## 用户鉴权

### 注册

`POST /api/v1/auth/register`

请求体

| 字段      | 类型   | 说明   |
| --------- | ------ | ------ |
| username  | string | 用户名 |
| password  | string | 密码   |

响应包裹

| 字段      | 类型   | 说明             |
| --------- | ------ | ---------------- |
| token     | string | 自动登录token    |

### 登录

`POST /api/v1/auth/login`

请求体

| 字段      | 类型   | 说明   |
| --------- | ------ | ------ |
| username  | string | 用户名 |
| password  | string | 密码   |

响应包裹

| 字段      | 类型   | 说明             |
| --------- | ------ | ---------------- |
| token     | string | 自动登录token    |

## 家庭成员管理

### 创建家庭

`POST /api/v1/family`

### 家庭成员邀请

`POST /api/v1/family/members`

请求体

| 字段      | 类型   | 说明   |
| --------- | ------ | ------ |
| username  | string | 用户名 |

响应包裹

| 字段      | 类型   | 说明     |
| --------- | ------ | -------- |
| username  | string | 用户名   |

### 家庭成员列表

`GET /api/v1/family/members`

响应包裹

| 字段      | 类型     | 说明     |
| --------- | -------- | -------- |
| members   | array    | 成员列表 |

成员对象：

| 字段      | 类型   | 说明   |
| --------- | ------ | ------ |
| username  | string | 用户名 |
| email     | string | 邮箱   |

## 账单管理

### 上传账单

`POST /api/v1/bills`

请求体

| 字段        | 类型   | 说明                         |
| ----------- | ------ | ---------------------------- |
| type        | number | 账单类型（0表示支出，1表示收入） |
| amount      | number | 账单数额                     |
| category    | string | 账单类别                     |
| occurred_at | string | 发生时间                     |
| note        | string | 备注                         |

### 查询账单

`GET /api/v1/bills`

请求参数

| 字段        | 类型   | 说明                         |
| ----------- | ------ | ---------------------------- |
| type        | number | 账单类型（0表示支出，1表示收入） |
| start_date  | string | 开始日期                     |
| end_date    | string | 结束日期                     |
| category    | string | 账单类别                     |
| member      | string | 家庭成员                     |

响应包裹

| 字段      | 类型   | 说明     |
| --------- | ------ | -------- |
| bills     | array  | 账单列表 |

账单对象：

| 字段        | 类型   | 说明       |
| ----------- | ------ | ---------- |
| id          | number | 账单ID     |
| type        | number | 账单类型   |
| amount      | number | 账单数额   |
| category    | string | 账单类别   |
| occurred_at | number | 发生时间戳 |
| note        | string | 备注       |

### 删除账单

`DELETE /api/v1/bills/:id`

路径参数

| 字段 | 类型   | 说明   |
| ---- | ------ | ------ |
| id   | number | 账单ID |

响应包裹

| 字段    | 类型   | 说明     |
| ------- | ------ | -------- |
| message | string | 操作结果 |

### 添加定期收支

`POST /api/v1/bills/recurring`

请求体

| 字段        | 类型   | 说明                                |
| ----------- | ------ | ----------------------------------- |
| type        | number | 账单类型（0表示支出，1表示收入）      |
| amount      | number | 账单数额                            |
| category    | string | 账单类别                            |
| occurred_at | string | 发生时间                            |
| note        | string | 备注                                |
| interval    | string | 账单周期（支持 daily, weekly, monthly） |

响应包裹

| 字段        | 类型   | 说明       |
| ----------- | ------ | ---------- |
| id          | number | 账单ID     |
| type        | number | 账单类型   |
| amount      | number | 账单数额   |
| category    | string | 账单类别   |
| occurred_at | number | 发生时间戳 |
| note        | string | 备注       |
| interval    | string | 账单周期   |

### 查询定期收支

`GET /api/v1/bills/recurring`

响应包裹

| 字段      | 类型   | 说明     |
| --------- | ------ | -------- |
| bills     | array  | 账单列表 |

账单对象：

| 字段        | 类型   | 说明       |
| ----------- | ------ | ---------- |
| id          | number | 账单ID     |
| type        | number | 账单类型   |
| amount      | number | 账单数额   |
| category    | string | 账单类别   |
| occurred_at | number | 发生时间戳 |
| note        | string | 备注       |
| interval    | string | 账单周期   |

## 预算和支出统计

### 设置预算

`POST /api/v1/budget`

请求体

| 字段       | 类型   | 说明     |
| ---------- | ------ | -------- |
| amount     | number | 预算金额 |
| start_date | string | 开始日期 |
| category   | string | 预算类别 |
| note       | string | 备注     |

### 查询预算

`GET /api/v1/budget`

请求参数

| 字段       | 类型   | 说明     |
| ---------- | ------ | -------- |
| start_date | number | 开始日期 |
| category   | string | 预算类别 |

相应包裹

| 字段      | 类型   | 说明     |
| --------- | ------ | -------- |
| budget    | object | 预算对象 |

预算对象：

| 字段       | 类型   | 说明     |
| ---------- | ------ | -------- |
| start_date | number | 开始日期 |
| amount     | number | 预算金额 |
| category   | string | 预算类别 |
| note       | string | 备注     |

### 查询支出统计

```GET /api/v1/outcome```

请求参数

| 参数名     | 类型   | 描述           |
| ---------- | ------ | -------------- |
| start_date | number | 开始日期       |
| end_date   | number | 结束日期       |
| category   | string | 账单类别       |

响应包裹

| 字段      | 类型   | 说明 |
| --------- | ------ | ---- |
| amount    | number | 金额 |
| category  | string | 类别 |

### 查询收入统计

```GET /api/v1/bills/income```

请求参数

| 字段       | 类型   | 说明     |
| ---------- | ------ | -------- |
| start_date | string | 开始日期 |
| end_date   | string | 结束日期 |
| category   | string | 类别     |

响应包裹

| 字段      | 类型   | 说明 |
| --------- | ------ | ---- |
| amount    | number | 金额 |
| category  | string | 类别 |
