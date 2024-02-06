go mod init cobra

go install github.com/spf13/cobra-cli@latest

export PATH=$PATH:~/go/bin/

cobra-cli init
Your Cobra application is ready at
/Users/vedantj/Desktop/temp_prac/Golang_prac/gocobra

cobra-cli add timezone

go build -o cobratry .

./cobratry timezone EST
Tue, 06 Feb 2024 01:43:33 EST

./cobratry timezone EST --date 01-02-2006
Current date in EST: 02-06-2024