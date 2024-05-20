# first command
docker pull jenkins/jenkins
# second command 
docker run -d -p 8080:8080 docker.io/jenkins/jenkins:latest
# third command
docker ps

-> copy the container id
# forth command 
docker exec <container_id> /var/jenkins_home/secrets/initialAdminPassword

-> copy the password 
# fifth commnad 
local:8080

paste the password
