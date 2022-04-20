#!/bin/bash
# https://github.com/pressly/goose
# https://qiita.com/H-A-L/items/fe8cb0e0ee0041ff3ceb
for ARGUMENT in "$@"
do
   KEY=$(echo $ARGUMENT | cut -f1 -d=)

   KEY_LENGTH=${#KEY}
   VALUE="${ARGUMENT:$KEY_LENGTH+1}"

   export "$KEY"="$VALUE"
done

if [[ -z $DB_USER ]];then
    export DB_USER="root"
fi

if [[ -z $DB_PASSWORD ]];then
    DB_PASSWORD="root"
fi

OPTS=("migrate" "rollback" "create" "enter" "version" "exit" )
OPT=""

set_operation() {
  echo "実行コマンドを選択してください。"
  PS3="> "
  select OPT in "${OPTS[@]}";
  do
      case $REPLY in
          [1-$((${#OPTS[@]}-1))])
              return
              ;;
          ${#OPTS[@]}) exit ;;
      esac
  done
}
set_operation


if [[ $OPT = enter ]];then
    psql -h 127.0.0.1 -U root -W
    exit
fi

if [[ -z $db ]];then
    echo -n "Type database name and press enter: ";
    read;
    db=${REPLY}
fi


if [[ $OPT = create && -n $title ]];then
    goose -dir=$db create $title sql
fi
if [[ $OPT = migrate ]];then
  goose -dir=$db postgres "user=${DB_USER} password=${DB_PASSWORD} dbname=${db} sslmode=disable" up
fi

if [[ $OPT = version ]];then
    goose postgres "user=${DB_USER} password=${DB_PASSWORD} dbname=${db} sslmode=disable" version
fi



# goose postgres "user=${DB_USER} password=${DB_PASSWORD} dbname=${db} sslmode=disable" status

# goose --dir="./task" postgres "user=root password=root dbname=task sslmode=disable" up
