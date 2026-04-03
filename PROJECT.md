# 后台管理系统 - 项目设计方案

---

## 1. 项目模块划分

```
后台管理系统
├── 用户管理模块 (User)
│   ├── 用户 CRUD
│   ├── 密码管理
│   └── 角色分配
│
├── 角色管理模块 (Role)
│   ├── 角色 CRUD
│   └── 权限分配
│
├── 权限管理模块 (Permission)
│   ├── 权限 CRUD
│   └── 权限树
│
├── 系统设置模块 (System)
│   ├── 菜单管理
│   ├── 配置管理
│   └── 日志管理
│
└── Dashboard 模块
    ├── 数据统计
    └── 工作台
```

---

## 2. 前后端目录结构

### 后端 (Go + Gin) - `backend/`

```
backend/
├── cmd/
│   └── main.go                  # 入口文件
├── internal/
│   ├── api/
│   │   ├── v1/                  # API 路由 (user, role, permission)
│   │   └── middleware/         # 中间件 (auth, log, cors)
│   ├── model/                   # 数据模型
│   ├── repository/              # 数据访问层
│   ├── service/                  # 业务逻辑层
│   ├── dto/                      # 数据传输对象
│   ├── config/                   # 配置管理
│   └── utils/                    # 工具类
├── scripts/                      # 迁移脚本
├── docker/                       # Docker 配置
├── go.mod
└── go.sum
```

### 前端 (Vue3 + TypeScript) - `frontend/`

```
frontend/
├── src/
│   ├── api/                      # API 请求封装
│   ├── views/                    # 页面
│   │   ├── user/                 # 用户管理
│   │   ├── role/                 # 角色管理
│   │   ├── permission/           # 权限管理
│   │   └── dashboard/            # 工作台
│   ├── components/               # 公共组件
│   ├── stores/                   # Pinia 状态管理
│   ├── router/                   # 路由配置
│   ├── directives/               # 自定义指令
│   ├── utils/                    # 工具类
│   └── assets/                   # 静态资源
├── public/                       # 公共资源
├── package.json
├── tsconfig.json
└── vite.config.ts
```

---

## 3. 技术栈建议

### 后端

| 技术 | 说明 |
|------|------|
| Go 1.21+ | 开发语言 |
| Gin | Web 框架 |
| GORM | ORM |
| PostgreSQL | 数据库 |
| Redis | 缓存/会话 |
| JWT | 认证 |
| Zap | 日志 |

### 前端

| 技术 | 说明 |
|------|------|
| Vue 3.4+ | 框架 |
| TypeScript | 语言 |
| Vite | 构建工具 |
| Pinia | 状态管理 |
| Vue Router | 路由 |
| Element Plus | UI 组件库 |
| Axios | HTTP 客户端 |

---

## 4. 数据库拆分

**单体架构：逻辑分表**

```
admin_db
├── users              # 用户表
├── roles             # 角色表
├── permissions       # 权限表
├── user_roles        # 用户-角色关联
├── role_permissions  # 角色-权限关联
├── menus             # 菜单表
├── operation_logs    # 操作日志
└── system_configs    # 系统配置
```

---

## 5. 微服务 / 单体建议

**推荐：单体部署**

- 适合后台管理系统规模
- 简化开发和运维
- 减少部署复杂度

---

## 6. 基础能力

### Docker

```yaml
# docker-compose.yml
services:
  app:
    build: ./backend
    ports: ["8080:8080"]
  
  postgres:
    image: postgres:15
  
  redis:
    image: redis:7-alpine
```

### CI/CD

- GitHub Actions / GitLab CI
- 自动化测试 → 构建 → 部署

### 权限控制

- RBAC (Role-Based Access Control)
- 前端路由守卫 + 后端中间件校验

### 日志

- 结构化日志 (zap)
- 操作日志记录

### 监控

- Prometheus + Grafana

---

## 下一步

使用 CRUD 模板生成具体模块：

```text
/crud-template 用户管理
/crud-template 角色管理
/crud-template 权限管理
```
