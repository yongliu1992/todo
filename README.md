# todo-list-service
[![Build Status](https://travis-ci.com/yongliu1992/todo.svg?branch=master)](https://travis-ci.com/yongliu1992/todo)
[![codecov](https://codecov.io/gh/yongliu1992/todo/branch/master/graph/badge.svg)](https://codecov.io/gh/yongliu1992/todo)
[![Go Report Card](https://goreportcard.com/badge/github.com/yongliu1992/todo)](https://goreportcard.com/report/github.com/yongliu1992/todo)
[![Open Source Helpers](https://www.codetriage.com/yongliu1992/todo/badges/users.svg)](https://www.codetriage.com/yongliu1992/todo)
[![Join the chat at https://gitter.im/go_do/community](https://badges.gitter.im/go_do/community.svg)](https://gitter.im/go_do/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/yongliu1992/todo)](https://goreportcard.com/report/github.com/yongliu1992/todo)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/yongliu1992/todo)](https://www.tickgit.com/browse?repo=github.com/yongliu1992/todo)

Todo List Service is a GO API service which serves Todo data.
# Requests 

- Add a Todo task
```
curl -X POST -H "Content-Type: application/json" -d '{                                                                                           
  "task": "Fix bug 130320",
  "endDate": "2020-03-20",
  "labels": "Dev",
  "comm": "In progress"
  "uid":1,
}' http://127.0.0.1:8080/api/v1/todo/
```

- Update a Todo task
```
curl -X PUT -H "Content-Type: application/json" -d ~~'{~~
  "task": "Fix bug 130320",
  "endDate": "2020-03-20",
  "labels": "QA",
  "comments": "Verification"
}' http://127.0.0.1:8080/todo/{TID}
```

- Read all Todo tasks
```
curl -X GET -H "Content-Type: application/json"     http://127.0.0.1:8080/todo/
```

- Read a specific Todo task
```
curl -X GET -H "Content-Type: application/json"     http://127.0.0.1:8080/todo/{TID}
```

- Delete a specific Todo task
```
curl -X DELETE -H "Content-Type: application/json"  http://127.0.0.1:8080/todo/{TID}
