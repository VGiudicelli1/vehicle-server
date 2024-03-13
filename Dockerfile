FROM gcr.io/distroless/static-debian12

COPY dist/server /dst/server

ENTRYPOINT ["/dst/server"]