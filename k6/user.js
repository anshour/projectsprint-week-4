import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [{ duration: "1s", target: 1 }],
};

export default function () {
  const url = `http://localhost:8080/admin/user`;

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  let bodyContent = JSON.stringify({
    username: "User 1",
    password: "123456789",
    email: "user@gmail.com",
  });

  const res = http.post(url, bodyContent, params);
  console.log(res.body);
  check(res, { "status was 201": (r) => r.status == 201 });
  sleep(1);
}
