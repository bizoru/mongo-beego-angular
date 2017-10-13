FROM node
RUN npm install --unsafe-perm -g @angular/cli
COPY code/ /angular/code
WORKDIR /angular/code
CMD ng serve --host 0.0.0.0
