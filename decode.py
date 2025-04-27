import jwt  # 注意：这里导入的是 PyJWT 库
from datetime import datetime

#token = "Bearer eyJhbGciOiJIUzI1NiJ9.eyJ0b2tlbiI6IkJlYXJlciBleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbGVIQWlPakUzTkRVMU5ETXpOVGdzSW5KdmJHVWlPakVzSW5WelpYSnVZVzFsSWpvaWRHeHNJbjAuQzg3blJET21hbVpsUEs1dlJwbTV5eUJvZGVMZjMtdG5NQTlxMklodS11NCJ9.m-NH2JKUadSJx_4Fvhtkk0V8UZM-1taXv5nJoZEYxJs"
token="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDU1NDMzNTgsInJvbGUiOjEsInVzZXJuYW1lIjoidGxsIn0.C87nRDOmamZlPK5vRpm5yyBodeLf3-tnMA9q2Ihu-u4"
secret_key = "tll114514"

try:
    pure_token = token.split(" ")[1] if "Bearer" in token else token
    print(pure_token)
    # 使用 PyJWT 的解码方法
    decoded = jwt.decode(
        pure_token,
        secret_key,
        algorithms=["HS256"],
        options={"verify_exp": True}
    )

    print("解码结果：")
    print(f"- 用户名: {decoded['username']}")
    print(f"- 角色: {decoded['role']}")
    print(f"- 过期时间: {datetime.fromtimestamp(decoded['exp'])}")

except Exception as e:
    print(f"错误: {str(e)}")