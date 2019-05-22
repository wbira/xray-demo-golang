.PHONY: deps clean build package create-role

deps:
	go get -u ./...

clean:
	rm -rf ./main

build:
	GOOS=linux GOARCH=amd64 go build -o ./tmp/main

package:
	zip -j tmp/main.zip tmp/main

create-role:
	aws iam create-role --role-name xray-demo1-executor --assume-role-policy-document file://tmp/trust-policy.json

deploy: build package update-function

create-function:
	aws lambda create-function --function-name xray-demo --runtime go1.x \
--role arn:aws:iam::147021510662:role/xray-demo1-executor \
--handler main --zip-file fileb://tmp/main.zip

update-function:
	aws lambda update-function-code --function-name xray-demo \
--zip-file fileb://tmp/main.zip

put-xray-policy:
	aws iam put-role-policy --role-name xray-demo1-executor \
--policy-name xray-role \
--policy-document file://tmp/xray-policy.json
