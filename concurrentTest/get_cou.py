import threading

import requests

# /saler_1/coupons/saler_1_6

url_base='http://localhost:20080/api/users'

authorization_token_one = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImN1c3RvbWVyXzAiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwia2luZCI6ImN1c3RvbWVyIiwiZXhwIjoxNzIxMjg0MTE0LCJpc3MiOiJ0aGlzIGlzIGEgaXNzdWVyIiwibmJmIjoxNzIxMjc5NTE0fQ.ci1gnmS0pqUzZEOh0oGFVo-uyF68oEUtaX5yqiehsy0'

authorization_token_two = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImN1c3RvbWVyXzIiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwia2luZCI6ImN1c3RvbWVyIiwiZXhwIjoxNzIxMjg0MTQ1LCJpc3MiOiJ0aGlzIGlzIGEgaXNzdWVyIiwibmJmIjoxNzIxMjc5NTQ1fQ.kDR8GAy_CETcyfpRBNgLvw_tIa_2L8NvXVCuWCnuVTk'

authorization_token_three = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImN1c3RvbWVyXzMiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwia2luZCI6ImN1c3RvbWVyIiwiZXhwIjoxNzIxMjg0MTU5LCJpc3MiOiJ0aGlzIGlzIGEgaXNzdWVyIiwibmJmIjoxNzIxMjc5NTU5fQ.lV_dLoyrvDD4KgyhBs90vxDxogHrl7pzPF5mPA-bedE'
def reqOne():
    for i in range(5):
        for j in range(10):
            url = url_base + '/saler_'+str(i)+'/coupons/saler_'+str(i)+'_'+str(j)
            response = requests.patch(
                url,
                headers={
                    'Authorization': f'{authorization_token_one}',
                    'Content-Type': 'application/json'  # 根据实际需要设置Content-Type
                },
                # 如果数据是字典，需要转换为JSON格式
            )
def reqTwo():
    for i in range(5):
        for j in range(10):
            url = url_base + '/saler_'+str(i)+'/coupons/saler_'+str(i)+'_'+str(j)
            response = requests.patch(
                url,
                headers={
                    'Authorization': f'{authorization_token_two}',
                    'Content-Type': 'application/json'  # 根据实际需要设置Content-Type
                },
                # 如果数据是字典，需要转换为JSON格式
            )
def reqThere():
    for i in range(5):
        for j in range(10):
            url = url_base + '/saler_'+str(i)+'/coupons/saler_'+str(i)+'_'+str(j)
            response = requests.patch(
                url,
                headers={
                    'Authorization': f'{authorization_token_three}',
                    'Content-Type': 'application/json'  # 根据实际需要设置Content-Type
                },
                # 如果数据是字典，需要转换为JSON格式
            )

if __name__ == '__main__':
    # 创建线程列表
    threads = []
    # 创建线程
    thread = threading.Thread(target=reqOne)
    # 将线程加入到线程列表中
    threads.append(thread)
    # 创建线程
    thread = threading.Thread(target=reqTwo)
    # 将线程加入到线程列表中
    threads.append(thread)
    # 创建线程
    thread = threading.Thread(target=reqThere)
    # 将线程加入到线程列表中
    threads.append(thread)
    # 假设我们要创建5个线程来发送请求
    for i in range(3):
        # 启动线程
        threads[i].start()

    # 等待所有线程完成
    for thread in threads:
        thread.join()

    print('所有请求已完成。')