FROM golang:alpine
WORKDIR /app
COPY . .
EXPOSE 8080
LABEL authors="Samuel Uzoagba, Jeremiah Bakere & Jude Eze"
CMD ["go", "run", "."]
