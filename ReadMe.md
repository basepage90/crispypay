# 서버 구성
- Go 1.22
    - uber/fx
    - Gin
    - mysql

# 실행
- 서버 구동시, mysql 데이터 자동 초기화

로컬 실행 명령어
```
go run .
```

도커 실행 명령어
```
docker build -t crispypay .

docker run -e PARTNER_HOST=host.docker.internal --name crispypay -d -p 8080:8080  crispypay
```

mysql 접속 정보
- host: crispyblog.kr
- port: 3306
- db: crispypay
- user: user01
- password: user123

# API

- 헬스체크
    - method: GET
    - endpoint: /health

- 송금 요청
    - method: POST
    - endpoint: /api/transfer
    
- 단건 조회
    - method: GET
    - endpoint: /api/transfer/{partner_tx_id}
    - response: SearchTransferResponse
    
- 목록 조회
    - method: GET
    - endpoint: /api/transfers
    - reponse: []SearchTransferResponse

- 송금 취소
    - method: PUT
    - endpoint: /api/transfer/cancel/:partner_tx_id

- 콜백 처리
    - method: POST
    - endpoint: /api/v1/webhooks/partner-callback
    - request: PartnerCallback