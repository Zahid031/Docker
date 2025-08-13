FROM gradle:9.0.0-jdk21 AS build
WORKDIR /app
COPY . .
RUN ./gradlew clean build

FROM eclipse-temurin:21-jdk
WORKDIR /app
COPY --from=build /app/app/build/libs/*.jar app.jar
EXPOSE 8080
CMD ["java", "-jar", "app.jar"]
