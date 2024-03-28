FROM golang:1.22-alpine AS build
ENV TZ=Asia/Tokyo

WORKDIR /opt/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./calendeye

#------------------------------------------------------------
FROM gcr.io/distroless/base:nonroot AS runner
ENV TZ=Asia/Tokyo
ENV GOENV=production
ENV PORT=80

COPY --from=build /opt/app/calendeye /bin/calendeye

EXPOSE ${PORT}
USER nonroot
ENTRYPOINT [ "/bin/calendeye" ]
