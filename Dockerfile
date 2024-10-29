# Go 이미지를 기반으로 설정
FROM golang:1.23

# 작업 디렉토리 설정
WORKDIR /app

# Go 모듈 파일 복사 및 종속성 설치
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사 및 빌드
COPY . .
RUN go build -o server .

# 컨테이너 실행 시 gRPC 서버 실행
CMD ["./server"]

# 외부 노출 포트 지정
EXPOSE 50051
