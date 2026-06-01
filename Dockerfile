FROM amd64/alpine
EXPOSE 8080
WORKDIR /app
COPY smtp_admin_ui .
CMD ["/app/smtp_admin_ui"]
