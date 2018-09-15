# chatserver
A simple chatserver in golang
The chat server runs on http://localhost:8081 and supports the following REST API:

 1. GET /messages

     list 100 most recent messages, sorted by 'timestamp' posted to the chat server.

     example:
     ========

     ```
     curl -H "Content-Type: application/json" http://localhost:8081/messages

     {
       "messages: [
         {"timestamp": 1491345710, "user": "superman", "text": "hello"},
         {"timestamp": 1491345713, "user": "batman", "text": "hello"}
       ]
     }

     ```

 2. POST /message

     a request to post the given message.
     when the message is processed by the server a unix timestamp is recorded with each message.

     example:
     ========

     ```
     curl -X POST -H "Content-Type: application/json" --data '{"user":"superman", "text":"hello"}' http://localhost:8081/message

     {
       "ok": true
     }
     ```

 3. GET /users

     a request to return a set of users that have posted messages so far.

     example:
     ========

     ```
     curl -H "Content-Type: application/json" http://localhost:8081/users

     {
       "users": [
         "superman", "batman"
       ]
     }
     ```

the server should use appropriate http status codes for data and route validation

Instructions:
=============

Metrics you would monitor to track the performance of the chat server.
   1. CPU usage: % cpu used
   2. Memory usage: Memory footprint of the server
   3. Uptime: Once running, how long before it runs into problems
   4. Response time: In miliseconds. Round trip time from request to response. This would be typically measured like: time curl -H "Content-Type: application/json" http://localhost:8081/users
   5. Average Response Time(ART):Average of response times measured in a time window of few minutes
   6. Peak Response Time(PRT): In ms.Peak of response times measured in a time window of few minutes
   7. PRT vs ART: If PRT is closer to ART means all queries perform similarly.This can be a problem meaning all queries are slow. An occasional PRT being higher than ART implies a one-off slow down.

Improvements:
=============
   1. Need the ability to delete/purge messages 
   2. Need the ability to pass N to get N messages instead of fixed 100.
   3. In terms of the chat server itself, in order to scale it we need to back it with a database.

