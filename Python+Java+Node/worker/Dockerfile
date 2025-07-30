# worker/Dockerfile

# Stage 1: Build with Maven
FROM maven:3.9.6-eclipse-temurin-17-alpine AS build
WORKDIR /app
COPY . .
RUN mvn clean package 

# Stage 2: Runtime image
FROM openjdk:17-alpine
WORKDIR /app
COPY --from=build /app/target/worker-jar-with-dependencies.jar worker.jar
CMD ["java", "-jar", "worker.jar"]
