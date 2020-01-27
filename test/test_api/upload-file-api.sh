curl -v  http://localhost:8080/file/upload?username=admin123\&token=f77ebe225d9e89db03d01d162561b75815801397 \
    -X POST \
    -F "file=@$(pwd)/test.jpg"
