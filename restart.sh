docker-compose down
docker rm -f go-echo
docker rmi docker rm -f go-echo
docker-compose up --build api