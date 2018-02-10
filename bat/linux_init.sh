#!/bin/bash
echo "What is your name?"

if [ ! -d "./log" ]; then
  mkdir ./log
fi

if [ ! -d "./log/data_server" ]; then
  mkdir ./log/data_server
fi

if [ ! -d "./log/login_server" ]; then
  mkdir ./log/login_server
fi

if [ ! -d "./log/gate_server" ]; then
  mkdir ./log/gate_server
fi