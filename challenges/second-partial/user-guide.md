# User Guide

- **Login**

```
curl -X POST -d "username=admin;password=password" http://localhost:8080/login
```

- **Logout**

```
curl -X POST -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/logout
```

- **Upload**

```
curl -X POST -F "data=@path/to/local/image.png" -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/upload
```

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU4MDI3ODAsInVzZXIiOiJhZG1pbiJ9.JrgK_5QBRL5UPYew4kAqF8YC5ccitU4TQcY52um_x7A

- **Status**

```
curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/status
```
