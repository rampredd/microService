PLEASE SET GOPATH AND GOROOT env variables

cd GOPATH/src

git clone http://github.com/rampredd/microService

cd microService
if any problem with .sh file permissions please set it
command to set permission
chmod 776 FILENAME

please check for any go dependencies in goDepend.sh file and run it if needed

This need to connect to maria db server
So change database parameters in config/config.json file
run ./commands.sh file


curl -H 'Content-Type: application/json' -d '{"req":{"q": "chicken","appKey":"APP_KEY","appId":"APP_ID"}}' "http://localhost:8080/gateway/search"
