import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 10,
  duration: '10s',
};

export default function () {
  // const res = http.get('http://localhost:8080/users/1');
  const res = http.get('http://localhost:8080/users');

  console.log(`Status: ${res.status}`);
  console.log(`Body: ${res.body}`);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  sleep(1);
}
