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
curl -X POST http://127.0.0.1:4000/projects -d '{"name": "Project 1"}'
```

#### 2. List Projects

```sh
curl http://127.0.0.1:4000/projects
```

#### 3. Get Project By ID

```sh
curl -X GET http://127.0.0.1:4000/projects/{id}
```

#### 4. Update Project

```sh
curl -X PUT http://127.0.0.1:4000/projects/{id} -d '{"name": "Updated Project Name"}'
```

#### 5. Delete Project

```sh
curl -X DELETE http://127.0.0.1:4000/projects/{id}
```

### Log Management API

#### TOC

- [1. Create Logs (Bulk)](#1-create-logs-bulk)
- [2. List Logs](#2-list-logs)
- [3. Delete Logs](#3-delete-logs)

#### 1. Create Logs (Bulk)

```sh
curl -X POST http://127.0.0.1:4000/logs -d '[{"time": 1615760000, "level": "info", "message": "this is a test log", "subject": "test", "user_id": "1", "project_id": "1"}]'
```

#### 2. List Logs

```sh
curl http://127.0.0.1:4000/logs?project_id=1
```

Or you add more query params to filter the logs

```sh
curl http://127.0.0.1:4000/logs?project_id=1&user_id=1&log_level=info&start_time=1615760000&end_time=1615760000
```

#### 3. Delete Logs

```sh
curl -X DELETE http://127.0.0.1:4000/logs?project_id=1
```

Or you can add more query params to filter the logs

```sh
curl -X DELETE http://127.0.0.1:4000/logs?project_id=1&user_id=2
```
