#!/bin/sh
echo -e "\033[41;36m start deploy my blog \033[0m"

echo `pwd`
cd ~/www/blog/web
npm i
npm run build
echo -e "\033[41;36m 前端代码构建完成 \033[0m"

nohup go run main.go &
