docker pull postgres:15.2-alpine
# SEATMAP APP

- Project được viết bằng Go
- Sử dụng framework Gin, ORM: GORM
- Frontend project: https://github.com/neil-go-phan/seatmap-frontend
- Database docker container chạy trên port 2345
- Migration: [golang-migration](https://github.com/golang-migrate/migrate)

## HOW TO RUN
- Clone source code từ [commit 7242661c97e236e0adaa2e277d9ee5f10236c06c](https://github.com/neil-go-phan/seatmap-backend/tree/7242661c97e236e0adaa2e277d9ee5f10236c06c). 
- Mở terminal
- CD vào folder project
- Chạy các lệnh sau:
  - `make pull_docker_img`
  - `make postgres`
  - `make server`

<!-- ### `npm install`
### `npm start` -->
## Entity Relationship Diagrams
[![](https://mermaid.ink/img/pako:eNqNUcsKAjEM_JWSs_5Ab4IeBE96XZCwzWqhjyVJEdH9d7uuoLIezCXtzDTMpDdosyOwQLz2eGKMTTK1ihDLkJdLc78bzoHEWHNG-WDNbbqMddjst6ud8e4NibJPp6c0YaQZ0aPIJfP8hWDQGdiVEI4_54zm3qD6SKIYe9MyoZI7ov5iS---2GFqU9K_go3SD0cDLCASR_SubvM5oQE9U-XB1qOjDksNBk0apVg0H66pBatcaAGTn9cXgO0wCA0P7LiBKw?type=png)](https://mermaid.live/edit#pako:eNqNUcsKAjEM_JWSs_5Ab4IeBE96XZCwzWqhjyVJEdH9d7uuoLIezCXtzDTMpDdosyOwQLz2eGKMTTK1ihDLkJdLc78bzoHEWHNG-WDNbbqMddjst6ud8e4NibJPp6c0YaQZ0aPIJfP8hWDQGdiVEI4_54zm3qD6SKIYe9MyoZI7ov5iS---2GFqU9K_go3SD0cDLCASR_SubvM5oQE9U-XB1qOjDksNBk0apVg0H66pBatcaAGTn9cXgO0wCA0P7LiBKw)
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
