curl -v  http://localhost:8080/file/fast/upload?username=admin123\&token=f77ebe225d9e89db03d01d162561b75815801397\&filehash=1d1a1e3a51670715fca2805a2d9747c238b3c950\&filename=test.jpg\&filesize=580760 \
    -X POST \
    -F "file=@$(pwd)/test.jpg"

# 秒传接口设计
# POST 请求
# 参数: username, token, filehash, filename, filesize
