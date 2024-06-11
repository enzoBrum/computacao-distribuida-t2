FROM maven:3.9.7-eclipse-temurin-21-alpine

WORKDIR /app

COPY pom.xml .

RUN mvn dependency:go-offline

COPY src src

RUN mvn clean package

CMD [ "java", "-Dlog4j.configurationFile=src/main/resources/log4j2.xml", "-Draft_id=${RAFT_ID}", "-jar", "target/espaco-tuplas-1.0-SNAPSHOT.jar" ]
