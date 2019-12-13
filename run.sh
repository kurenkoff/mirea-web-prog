#!/usr/bin/env bash
cd frontend
npm i
npm run build
cd ..
rm -rf /ngnix/dist
mv /frontend/dist /ngnix
docker-compose build && docker-compose up