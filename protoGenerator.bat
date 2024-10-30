@echo off
REM 모든 .proto 파일에 대해 Protobuf 코드를 컴파일하는 배치 파일

REM 현재 디렉토리 출력
echo Current directory: %CD%

REM proto 디렉토리로 이동
cd proto

REM 환경변수에서 C# 코드 생성 경로를 가져옴
if "%UNITY_PROTO_PATH%"=="" (
    echo Error: UNITY_PROTO_PATH 환경변수가 설정되지 않았습니다.
    exit /b 1
)

REM 환경 변수 값 출력
echo UNITY_PROTO_PATH: "%UNITY_PROTO_PATH%"

REM 모든 .proto 파일에 대해 Go 및 C# 코드를 컴파일
for %%f in (*.proto) do (
    echo Generating Protobuf code for %%f...

    REM Go 코드 생성 (gRPC 제외)
    protoc --proto_path=. --go_out=../src/pb --go_opt=paths=source_relative %%f

    REM C# 코드 생성 (gRPC 제외)
    protoc --proto_path=. --csharp_out="%UNITY_PROTO_PATH%/Proto" %%f
)

REM message.proto에 대해서만 Protobuf 코드 생성
echo Generating Protobuf code for message.proto

@REM @REM Message에 대한 .proto 파일에 대한 GO, C# 코드를 컴파일
@REM protoc --proto_path=. --go_out=../src/pb --go_opt=paths=source_relative messages.proto
@REM protoc --proto_path=. --csharp_out="%UNITY_PROTO_PATH%/Messages" messages.proto

echo Protobuf code generation completed for all .proto files.
pause
