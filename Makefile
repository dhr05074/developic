openapi:
	cd backend/gateway && oapi-codegen -generate types,server,spec,strict-server -package gateway -o server.gen.go ../../api/openapi/api.yaml
	aws s3 cp ./api/openapi/api.yaml s3://coplay-test/code_connect.yaml --acl public-read
	portman -l ./api/openapi/api.yaml -b http://localhost:3000 --syncPostman --postmanFastSync

jaehwan:
	cd front && sh openapi.sh
