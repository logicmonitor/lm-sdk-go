FROM amd64/golang:1.15
WORKDIR /lmsdk
ARG VERSION
RUN mkdir -p $GOPATH/src/github.com/logicmonitor/lm-sdk-go
COPY ./ ./
RUN go run /lmsdk/argus-customize.go /lmsdk/swagger.json /lmsdk/new-swagger.json ${VERSION}
RUN curl -L -o swagger https://github.com/go-swagger/go-swagger/releases/download/v0.27.0/swagger_linux_amd64 && chmod +x swagger && sync && mv swagger /bin
RUN swagger generate client -T=/lmsdk/templates -A "lm-sdk-go" -f /lmsdk/new-swagger.json --additional-initialism=LM --additional-initialism=SDT -t $GOPATH/src/github.com/logicmonitor/lm-sdk-go
