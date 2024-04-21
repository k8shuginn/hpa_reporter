# hpa_reporter
hpa_reporter는 Kubernetes의 Horizontal Pod Autoscaler (HPA)에서 발생하는 이벤트를 모니터링하여 사용자에게 현재 상태를 보고하는 프로그램입니다.
이 프로그램은 지정된 HPA에 대해 특정 개수 이상의 replica가 생성되었을 때 이벤트를 수집합니다.
설정된 threshold 값 이상의 replica가 생성되면 "Warning"으로, HPA 설정의 최대 허용치인 maxReplicas 이상의 replica가 생성되었을 경우 "Critical"로 알림을 보내 사용자가 상황을 인지할 수 있도록 합니다.
이렇게 hpa_reporter는 Kubernetes 환경에서 리소스 사용 상황을 효과적으로 관리하고 조정하는 데 도움을 줍니다.
