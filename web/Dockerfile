FROM node:20-alpine AS build
ENV TZ=Asia/Tokyo

WORKDIR /opt/app

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .
RUN npm run build

#------------------------------------------------------------
FROM nginx:latest
ENV TZ=Asia/Tokyo

WORKDIR /etc/nginx/html

RUN rm -rf /usr/share/nginx/html/*
COPY --from=build /opt/app/dist/ /usr/share/nginx/html/

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]
