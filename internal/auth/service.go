package auth

import (
	"errors"
	"fmt"
	"time"

	"ai-financial-api/config"
	"ai-financial-api/models"
	"ai-financial-api/pkg/token"

	"github.com/google/uuid"
)

// func SendOTP(mobile string) error {
//     otp := generateOTP()
//     err := SendSMS(mobile, fmt.Sprintf("Your OTP is: %s", otp))
//     if err != nil {
//         return err
//     }

//     record := models.OTPVerification{
//         Mobile:    mobile,
//         OTPCode:   otp,
//         IsVerified: false,
//     }
//     return config.DB.Create(&record).Error
// }

// func VerifyOTP(mobile, otp string) (string, error) {
//     var record models.OTPVerification
//     err := config.DB.Where("mobile = ? AND otp_code = ?", mobile, otp).Last(&record).Error
//     if err != nil || record.ID == 0 {
//         return "", errors.New("invalid OTP")
//     }

//     record.IsVerified = true
//     config.DB.Save(&record)

//     // Create or update user
//     var user models.User
//     if err := config.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
//         user = models.User{
//             UUID:   uuid.New().String(),
//             Name:   "User_" + mobile,
//             Mobile: mobile,
//         }
//         config.DB.Create(&user)
//     }

//     jwtToken, err := token.GenerateJWT(user.UUID)
//     if err != nil {
//         return "", err
//     }

//     user.JWTToken = jwtToken
//     config.DB.Save(&user)

//     return jwtToken, nil
// }

func SendOTP(mobile string) error {
    otp := "123456" // 🔐 Fixed OTP for development

    // Log instead of sending via Twilio
    fmt.Printf("[DEV] OTP for %s is %s\n", mobile, otp)

    record := models.OTPVerification{
        Mobile:     mobile,
        OTPCode:    otp,
        IsVerified: false,
    }
    return config.DB.Create(&record).Error
}


func VerifyOTP(mobile, otp string) (string, error) {
    var record models.OTPVerification
    err := config.DB.Where("mobile = ? AND otp_code = ?", mobile, otp).Last(&record).Error
    if err != nil || record.ID == 0 {
        return "", errors.New("invalid OTP")
    }

    record.IsVerified = true
    config.DB.Save(&record)

    var user models.User
    if err := config.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
        user = models.User{
            UUID:   uuid.New().String(),
            Name:   "User_" + mobile,
            Mobile: mobile,
        }
        config.DB.Create(&user)
    }

    jwtToken, err := token.GenerateJWT(user.UUID)
    if err != nil {
        return "", err
    }

    user.JWTToken = jwtToken
    config.DB.Save(&user)

    return jwtToken, nil
}


func generateOTP() string {
    return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}
