
while true
do
  inotifywait -qq -r -e create,close_write,modify,move,delete ./ && go test  -v  $1                                   
done 