# Go语言静态网站生成器

一个使用Go语言构建的轻量级静态网站生成器，可轻松部署到任何静态托管服务。

[English Documentation](README.md)

## 功能特点

- 易于自定义和扩展
- 通过JSON API提供数据，便于前端使用
- 支持ICP备案信息展示
- 响应式设计，适配移动端和桌面端
- 简洁的个人卡片式布局
- 支持图标链接和网站链接的自定义

## 部署方法

### 快速开始

项目包含Makefile以简化构建和部署流程：

```bash
# 显示所有可用命令
make help

# 构建服务器应用
make build

# 构建并运行服务器
make run

# 生成静态网站
make static

# 一步完成：构建所有组件并生成静态网站
make all-in-one

# 清理生成的文件
make clean
```

### 方法一：本地服务器

1. 安装Go（版本1.16或更高）
2. 克隆此仓库
3. 进入项目目录
4. 运行`make run`
5. 打开浏览器，访问http://localhost:8080

### 方法二：静态生成（推荐用于托管）

项目提供了专门的工具，可生成完全静态的HTML网站：

1. 运行`make static`
2. 静态网站将生成在`dist`目录中
3. 将`dist`目录中的文件上传到任何静态托管服务（GitHub Pages、Netlify、Vercel等）

## 自定义

要自定义网站，请修改`data/pagedata.go`中的数据。该文件包含所有网站的数据结构和内容。

## 项目结构

- `main.go` - 主服务器应用
- `data/pagedata.go` - 数据结构和内容
- `cmd/generate/main.go` - 静态网站生成器
- `templates/template.html` - HTML模板
- `static/` - 静态资源（CSS、图片等）
- `dist/` - 生成的静态网站（运行`make static`后）

## 版权声明

本站资源均来自互联网收集，仅供用于学习和交流，请勿用于商业用途。如有侵权，请联系网站管理并出示版权证明以便删除！

Copyright © 2020-2025 Linqi All Rights Reserved. 