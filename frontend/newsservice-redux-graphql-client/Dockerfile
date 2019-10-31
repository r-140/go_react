FROM nginx:1.15.10-alpine
COPY ./build /var/www
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
ENTRYPOINT ["nginx","-g","daemon off;"]

# build environment
#FROM node:12.2.0-alpine as build
#WORKDIR /app
#ENV PATH /app/node_modules/.bin:$PATH
#COPY package.json /app/package.json
#COPY config-overrides.js /app/config-overrides.js
#RUN npm install --silent
#RUN npm install react-scripts@3.0.1 -g --silent
#RUN npm install antd react-router-dom --save
#RUN npm install react-app-rewired babel-plugin-import react-app-rewire-less --save
#COPY . /app
#RUN npm run build
#
## production environment
#FROM nginx:1.16.0-alpine
#COPY --from=build /app/build /usr/share/nginx/html
#RUN rm /etc/nginx/conf.d/default.conf
#COPY nginx.conf /etc/nginx/conf.d
#EXPOSE 80
#CMD ["nginx", "-g", "daemon off;"]
