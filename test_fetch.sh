#!/bin/bash

# Test script for brokolisql fetch mode
echo "Building brokolisql..."
go build -o brokolisql

echo "Testing fetch mode with a public REST API..."
./brokolisql --fetch --source https://jsonplaceholder.typicode.com/users --output ./examples/users.sql --table users --create-table

if [ $? -eq 0 ]; then
    echo "Fetch test successful!"
    echo "Generated SQL file:"
    head -n 20 ./examples/users.sql
else
    echo "Fetch test failed!"
fi