# hpa_reporter
Kubernetes에서 자동 수평 확장을 위해 자주 사용되는 기능인 Horizontal Pod Autoscaler (HPA)는 설정된 maxReplicas 값까지만 pod를 확장하고, 그 이상 확장되지 않습니다. 이로 인해 사용자가 HPA가 최대 설정값까지 확장되었는지 직접 확인해야 하는 불편함이 있습니다. 이 문제를 해결하기 위해 hpa-reporter라는 도구가 개발되었습니다.
hpa-reporter는 지정된 namespace 내의 모든 HPA를 모니터링하며, 각 HPA가 관리하는 deployment의 현재 replica 수와 maxReplicas 값을 비교합니다. 만약 현재 replica 수가 특정 기준을 초과하면, 사용자에게 알림을 보내어 현재 상태를 알 수 있게 합니다. 이 시스템을 통해 사용자는 별도의 수동 검사 없이도 deployment가 최대로 설정된 replica 수에 도달했는지 쉽게 파악할 수 있습니다.
