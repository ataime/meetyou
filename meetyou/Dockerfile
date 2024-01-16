FROM node:20

WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY ./package*.json ./

# 安装依赖
RUN npm install

# 复制项目文件
COPY . .

EXPOSE 8080

CMD ["npm", "run", "serve"]
