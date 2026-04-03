---
name: crud-template
description: 根据实体名称自动生成标准化 CRUD 前后端模版，包括数据库、Go API、Vue3 页面、TypeScript 类型定义
---

# CRUD Template Skill

当用户输入：

```text
/crud-template 用户管理
```

或：

```text
/crud-template 商品管理
```

请自动完成以下内容：

1. 识别实体名称
2. 自动推导字段
3. 自动生成数据库表
4. 自动生成 Go 后端 CRUD
5. 自动生成 Vue3 + Element Plus 页面
6. 自动生成 API 文件与类型定义

请严格按以下结构输出：并把相关的前端代码和文件放在 `frontend` 目录下，后端代码和文件放在 `backend` 目录下。

# 1. 数据表设计

输出 MySQL 建表 SQL，例如：

```sql
CREATE TABLE users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL,
  nickname VARCHAR(50),
  phone VARCHAR(20),
  status TINYINT DEFAULT 1,
  created_at DATETIME,
  updated_at DATETIME
);
```

# 2. Go Model

生成：

```text
internal/model/user.go
```

要求：

- 使用 GORM
- 带 json 与 gorm tag
- 自动包含 ID、创建时间、更新时间

# 3. DTO

生成：

```text
internal/dto/user_dto.go
```

包含：

- CreateXXXReq
- UpdateXXXReq
- QueryXXXReq

# 4. Repository

生成：

```text
internal/repository/user_repository.go
```

包含：

- Create
- Update
- Delete
- Detail
- List

# 5. Service

生成：

```text
internal/service/user_service.go
```

要求：

- 封装业务逻辑
- 支持分页
- 支持条件查询

# 6. Controller

生成：

```text
internal/controller/user_controller.go
```

自动生成以下接口：

- POST /user/create
- POST /user/update
- POST /user/delete
- GET /user/detail
- GET /user/list

# 7. Router

生成：

```text
internal/router/user_router.go
```

# 8. 前端 API 文件

生成：

```text
src/api/user.ts
```

要求：

- 使用 axios
- 包含 create、update、delete、detail、list 方法

# 9. TypeScript 类型定义

生成：

```text
src/types/user.ts
```

# 10. Vue3 页面模版

生成：

```text
src/views/user/index.vue
```

页面需自动包含：

- 查询表单
- 表格
- 新增按钮
- 编辑按钮
- 删除按钮
- 分页
- 弹窗表单

技术要求：

- Vue3 + script setup
- TypeScript
- Element Plus
- 支持响应式

# 11. 可选增强

如果用户需求中包含以下关键词，请自动追加：

- “权限” → 自动生成权限控制代码
- “SSO” → 自动生成用户登录态校验
- “导入导出” → 自动生成 Excel 导入导出接口
- “软删除” → 自动添加 deleted_at 字段
- “树形结构” → 自动生成 parent_id
- “状态切换” → 自动生成启用 / 禁用开关

最后请补充：

- 哪些文件建议由 AI 自动生成
- 哪些业务逻辑需要人工补充
- 后续建议继续使用哪些 Skill，例如：

  - /review
  - /security-check
  - /refactor
