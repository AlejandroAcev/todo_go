<div align="center">

# TODO_GO <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" height="25px" alt="golang icon" />

</div>

TODO_GO is a simple `todo` API created in Golang with a minimum number of dependencies and configuration.

## Content

- [Inspiration](#inspiration)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Routes](#routes)

## Inspiration

As many times, when you want to start a new frontend project, you will need to create or to find a simple backend API to structure your application. Usually using something like <a href="https://rickandmortyapi.com/documentation">The Rick and Morty API</a> which is very complete, but doesn't allow you make any change on the API. In this case, you will be able to use the API and change it as much as you need.

## Prerequisites

In order to run this project you will need Go v1.17 installed on your local environment: <a href="https://go.dev/doc/install">Golang</a>.

Also you will need to create a database in MongoDB. To do this you can install MonogoDB on your computer or use a free cloud solution: <a href="https://www.mongodb.com/">MongoDB</a>.

## Installation


### 1. Clone the project

Clone this project using the following command:

```bash
git clone git@github.com:AlejandroAcev/todo_go.git
```

### 2. Install the dependencies

Install all the required dependencies using:

```bash
go get ./...
```

### 3. Add the configuration

Add the required values to connect the API with the MongoDB database and the desired port in the `.env` file:

```bash
PORT=8080
MONGO_PASS=MySecretPass
MONGO_URI=mongodb+srv://<your_user>:${MONGO_PASS}@mongogeneric.ujgee.mongodb.net/<database_name>?retryWrites=true&w=majority
```

### 4. Run the project

The simplest way to run the code is usgin the following command. 

```bash
go run main.go
```

But if you prefer you can compile the code and generate an executable to run on your machine:

```bash
go build -o todo_go #or the name that you prefer
```

## Routes

In order to connect your web application you can use any of the following endpoints:

```
GET         /api/tasks               # Shows all the task created
POST        /api/task                # Creates a new task with the specified data
GET         /api/task/<id>           # Shows the details of a specific task
POST | PUT  /api/task/<id>           # Modify/Update the task selected
DELETE      /api/task/<id>           # Delete the specified task
```

Also you can find a <a href="https://www.thunderclient.com/">Thunder collection</a> and a thunder environment file with all the required data and information to use the app. 

Example of body for create a new task:

```javascript
POST /api/task

{
    "title": "Title example!",
    "description": "This is a very long description",
    "tags": ["Tag1", "Tag2"],
    "completed": false
}
```

## Credits
Alejandro Acevedo - @AlejandroAcev

## License
The MIT License (MIT)

Copyright (c) 2022 Alejandro Acevedo

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.