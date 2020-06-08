FROM debian:bullseye-slim
#RUN apk --no-cache add ca-certificates mailcap bash && addgroup -S app && adduser -S app -G app
# extract SAR archive, run ldconfig
USER root
WORKDIR /app
EXPOSE 8000

COPY go-admin /bin
#COPY .env /app
COPY config/rbac_model.conf /app
COPY config/settings.yml /app
COPY config/db.sql /app
#ADD app.tar.gz /

#ENV PATH="${PATH}:/usr/sap/nwrfcsdk/bin"

ENTRYPOINT [ "/bin/go-admin" ]
