curl -v  http://localhost:8080/file/fast/upload?username=admin123\&token=3d08a2fd99afb48e7f1d06e7fcb9078c15805754\&filehash=1d1a1e3a51670715fca2805a2d9747c238b3c950\&filename=test.jpg\&filesize=580760 \
    -X POST \
    -F "file=@$(pwd)/test.jpg"

# 秒传接口设计
# POST 请求
# 参数: username, token, filehash, filename, filesize
