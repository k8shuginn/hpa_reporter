import http from 'k6/http';

export let options = {
    stages: [
        { duration: '3m', target: 200 },  // 5분 동안 사용자 수를 0에서 100으로 증가
        { duration: '10m', target: 200 }, // 10분 동안 100 사용자 유지
        { duration: '5m', target: 0 },    // 마지막 5분 동안 사용자 수를 100에서 0으로 감소
    ],
};

export default function() {
    http.get('http://localhost:30080');
}