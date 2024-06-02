import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [{ duration: "1s", target: 1 }],
};

export default function () {
  const url = `http://localhost:8080/admin/mechant`;

  const params = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  let bodyContent = JSON.stringify({
    name: "Merchant 1",
    merchantCategory: "MediumRestaurant",
    imageUrl: "http://image.png",
    location: {
      lat: 1,
      long: 1,
    },
  });

  const res = http.post(url, bodyContent, params);
  console.log(res.body);
  check(res, { "status was 201": (r) => r.status == 201 });
  sleep(1);
}
