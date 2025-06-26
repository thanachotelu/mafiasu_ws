import axios from "axios";
import { authService } from "./authService";

axios.interceptors.response.use(
  response => response,
  async error => {
    const originalRequest = error.config;
    if (
      error.response &&
      error.response.status === 401 &&
      !originalRequest._retry
    ) {
      originalRequest._retry = true;
      try {
        const refresh_token = localStorage.getItem("refresh_token");
        const res = await axios.post("http://localhost:8000/api/v1/auth/refresh", { refresh_token });
        const { token, refresh_token: newRefreshToken } = res.data;
        localStorage.setItem("token", token);
        localStorage.setItem("refresh_token", newRefreshToken);
        axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
        originalRequest.headers["Authorization"] = `Bearer ${token}`;
        return axios(originalRequest); // retry request เดิม
      } catch (refreshError) {
        authService.logout();
        window.location.href = "/login";
      }
    }
    return Promise.reject(error);
  }
);