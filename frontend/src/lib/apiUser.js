import { json } from "@sveltejs/kit";

const urlapi = "http://localhost:8080/api";

export const postRegister = async ({ pengguna, sandi, email }) => {
  return await fetch(`${urlapi}/pengguna`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify({ pengguna, sandi, email }),
  });
};
export const postLogin = async ({ pengguna, sandi }) => {
  return await fetch(`${urlapi}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify({ pengguna, sandi }),
    credentials: "include",
  });
};
export const postLogout = async () => {
  return await fetch(`${urlapi}/logout`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
  });
};
