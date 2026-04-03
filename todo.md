## **用户中心 + SSO - 项目设计方案**

---

### 1. **项目模块划分**

```
用户中心 + SSO 系统模块划分：
├── 用户管理模块
│   ├── 用户注册/登录/注销
│   ├── 用户信息管理（Profile）
│   ├── 密码/安全设置
│   └── 账户状态管理
│
├── 认证模块 (Auth)
│   ├── 多方式认证（密码、短信、邮箱、OAuth2、LDAP）
│   ├── JWT/Session 会话管理
│   └── Token 刷新/撤销机制
│
├── SSO 模块
│   ├── 统一登录入口（IdP）
│   ├── SAML/OAuth2/OpenID Connect 协议支持
│   ├── Session 共享与单点登出
│   └── Client App 集成 SDK
│
├── 权限管理模块 (RBAC)
│   ├── 用户-角色-权限 关系管理
│   ├── 权限校验中间件
│   └── 动态权限配置
│
└── 审计与日志模块
    ├── 登录/操作日志
    ├── 异常行为检测
    └── 审计报表
```

---

### 2. **前后端目录结构**

**后端 (Go + Gin + PostgreSQL)**

```
user-center-ssgo/
├── cmd/
│   └── main.go
├── internal/
│   ├── api/              # HTTP API 层
│   │   ├── v1/          # API 版本控制
│   │   └── middleware/  # 中间件（认证、日志、限流）
│   ├── service/          # 业务逻辑层
│   │   ├── user.go
│   │   ├── auth.go
│   │   ├── sso/
│   │   └── permission.go
│   ├── model/            # 数据模型
│   │   ├── user.go
│   │   ├── role.go
│   │   ├── permission.go
│   │   └── sso/
│   ├── repository/       # 数据访问层
│   │   ├── user_repo.go
│   │   └── ...
│   ├── config/           # 配置管理
│   └── utils/            # 工具类（JWT、加密、邮件等）
├── scripts/              # 迁移脚本、部署脚本
├── docker/               # Docker 配置
├── go.mod
└── go.sum
```

**前端 (Vue3 + TypeScript + Pinia)**

```
user-center-frontend/
├── src/
│   ├── api/              # API 调用封装
│   ├── views/            # 页面
│   │   ├── user/        # 用户相关页面
│   │   ├── auth/        # 认证页面（登录、注册）
│   │   ├── sso/         # SSO 相关页面
│   │   └── permission/  # 权限管理页面
│   ├── components/       # 公共组件
│   ├── stores/           # Pinia 状态管理
│   ├── router/           # 路由配置
│   ├── directives/       # 自定义指令（权限 directives）
│   ├── utils/            # 工具类（请求封装、加密）
│   └── App.vue
├── vite.config.ts
└── package.json
```

---

### 3. **技术栈建议**

**后端:**

- **框架**: Go + Gin (高性能、简洁)
- **数据库**: PostgreSQL (支持 JSONB, 灵活扩展)
- **ORM**: GORM (成熟生态)
- **认证**: JWT + Redis (会话缓存)
- **消息队列**: Redis Streams / NATS (异步任务，如邮件发送)
- **配置**: Viper
- **日志**: Zap (高性能结构化日志)
- **测试**: testify + go-sqlmock

**前端:**

- **框架**: Vue3 + TypeScript + Vite
- **状态管理**: Pinia
- **路由**: Vue Router
- **UI 框架**: Element Plus / Ant Design Vue
- **HTTP 客户端**: Axios
- **表单**: VeeValidate
- **图标**: Iconify + Phosphor Icons

---

### 4. **数据库拆分**

**单体部署下的逻辑拆分 (Table-level):**

```
数据库: user_center_db

核心表:
├── users (用户表)
├── user_profiles (用户扩展信息)
├── auth_methods (认证方式：password, oauth2, ldap)
├── sessions (会话表)
├── roles (角色表)
├── permissions (权限表)
├── user_roles (用户-角色关联)
├── role_permissions (角色-权限关联)
├── sso_clients (SSO 客户端应用注册表)
├── sso_sessions (跨应用 SSO 会话映射)
├── audit_logs (审计日志表)
└── oauth2_tokens (OAuth2 令牌表)
```

**分库分表建议 (未来扩展):**

- 用户主数据: `user_center_db`
- 认证/会话: `auth_db` (高并发读写)
- 审计日志: `audit_db` (写多读少，可时序数据库)

---

### 5. **微服务 / 单体建议**

**推荐：单体部署 (Monolith)**

**理由:**

1. **快速验证 MVP**：用户中心 + SSO 初期功能耦合度高，单体架构开发迭代快
2. **简化 DevOps**：一个服务，无需处理服务发现、分布式事务、链路追踪等复杂问题
3. **资源占用低**：适合中小团队运维
4. **后期可拆分**：单体架构可逐步模块化，未来通过模块划分再拆分为微服务

**未来演进路径:**

```
Monolith → Domain-Driven Modules ⮝ Microservices
                              (users, auth, sso, audit)
```

---

### 6. **基础能力**

**Docker:**

```yaml
# docker-compose.yml
services:
  app:
    build: .
    ports: ["8080:8080"]
    depends_on: ["postgres", "redis"]

  postgres:
    image: postgres:15
    ports: ["5432:5432"]
    volumes: ["postgres_data:/var/lib/postgresql/data"]

  redis:
    image: redis:7-alpine
    ports: ["6379:6379"]

  redis-commander:
    image: rediscommander/redis-commander
    ports: ["8081:8081"]
    environment: ["REDIS_HOST=redis"]
```

**CI/CD:**

- **GitLab CI / GitHub Actions**
- 自动化测试 → 构建镜像 → 推送仓库 → 部署 (Kubernetes / Docker Swarm / 直接服务器)

**权限控制:**

- **RBAC (Role-Based Access Control)** + **ABAC (Attribute-Based)**
- 前端路由权限、按钮权限（v-permission 指令）
- 后端中间件校验

**日志:**

- **结构化日志 (JSON)**：zap log
- **访问日志**: Nginx 日志 + 应用中间件日志
- **日志收集**: Loki + Grafana / ELK

**监控:**

- **指标**: Prometheus + Node Exporter
- **告警**: Alertmanager
- **链路追踪**: Jaeger / Zipkin (未来微服务化时启用)
- **APM**: SkyWalking (可选)

---

### ✅ 建议下一步：生成模块 CRUD 模版

是否继续生成以下模块的 CRUD 模版？

```
/crud-template 用户管理
/crud-template 角色管理
/crud-template 权限管理
/crud-template SSO 客户端管理
```

输入 `/crud-template xxx` 即可生成：数据库表、Go API、Vue3 页面、TypeScript 类型定义
