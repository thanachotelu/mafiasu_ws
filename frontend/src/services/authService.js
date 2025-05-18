import axios from "axios";

const API_URL = "http://localhost:8000/api/v1";

export const authService = {
  async register(data) {
    try {
      const response = await axios.post(`${API_URL}/users`, data);
      return response.data;
    } catch (error) {
      throw new Error(error.response?.data?.message || "Registration failed");
    }
  },

  async login(credentials) {
    try {
      const response = await axios.post(`${API_URL}/auth/login`, credentials);
      const { token, user } = response.data;
      
      // Store the token
      localStorage.setItem("token", token);
      // Set the default Authorization header
      axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
      
      return { token, user };
    } catch (error) {
      throw new Error(error.response?.data?.message || "Login failed");
    }
  },

  isAuthenticated() {
    const token = localStorage.getItem("token");
    return !!token;
  },

  logout() {
    localStorage.removeItem("token");
    delete axios.defaults.headers.common["Authorization"];
  },
};
