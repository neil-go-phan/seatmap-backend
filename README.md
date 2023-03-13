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
## Entity Relationship Diagrams
[![](https://mermaid.ink/img/pako:eNqNUcuKwzAM_BWjc_sDvi3sHgp7aq-BImIlNfgRJJmyNPn3Ok0hXdJDdZE9MxYz8g3a7AgsEH977Bljk0ytIsRixnG_H0fDOZAYa9qcFH16VdyWy1ynn-Ph69d4t0Ki7FP_kCaMtCEGFLlm3r4QDLoBuxLC-e2c2eAKqo8kinEwLRMquTPqO7YM7h87LW1J-1GwWfriaIIdROKI3tWNPiY0oBeqPNh6dNRhqcGgSbMUi-bTX2rBKhfaweLn-Q1gOwxC0x2q0YLn?type=png)](https://mermaid.live/edit#pako:eNqNUcuKwzAM_BWjc_sDvi3sHgp7aq-BImIlNfgRJJmyNPn3Ok0hXdJDdZE9MxYz8g3a7AgsEH977Bljk0ytIsRixnG_H0fDOZAYa9qcFH16VdyWy1ynn-Ph69d4t0Ki7FP_kCaMtCEGFLlm3r4QDLoBuxLC-e2c2eAKqo8kinEwLRMquTPqO7YM7h87LW1J-1GwWfriaIIdROKI3tWNPiY0oBeqPNh6dNRhqcGgSbMUi-bTX2rBKhfaweLn-Q1gOwxC0x2q0YLn)
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
