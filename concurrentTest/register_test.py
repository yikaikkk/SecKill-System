import requests
import json


if __name__ == '__main__':
    url = 'http://localhost:20080/api/users'
    # POST请求的数据，通常是一个字典

    # for saler_id in range(5):
    #     username = 'saler_' + str(saler_id)
    #     password = '123456'
    #     payload = {'username': username,'kind':"saler", 'password': password}
    #     reqdata = json.dumps(payload)
    #     r = requests.post(url, data=reqdata)
    #     print(r.text)
    for saler_id in range(200):
        username = 'customer_' + str(saler_id)
        password = '123456'
        payload = {'username': username,'kind':"customer", 'password': password}
        reqdata = json.dumps(payload)
        r = requests.post(url, data=reqdata)
        print(r.text)

    # data = {
    #     'key1': 'value1',
    #     'key2': 'value2'
    # }
    #
    # # 发送POST请求
    # response = requests.post(url, data=data)
    #
    # # 检查请求是否成功
    # if response.status_code == 200:
    #     # 打印响应内容
    #     print(response.text)
    # else:
    #     print('请求失败，状态码：', response.status_code)
