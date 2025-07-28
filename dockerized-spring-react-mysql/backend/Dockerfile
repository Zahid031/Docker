FROM maven:3.9-eclipse-temurin-17-alpine as build

WORKDIR /app

COPY pom.xml ./
COPY .mvn .mvn
COPY mvnw .
COPY src src/

RUN mvn clean package -DskipTests

FROM tomcat:10.1.24-jdk21

VOLUME /tmp

EXPOSE 8080

COPY --from=build /app/target/*.war /usr/local/tomcat/webapps/ROOT.war
