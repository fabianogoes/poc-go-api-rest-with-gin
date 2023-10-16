curl http://localhost:3000/customers/3b8aab72-0b33-4025-9f7f-96d23a2aabf1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"name": "Customer updated"}'