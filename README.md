
根据之前为你编写的 README 内容，若需要将其添加到远程仓库 `https://github.com/SBHHX/Gocunment.git` 中，请按照以下步骤操作：


### **步骤 1：克隆远程仓库到本地**
```bash
git clone https://github.com/SBHHX/Gocunment.git
cd Gocunment
```


### **步骤 2：创建或编辑 `README.md` 文件**
#### **情况 1：仓库中无 `README.md`（初始状态）**  
直接创建文件并写入内容：
```bash
echo "# Gocument 文档管理系统" > README.md  # 初始化文件
nano README.md  # 使用文本编辑器打开（粘贴之前生成的 README 内容）
```

#### **情况 2：仓库中已有 `README.md`**  
编辑现有文件：
```bash
nano README.md  # 清空原有内容，粘贴新的 README 内容
```

**粘贴以下内容**（即之前生成的 README 全文）：
```markdown
# Gocument 文档管理系统

Gocument 是一个基于 Go 语言和 Gin 框架开发的文档管理系统，提供用户认证、文档创建、读取、更新和删除等功能。

## 特性
- 用户注册与登录（JWT 认证）
- 文档的增删改查
- 简单的权限控制（文档公开/私有）

## 安装与运行

### 环境要求
- Go 1.18+
- MySQL 8.0+

### 步骤
1. **克隆项目**
```bash
git clone <项目仓库地址>
cd gocument
```

2. **配置数据库**
- 创建数据库 `gocument`
- 编辑 `config/config.yaml`，更新数据库连接信息：
```yaml
database:
  dsn: "root:你的密码@tcp(localhost:3306)/gocument?charset=utf8mb4&parseTime=True&loc=Local"
```

3. **安装依赖**
```bash
go mod tidy
```

4. **运行项目**
```bash
go run main.go
```
服务启动后，访问 `http://localhost:8080`

## 接口文档

### 用户注册
- **URL**: `/api/auth/register`
- **方法**: `POST`
- **参数**:
```json
{
  "username": "用户名",
  "password": "密码",
  "nickname": "昵称"
}
```
- **示例响应**:
```json
{
  "message": "注册成功"
}
```

### 用户登录
- **URL**: `/api/auth/login`
- **方法**: `POST`
- **参数**:
```json
{
  "username": "用户名",
  "password": "密码"
}
```
- **示例响应**:
```json
{
  "token": "生成的JWT令牌"
}
```

### 创建文档
- **URL**: `/api/documents`
- **方法**: `POST`
- **请求头**: `Authorization: Bearer <JWT令牌>`
- **参数**:
```json
{
  "title": "文档标题",
  "content": "文档内容",
  "is_public": true  // 是否公开
}
```
- **示例响应**:
```json
{
  "message": "文档创建成功"
}
```

### 获取文档
- **URL**: `/api/documents/:id`
- **方法**: `GET`
- **请求头**: `Authorization: Bearer <JWT令牌>`
- **参数**: `id` 为文档 ID
- **示例响应**:
```json
{
  "id": 1,
  "user_id": 1,
  "title": "文档标题",
  "content_path": "存储路径",
  "is_public": true,
  "created_at": "时间戳",
  "updated_at": "时间戳"
}
```

### 更新文档
- **URL**: `/api/documents/:id`
- **方法**: `PUT`
- **请求头**: `Authorization: Bearer <JWT令牌>`
- **参数**:
```json
{
  "title": "新标题",
  "content": "新内容",
  "is_public": false
}
```
- **示例响应**:
```json
{
  "message": "文档更新成功"
}
```

### 删除文档
- **URL**: `/api/documents/:id`
- **方法**: `DELETE`
- **请求头**: `Authorization: Bearer <JWT令牌>`
- **示例响应**:
```json
{
  "message": "文档删除成功"
}
```

## 贡献
欢迎提交 PR 或 Issue 来改进项目。

## 许可证
MIT 许可证，详情见 [LICENSE](LICENSE)。
```


### **步骤 3：提交并推送更改到远程仓库**
```bash
git add README.md         # 添加文件到暂存区
git commit -m "Update README.md"  # 提交更改
git push origin main      # 推送到远程仓库的 main 分支
```


### **验证远程仓库**
访问 [https://github.com/SBHHX/Gocunment.git](https://github.com/SBHHX/Gocunment.git)，刷新页面后应看到更新后的 `README.md` 内容。

如果遇到权限问题，请确保你有仓库的写入权限（如需添加 SSH 密钥或使用 HTTPS 认证）。
