const urlapi = "http://localhost:8080/api";

export const getProduct = async () => {
  return await fetch(`${urlapi}/barang`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
};
