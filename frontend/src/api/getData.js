import { store } from "../store/store";

export async function apiGet(route) {
  const url = `http://${window.location.hostname}:${store.state.PORT}/${route}`;

  const respData= fetch(url)
  .then(response => response.json())
  .then(data => {return data});

  return respData
}
