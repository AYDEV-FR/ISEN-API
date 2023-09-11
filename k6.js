import http from 'k6/http';

export default function () {
  const url = 'http://localhost:8080/v1/notations';
//   const payload = JSON.stringify({
//     email: 'johndoe@example.com',
//     password: 'PASSWORD',
//   });

  const params = {
    headers: {
          'Content-Type': 'application/json',
        'Token': '<REDACTED>'
    },
  };

  http.get(url, params);
}