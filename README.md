# Glossika assignment

## Test step

### Set env
1. Redis
   
```
docker run -d --name redis -p 6379:6379 redis
```

2. MySQL
```
docker run -d \
  --name mysql \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=glossika \
  -e MYSQL_USER=glossika \
  -e MYSQL_PASSWORD=glossika \
  -e MYSQL_DATABASE=glossika \
  mysql:9.1
```

3. Start server
```
go run . server -d mysql
```
   


