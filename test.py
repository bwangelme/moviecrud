#!/usr/bin/env python3

####
# 执行功能测试的脚本
####

import requests
from subprocess import check_output, Popen, PIPE

MOVIE_DEMO_HOST="localhost:8080"


def reset_database():
    """ 重置数据库 """

    check_output(['mysql', '-u', 'root', '-ppasswd', '-e', 'DROP database if exists test_subject'])
    check_output(['mysql', '-u', 'root', '-ppasswd', '-e', 'CREATE database test_subject'])
    p = Popen(
        ['mysql', '-u', 'root', '-ppasswd', 'test_subject'],
        stdin=PIPE, stdout=PIPE, stderr=PIPE, close_fds=True
    )
    fd = open('config/schema.sql')
    sql = fd.read()
    fd.close()
    p.stdin.write(sql.encode())
    out, err = p.communicate()


def api_request(url, method="get", data=None, expected_code=200) -> {}:
    url = "http://%s%s" % (MOVIE_DEMO_HOST, url)
    kwargs = {
        "url": url,
        "method": method,
        "data": data,
    }
    resp = requests.request(**kwargs)
    assert resp.status_code == expected_code, "响应码错误: %s" % resp.status_code
    return resp.json()


def main():
    reset_database()
    print("小明创建了一部电影: 肖申克的救赎")
    url = "/api/movies"
    data = api_request(url, method="post", data={
        'title': '肖申克的救赎',
        'pubdate': '1994-09-10',
        'country': '美国'
    })
    assert data['error'] is None, "错误不为空: %s" % data['err']
    print("他创建的这部电影的ID是1")
    assert 1 == data['movie']['id']

    print("小明列出了网站所有的电影，目前只有一部")
    url = "/api/movies/list"
    data = api_request(url)
    assert 1 == data['total']

    print("小明查询了ID = 1 的电影，电影的名称叫 “肖申克的救赎” ")
    url = "/api/movie/1"
    data = api_request(url)
    assert "肖申克的救赎" == data['movie']['title']

    print("小明查询了ID = abc 的电影，得到了404的响应")
    url = "/api/movie/abc"
    api_request(url, expected_code=404)

    print("小明删除了ID = 1 的电影")
    url = "/api/movie/1"
    api_request(url, method="delete")

    print("小明查询了网站的所有电影，目前一部都没有")
    url = "/api/movies/list"
    data = api_request(url)
    assert 0 == data['total']


if __name__ == '__main__':
    main()
