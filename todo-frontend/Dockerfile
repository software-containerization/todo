FROM node as build

WORKDIR /app

COPY package*.json /app/

RUN npm install

COPY . /app

ARG configuration=production

RUN npm run build

FROM nginx

COPY --from=build /app/build /usr/share/nginx/html

COPY /nginx-custom.conf /etc/nginx/conf.d/default.conf