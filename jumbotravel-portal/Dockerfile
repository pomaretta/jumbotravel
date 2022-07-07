FROM node:alpine AS build-app

WORKDIR /usr/src/app

COPY package.json .
COPY package-lock.json .
COPY . .

ARG SCHEMA=https
ARG HOSTNAME=api.jumbotravel.carlospomares.es
ARG ENV=PROD

ENV REACT_APP_SCHEMA=$SCHEMA
ENV REACT_APP_HOSTNAME=$HOSTNAME
ENV REACT_APP_ENVIRONMENT=$ENV

RUN npm i
RUN npm run build

# ============================ #

FROM nginx:1.15 as server-app
EXPOSE 80

COPY --from=build-app /usr/src/app/build /usr/share/nginx/html
COPY --from=build-app /usr/src/app/nginx.conf /etc/nginx/conf.d/default.conf