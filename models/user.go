package models 

type User struct {
    ID       int 
    Name     string
    Email    string
    Password string
    Country string
    Occupation string 
    Phone string
}

type LoginResponse struct {
    Token string
    Email    string
}

type CreateResponse struct{
    Id string 
    Email string 
    Name string
}

type ForgotPasswordRequest struct{
    Email string
}

type ResetPasswordRequest struct {
    NewPassword string 
    ConfirmNewPassword string
}
