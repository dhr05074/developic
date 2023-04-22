openapi:
	cd backend/gateway && oapi-codegen -generate types,server,spec,strict-server -package gateway -o server.gen.go ../../api/openapi/openapi.yaml
	aws s3 cp ./api/openapi/openapi.yaml s3://coplay-test/code_connect.yaml --acl public-read