# gotodo

gotodo is a simple JSON api for interacting with MongoDB. It was built with intention of being a test application, so it probably
shouldn't be used as a serious application. You have been warned. :smile:


Example CURL commands:

_GET (read)_

	curl -X GET http://localhost:9090/person/5514797a7143751558000006

_POST (create)_

	curl -H "Content-Type: application/json" -d '{"firstName":"john","lastName":"smith","email":"jsmith@smithwesson.com","phoneNumber":"715-555-5555"}' http://localhost:9090/person

_PUT (update)_

	curl -H "Content-Type: application/json" -X PUT -d '{"id":"5514797a7143751558000006","firstName":"john","lastName":"smith","email":"johnsmith@smithwesson.com","phoneNumber":"715-555-5555"}' http://localhost:9090/person

_DELETE (exactly what it suggests)_

	curl -X DELETE http://localhost:9090/person/5514797a7143751558000006
