# 使用官方 Node.js 镜像
FROM node:14

# 设置工作目录
WORKDIR /app

# 将 package.json 和 package-lock.json 复制到工作目录
COPY ./package*.json ./

# 安装项目依赖
RUN npm install

# 拷贝项目文件和文件夹到工作目录
COPY . .

# 暴露端口
EXPOSE 8080

# 运行 Vue 应用
CMD ["npm", "run", "serve"]