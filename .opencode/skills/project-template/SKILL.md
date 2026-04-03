---
name: project-template
description: 根据项目类型自动生成标准化前后端项目目录结构和初始化方案
---

# Project Template Skill

输入示例：

```text
/project-template 后台管理系统
```

```text
/project-template 用户中心 + SSO
```

请输出：

1. 项目模块划分
2. 前后端目录结构
3. 技术栈建议
4. 数据库拆分
5. 微服务 / 单体建议
6. Docker、CI/CD、权限、日志、监控等基础能力

重点：需要把相关项目目录结构生成到当前所属目录中。

前端的整体项目名为 `frontend`，后端的整体项目名为 `backend`。

如果适合进一步生成具体模块，请提示用户继续使用：

```text
/crud-template 用户管理
/crud-template 角色管理
/crud-template 权限管理
```
