#!/bin/bash

curl -X POST -H "Content-Type: application/json" --data '{"user":"batman101", "text":"hello"}' http://localhost:8081/message &
sleep $((RANDOM % 5));curl -X POST -H "Content-Type: application/json" --data '{"user":"batman102", "text":"hello"}' http://localhost:8081/message &
sleep $((RANDOM % 4));curl -X POST -H "Content-Type: application/json" --data '{"user":"batman103", "text":"hello"}' http://localhost:8081/message &
sleep $((RANDOM % 2));curl -X POST -H "Content-Type: application/json" --data '{"user":"batman104", "text":"hello"}' http://localhost:8081/message &
curl "http://localhost:8081/messages"
curl "http://localhost:8081/users"

curl -X POST -H "Content-Type: application/json" --data '{"user":"batman101", "text":"hello"}' http://localhost:8081/message &
sleep $((RANDOM % 5));curl -X POST -H "Content-Type: application/json" --data '{"user":"batman102", "text":"hello"}' http://localhost:8081/message &
sleep $((RANDOM % 4));curl -X POST -H "Content-Type: application/json" --data '{"user":"batman103", "text":"hello"}' http://localhost:8081/message &
sleep $((RANDOM % 3));curl -X POST -H "Content-Type: application/json" --data '{"user":"batman104", "text":"hello"}' http://localhost:8081/message &
curl "http://localhost:8081/messages"
curl "http://localhost:8081/users"

wait
