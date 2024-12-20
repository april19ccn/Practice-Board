# Elysia with Bun runtime

## Getting Started
To get started with this template, simply paste this command into your terminal:
```bash
bun create elysia ./elysia-example
```

## Development
To start the development server run:
```bash
bun run dev
```

Open http://localhost:3000/ with your browser to see the result.


## Directory Structure
/
├── config/                 # 配置相关的文件（如数据库配置、环境变量等）
│   └── db.js               # 数据库配置文件
├── controllers/            # 控制器层，处理业务逻辑
│   └── userController.js   # 用户相关控制器
├── middlewares/            # 中间件
│   └── authMiddleware.js   # 认证中间件
├── models/                 # 数据模型，通常用于与数据库交互
│   └── userModel.js        # 用户模型
├── routes/                 # 路由层，定义 API 路由
│   └── userRoutes.js       # 用户相关的路由
├── services/               # 服务层，封装复杂业务逻辑和数据库操作
│   └── userService.js      # 用户服务
├── utils/                  # 工具库，存放通用的工具函数
│   └── logger.js           # 日志工具
├── validations/            # 数据验证
│   └── userValidation.js   # 用户数据验证逻辑
├── app.js                  # 应用入口，初始化 Express 应用
├── server.js               # 启动文件，启动 Express 服务
├── .env                    # 环境变量文件
├── package.json            # 项目配置文件
└── README.md               # 项目说明文档