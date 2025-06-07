FROM node:lts-alpine
LABEL org.opencontainers.image.source=https://github.com/UniversalOJ/UOJ-System
LABEL org.opencontainers.image.description="UOJ Web Aux - Markdown Rendering Server"
LABEL org.opencontainers.image.licenses=MIT

WORKDIR /opt/uoj
COPY package*.json .
RUN npm install
COPY . .

ENV LANG=C.UTF-8 TZ=Asia/Shanghai
EXPOSE 7513
CMD ["npm", "start"]
