## Описание
1. Все в одном файле, для упрощения
2. Генерирует хэш и один раз дергает сервер

## Build
1. docker run --net=test_isolated --ip 172.28.5.5 tcp_client 
2. docker build -t tcp_client .
