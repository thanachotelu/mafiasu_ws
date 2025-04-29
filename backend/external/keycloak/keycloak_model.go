package keycloak

// Credential สำหรับส่งไปกับ CreateUserRequest
type Credential struct {
    Type      string `json:"type"`
    Value     string `json:"value"`
    Temporary bool   `json:"temporary"`
}

// CreateUserRequest payload สำหรับสร้าง user ใน Keycloak
type CreateUserRequest struct {
    Username    string      `json:"username"`
    Email       string      `json:"email"`
    FirstName   string      `json:"firstName"`
    LastName    string      `json:"lastName"`
    Password    string      `json:"password"`
    Enabled     bool        `json:"enabled"`
    Credentials []Credential `json:"credentials,omitempty"` // เพิ่ม field นี้
}

// RoleRepresentation ใช้ตอน assign role
type RoleRepresentation struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
