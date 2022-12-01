# Learn Golang

## The Strengths of Golang
1. Concurrency
    - GoRoutine을 사용한 동시 처리
    - 한 프로세스로 모든 코어를 활용하면서 적은 메모리를 사용
    - 직접 쓰레드 사용(Java, C++) vs GoRoutine 사용(Go)
        - 프로세스 혹은 프로세스를 쪼갠 쓰레드는 OS의 스케줄링에 따라서 움직이기 때문에 동작 전 전처리가 필요하고 이로 인해 발생하는 비용이 큼 
        - 이와 달리 Go는 OS에 리소스를 요청하는 것이 아니라 Go스케줄러가 스케줄링
        - 다시 말해 런타임에서 설치 및 철거 비용이 매우 저렴
2. Very Fast


## Reference
- https://nomadcoders.co/go-for-beginners
- https://darrengwon.tistory.com/1011