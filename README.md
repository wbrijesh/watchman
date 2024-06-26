# Watchman

## Overview

Watchman is a hacky small scale logging tool for my personal projects kind of like that art you made in kindergarden (awesome, but not exactly museum worthy).

i wrote this purely for my own amusement, since it doesn't even have basic auth, this project is strictly a "don't try this at home" situation.

Maybe someday I'll get around to adding auth, but even then, this is best used with toy projects.

### Project Management API

#### TOC

- [1. Create Project](#1-create-project)
- [2. List Projects](#2-list-projects)
- [3. Get Project By ID](#3-get-project-by-id)
- [4. Update Project](#4-update-project)
- [5. Delete Project](#5-delete-project)

#### 1. Create Project

```sh
curl -X POST http://127.0.0.1:4000/projects -d '{"Name": "Project 1"}'
```

#### 2. List Projects

```sh
curl http://127.0.0.1:4000/projects
```

#### 3. Get Project By ID

```sh
curl -X POST http://127.0.0.1:4000/project -d '{"ID": "1"}'
```

#### 4. Update Project

```sh
curl -X PUT http://127.0.0.1:4000/projects -d '{"ID": "1", "Name": "Project 1 Updated"}'
```

#### 5. Delete Project

```sh
curl -X DELETE http://127.0.0.1:4000/projects -d '{"ID": "1"}'
```

### Log Management API

#### TOC

- [1. Create a New Log](#1-create-a-new-log)
- [2. Get All Logs](#2-get-all-logs)
- [3. Get Logs by Project ID](#3-get-logs-by-project-id)
- [4. Get Logs by User ID](#4-get-logs-by-user-id)
- [5. Get Logs in Time Range](#5-get-logs-in-time-range)
- [6. Get Logs by Level](#6-get-logs-by-level)
- [7. Delete Logs by Project ID](#7-delete-logs-by-project-id)
- [8. Delete Logs by User ID](#8-delete-logs-by-user-id)
- [9. Delete Logs by Time Range](#9-delete-logs-by-time-range)

#### 1. Create a New Log

```sh
curl -X POST http://127.0.0.1:4000/logs -H "Content-Type: application/json" -d '{"level": "INFO", "message": "Log message", "subject": "Log subject", "user_id": "user1", "project_id": "project1"}'
```

#### 2. Get All Logs

```sh
curl -X GET http://127.0.0.1:4000/logs
```

#### 3. Get Logs by Project ID

```sh
curl -X GET 'http://127.0.0.1:4000/logs/project?project_id=project1'
```

#### 4. Get Logs by User ID

```sh
curl -X GET 'http://127.0.0.1:4000/logs/user?user_id=user1'
```

#### 5. Get Logs in Time Range

```sh
curl -X GET 'http://127.0.0.1:4000/logs/time?start_time=TIMESTAMP&end_time=TIMESTAMP'
```

#### 6. Get Logs by Level

```sh
curl -X GET 'http://127.0.0.1:4000/logs/level?level=INFO'
```

#### 7. Delete Logs by Project ID

```sh
curl -X DELETE 'http://127.0.0.1:4000/logs/project?project_id=project1'
```

#### 8. Delete Logs by User ID

```sh
curl -X DELETE 'http://127.0.0.1:4000/logs/user?user_id=user1'
```

#### 9. Delete Logs by Time Range

```sh
curl -X DELETE 'http://127.0.0.1:4000/logs/time?start_time=TIMESTAMP&end_time=TIMESTAMP'
```
