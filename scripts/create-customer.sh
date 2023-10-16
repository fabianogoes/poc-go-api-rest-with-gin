curl http://localhost:3000/customers \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "Customer Test"}'