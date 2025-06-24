// services/authService.js
import axios from "axios";
import { ref } from "vue";

const API_URL = "http://localhost:8000/api/v1";

// Create reactive states that can be shared across components
export const isAuthenticated = ref(false);
export const userRole = ref(null);
export const currentUser = ref(null);

// Initialize the state
const initializeAuthState = () => {
  const token = localStorage.getItem("token");
  const role = localStorage.getItem("userRole");
  const user = localStorage.getItem("currentUser");
  
  isAuthenticated.value = !!token;
  userRole.value = role;
  currentUser.value = user ? JSON.parse(user) : null;
  
  if (token) {
    axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  }
};

// Initialize on module load
initializeAuthState();

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
      const { token, user, role } = response.data;
      
      // Store the token, user, and role
      localStorage.setItem("token", token);
      localStorage.setItem("userRole", role);
      localStorage.setItem("currentUser", JSON.stringify(user));
      
      // Set the default Authorization header
      axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
      
      // Update reactive states
      isAuthenticated.value = true;
      userRole.value = role;
      currentUser.value = user;
      
      return { token, user, role };
    } catch (error) {
      throw new Error(error.response?.data?.message || "Login failed");
    }
  },

  isAuthenticated() {
    const token = localStorage.getItem("token");
    const role = localStorage.getItem("userRole");
    const user = localStorage.getItem("currentUser");
    
    const authStatus = !!token;
    isAuthenticated.value = authStatus;
    userRole.value = role;
    currentUser.value = user ? JSON.parse(user) : null;
    
    return authStatus;
  },

  logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("userRole");
    localStorage.removeItem("currentUser");
    delete axios.defaults.headers.common["Authorization"];
    
    // Update reactive states
    isAuthenticated.value = false;
    userRole.value = null;
    currentUser.value = null;
  },
};