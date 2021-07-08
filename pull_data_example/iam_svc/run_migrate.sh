#!/bin/sh

usage() {
  echo 
  echo " Error: No migration version provided."
  echo " Hints: Look into ./ops/db_migration and see what's the latest version"
  echo "        For example, this might be 000003 which means you can use ./run_migrate.sh 3"
  echo " Usage:"
  echo "   - Go to a specific version using ./run_migrate.sh {version}"
  echo "   - Force (after fixing errors) to specific version using ./run_migrate.sh {version} force"
  exit 1
}

if [ "$#" -lt 1 ] || [ "$1" = "-h" ] || [ "$1" = "-help" ]; then
  usage
  exit 1
fi

. ./env.sh
VERSION=$1


echo "Running migrate with DB_DSN=$DB_DSN and goto VERSION=$VERSION"

if [ "$#" -eq 1 ]; then
  migrate -path=./ops/db_migrations -database=$DB_DSN goto $1
elif [ "$#" -eq 2 ] && [ "$2" = "force" ]; then
  migrate -path=./ops/db_migrations -database=$DB_DSN force $1
else
  usage
fi

