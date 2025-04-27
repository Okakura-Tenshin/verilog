package utils
import(
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
)
func HashPasswd(pwd string)(string){
	hash,_:=bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)

}
func GenerateJWT(username string,role int)(string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username":username,
		"role":role,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	})
	//Generate encoded token and send it as response.
	tokenString,err:=token.SignedString([]byte("tll114514"))
		return "Bearer "+tokenString,err
}

func Cheakpasswd(plainpwd string,hashedpwd string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedpwd),[]byte(plainpwd))
	return err==nil

}

func ParseJWT(tokens string) (string, error) {
    if len(tokens) > 7 && tokens[:7] == "Bearer " {
        tokens = tokens[7:]
    }
    token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid token")
        }
        return []byte("tll114514"), nil
    })

    if err != nil {
        return "", errors.New("invalid token")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        username, ok := claims["username"].(string);
		if !ok {
		    return "", errors.New("invalid token: username not found")
		}
		return username,nil
    }

    return "", errors.New("invalid token")
}

func GetUserRoleFromJWT(tokens string) (int,error){
	//三个用户组，0，管理员，1，老师，2，学生
	//其中管理员可以查看所有信息，老师可以查看学生信息，学生只能查看自己的信息
	//所以需要从JWT中提取用户角色
	tokens=tokens[7:]

	token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		// 确保使用 HMAC 签名方式
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte("tll114514"), nil
	})

	if err != nil {
		return -1, errors.New("token parsing failed")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		roleFloat, ok := claims["role"].(float64) // 注意：JSON中的数字默认是float64
		if !ok {
			return -1, errors.New("role not found in token")
		}
		return int(roleFloat), nil
	}

	return -1, errors.New("invalid token")
}
