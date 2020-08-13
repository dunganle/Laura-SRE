# Laura-SRE

Installation instructions
=========================

This code runs a small web server that returns a value from Redis. 
It is designed to run on a variety of Linux distributions, both Docker Engine and Docker Compose will be needed to run successfully.
Instructions for installing these tools can be found at these links:
https://docs.docker.com/engine/install/
https://docs.docker.com/compose/install/

Ensure both services are running using "systemctl status docker"

Once the files are loaded to your system, simply run the command "docker-compose up -d"

Test the web service using the command "curl http://localhost:9001"





