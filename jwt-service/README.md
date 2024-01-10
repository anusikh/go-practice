### Jwt auth

- gin, GORM, bcrypt, postgres
- to run: `CompileDaemon --command="./jwt-service"  //"./GO_MOD_NAME"`
- create .env file in base as seen below
```
PORT=5000
DB=postgres://root:root@localhost:5432/jwt-service
SECRET="anusikh2001@gmail.com"
```