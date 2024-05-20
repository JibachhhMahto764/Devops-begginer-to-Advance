# first command
docker pull jenkins/jenkins
# second command 
docker run -d -p 8080:8080 docker.io/jenkins/jenkins:latest
# third command
docker ps

-> Copy the container_ID
# forth command 
docker exec <container_id> /var/jenkins_home/secrets/initialAdminPassword

-> Copy the password 
# fifth command 
local:8080

paste the password
