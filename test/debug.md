## Debugging 프로그램 인수
```
--app.config=/Users/hsw/0_develop/0_code/k8shuginn/hpa_reporter/test/config.yml
```

## Debugging 환경변수
```
LOG_LEVEL=DEBUG;LOG_PATH=/Users/hsw/0_develop/0_code/k8shuginn/hpa_reporter/test;KUBECONFIG=/Users/hsw/.kube/config
```

## 부하테스트
```bash
brew install k6
k6 run k6-test.js
```