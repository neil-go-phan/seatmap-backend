docker pull postgres:15.2-alpine
# SEATMAP APP

- Project được viết bằng Go
- Sử dụng framework Gin, ORM: GORM
- Backend project: https://github.com/neil-go-phan/seatmap-frontend

## HOW TO RUN
- Clone source code từ [commit 7242661c97e236e0adaa2e277d9ee5f10236c06c](https://github.com/neil-go-phan/seatmap-backend/tree/7242661c97e236e0adaa2e277d9ee5f10236c06c). 
- Mở terminal
- CD vào folder project
- Chạy các lệnh sau:
### `go run ./app main.go`
<!-- ### `npm install`
### `npm start` -->
## TODO
- OT: 
  - Handle error
  - Migration
  - connect to a postgres docker images local
  - create Makefile
  - readme.md
- Tomorow
  - Frontend: 
    - create map modal, style react-grid-layout
  - Backend:
    - fix error handler
        
NOTE: Refactor validate role 
