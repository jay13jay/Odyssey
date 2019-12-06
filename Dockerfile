# pull official base image
FROM golang:1.13.5-alpine

# set work directory
WORKDIR /usr/src/app

# copy project
COPY . /usr/src/app/

RUN go install .

# run entrypoint.sh
ENTRYPOINT ["/usr/src/app/entrypoint.sh"]

CMD ["go", "run", "chat.go"]