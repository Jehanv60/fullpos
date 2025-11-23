import { goto } from "$app/navigation";
export const fetchApi = async (responseStatus) => {
  if (responseStatus === 404) {
    goto("/");
  }
};
