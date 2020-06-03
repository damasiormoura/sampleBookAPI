FROM alpine:3.12

COPY ./sampleBookAPI /app/sampleBookAPI
RUN chmod +x /app/sampleBookAPI

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/sampleBookAPI
