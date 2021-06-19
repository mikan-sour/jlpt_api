FROM golang:1.13

ENV REPO_URL=github.com/jedzeins/jlpt_api/dictionaryService
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src

# /app/src/github.com/jedzeins/jlpt_api/dictionaryService/src

COPY src $WORKPATH

WORKDIR $WORKPATH

RUN go get github.com/lib/pq
RUN go build -o dictionaryService . 
CMD ["./dictionaryService"]

