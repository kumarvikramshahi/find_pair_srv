# FIND PAIR service
A Service that allows user to input array of integers and a target value, and get all pairs of indices where the sum of the elements at those indices equals the target value.

**Note**: *Algorithm assumes that input array can be un-sorted and can have duplicate values, and will return indices with respect to original input array.*

### Quick Start
To quicly run server, make sure you have properly installed go on your local then executes below commands.

```
git clone https://github.com/kumarvikramshahi/streak_assignment.git
```

```
cd streak_assignment
```
```
go mod download
```
```
go run main.go dev
```

### Architecture
Used Hexagonal Architecture (Port-Adapter) to build the web server.

* `configs` - contains configs & env vars related logics
* `env` - contains env vars files.
* `pkg` - here the main application logics relies
  
![Untitled-2024-03-18-1524](https://github.com/user-attachments/assets/90757274-c294-41a0-8723-8a4eede27341)



### Service have two endpoints for client:

#### `/find-pairs`  -  `POST`

```
curl --location 'localhost:8000/find-pairs' \
--header 'Content-Type: application/json' \
--data '{
  "numbers": [1,2,3,4,5],
  "target": 6
}
'
```
Example request:
```
{
  "numbers": [1,2,3,4,5],
  "target": 6
}
```
Example response:
```
{
    "solutions": [
        [0,4],
        [1,3]
    ]
}
```

#### `/` - `GET`
For health check.
```
curl --location 'localhost:8000' \
--header 'Content-Type: application/json' \
--data '{
  "numbers": [1,2,3,4,5],
  "target": 6
}
'
```
Example response:
```
{"message":"Working..."}
```

