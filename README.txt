PLEASE SET GOPATH AND GOROOT env variables

after unzip:

copy entire folder into GOPATH folder

if any problem with .sh file permissions please set it
command to set permission
chmod 776 FILENAME

please check for any go dependencies in goDepend.sh file and run it if needed


run ./commands.sh file


please replace APP_KEY and APP_ID with right values then copy curl command and paste and run in seperate linux terminal tab

curl -H 'Content-Type: application/json' -d '{"req":{"q": "chicken","appKey":"APP_KEY","appId":"APP_ID"}}' "http://localhost:8080/gateway/search"
