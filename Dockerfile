FROM maven:3.9.7-eclipse-temurin-21-alpine

WORKDIR /app

COPY pom.xml .

RUN mvn dependency:go-offline

COPY src src

RUN mvn clean package

CMD [ "java", "-Djgroups.bind_addr=127.0.0.1", "-Djava.net.preferIPv4Stack=true","-Dlog4j.configurationFile=src/main/resources/log4j2.xml", "-jar", "target/espaco-tuplas-1.0-SNAPSHOT.jar" ]
